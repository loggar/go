# Don’t just check errors, handle them gracefully

ref) https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

## Errors are just values

I’ve spent a lot of time thinking about the best way to handle errors in Go programs. I really wanted there to be a single way to do error handling, something that we could teach all Go programmers by rote, just as we might teach mathematics, or the alphabet.

However, I have concluded that there is no single way to handle errors. Instead, I believe Go’s error handling can be classified into the three core strategies.

## Sentinel errors

The first category of error handling is what I call sentinel errors.

```
if err == ErrSomething { … }
```

The name descends from the practice in computer programming of using a specific value to signify that no further processing is possible. So to with Go, we use specific values to signify an error.

Examples include values like io.EOF or low level errors like the constants in the syscall package, like syscall.ENOENT.

There are even sentinel errors that signify that an error did not occur, like go/build.NoGoError, and path/filepath.SkipDir from path/filepath.Walk.

Using sentinel values is the least flexible error handling strategy, as the caller must compare the result to predeclared value using the equality operator. This presents a problem when you want to provide more context, as returning a different error would will break the equality check.

Even something as well meaning as using fmt.Errorf to add some context to the error will defeat the caller’s equality test. Instead the caller will be forced to look at the output of the error‘s Error method to see if it matches a specific string.

### Never inspect the output of error.Error

As an aside, I believe you should never inspect the output of the error.Error method. The Error method on the error interface exists for humans, not code.

The contents of that string belong in a log file, or displayed on screen. You shouldn’t try to change the behaviour of your program by inspecting it.

I know that sometimes this isn’t possible, and as someone pointed out on twitter, this advice doesn’t apply to writing tests. Never the less, comparing the string form of an error is, in my opinion, a code smell, and you should try to avoid it.

### Sentinel errors become part of your public API

If your public function or method returns an error of a particular value then that value must be public, and of course documented. This adds to the surface area of your API.

If your API defines an interface which returns a specific error, all implementations of that interface will be restricted to returning only that error, even if they could provide a more descriptive error.

We see this with io.Reader. Functions like io.Copy require a reader implementation to return exactly io.EOF to signal to the caller no more data, but that isn’t an error.

### Sentinel errors create a dependency between two packages

By far the worst problem with sentinel error values is they create a source code dependency between two packages. As an example, to check if an error is equal to io.EOF, your code must import the io package.

This specific example does not sound so bad, because it is quite common, but imagine the coupling that exists when many packages in your project export error values, which other packages in your project must import to check for specific error conditions.

Having worked in a large project that toyed with this pattern, I can tell you that the spectre of bad design–in the form of an import loop–was never far from our minds.

### Conclusion: avoid sentinel errors

So, my advice is to avoid using sentinel error values in the code you write. There are a few cases where they are used in the standard library, but this is not a pattern that you should emulate.

If someone asks you to export an error value from your package, you should politely decline and instead suggest an alternative method, such as the ones I will discuss next.

## Error types

Error types are the second form of Go error handling I want to discuss.

```
if err, ok := err.(SomeType); ok { … }
```

An error type is a type that you create that implements the error interface. In this example, the MyError type tracks the file and line, as well as a message explaining what happened.

``` go
type MyError struct {
        Msg string
        File string
        Line int
}

func (e *MyError) Error() string { 
        return fmt.Sprintf("%s:%d: %s”, e.File, e.Line, e.Msg)
}

return &MyError{"Something happened", “server.go", 42}
```

Because MyError error is a type, callers can use type assertion to extract the extra context from the error.

```
err := something()
switch err := err.(type) {
case nil:
        // call succeeded, nothing to do
case *MyError:
        fmt.Println(“error occurred on line:”, err.Line)
default:
// unknown error
}
```

A big improvement of error types over error values is their ability to wrap an underlying error to provide more context.

An excellent example of this is the os.PathError type which annotates the underlying error with the operation it was trying to perform, and the file it was trying to use.

``` go
// PathError records an error and the operation
// and file path that caused it.
type PathError struct {
        Op   string
        Path string
        Err  error // the cause
}

func (e *PathError) Error() string
```

### Problems with error types

So the caller can use a type assertion or type switch, error types must be made public.

If your code implements an interface whose contract requires a specific error type, all implementors of that interface need to depend on the package that defines the error type.

This intimate knowledge of a package’s types creates a strong coupling with the caller, making for a brittle API.

### Conclusion: avoid error types

While error types are better than sentinel error values, because they can capture more context about what went wrong, error types share many of the problems of error values.

So again my advice is to avoid error types, or at least, avoid making them part of your public API.

## Opaque errors

Now we come to the third category of error handling. In my opinion this is the most flexible error handling strategy as it requires the least coupling between your code and caller.

I call this style opaque error handling, because while you know an error occurred, you don’t have the ability to see inside the error. As the caller, all you know about the result of the operation is that it worked, or it didn’t.

This is all there is to opaque error handling–just return the error without assuming anything about its contents. If you adopt this position, then error handling can become significantly more useful as a debugging aid.

``` go
import “github.com/quux/bar”

func fn() error {
        x, err := bar.Foo()
        if err != nil {
                return err
        }
        // use x
}
```

For example, Foo‘s contract makes no guarantees about what it will return in the context of an error. The author of Foo is now free to annotate errors that pass through it with additional context without breaking its contract with the caller.

### Assert errors for behaviour, not type

In a small number of cases, this binary approach to error handling is not sufficient.

For example, interactions with the world outside your process, like network activity, require that the caller investigate the nature of the error to decide if it is reasonable to retry the operation.

In this case rather than asserting the error is a specific type or value, we can assert that the error implements a particular behaviour. Consider this example:

``` go
type temporary interface {
        Temporary() bool
}
 
// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
        te, ok := err.(temporary)
        return ok && te.Temporary()
}
```

We can pass any error to IsTemporary to determine if the error could be retried.

If the error does not implement the temporary interface; that is, it does not have a Temporary method, then then error is not temporary.

If the error does implement Temporary, then perhaps the caller can retry the operation if Temporary returns true.

The key here is this logic can be implemented without importing the package that defines the error or indeed knowing anything about err‘s underlying type–we’re simply interested in its behaviour.

## Don’t just check errors, handle them gracefully

This brings me to a second Go proverb that I want to talk about; don’t just check errors, handle them gracefully. Can you suggest some problems with the following piece of code?

``` go
func AuthenticateRequest(r *Request) error {
        err := authenticate(r.User)
        if err != nil {
                return err
        }
        return nil
}
```

An obvious suggestion is that the five lines of the function could be replaced with

``` go
return authenticate(r.User)
```

But this is the simple stuff that everyone should be catching in code review. More fundamentally the problem with this code is I cannot tell where the original error came from.

If authenticate returns an error, then AuthenticateRequest will return the error to its caller, who will probably do the same, and so on. At the top of the program the main body of the program will print the error to the screen or a log file, and all that will be printed is: No such file or directory.

There is no information of file and line where the error was generated. There is no stack trace of the call stack leading up to the error. The author of this code will be forced to a long session of bisecting their code to discover which code path trigged the file not found error.

Donovan and Kernighan’s The Go Programming Language recommends that you add context to the error path using fmt.Errorf

``` go
func AuthenticateRequest(r *Request) error {
        err := authenticate(r.User)
        if err != nil {
                return fmt.Errorf("authenticate failed: %v", err)
        }
        return nil
}
```

But as we saw earlier, this pattern is incompatible with the use of sentinel error values or type assertions, because converting the error value to a string, merging it with another string, then converting it back to an error with fmt.Errorf breaks equality and destroys any context in the original error.

### Annotating errors

I’d like to suggest a method to add context to errors, and to do that I’m going to introduce a simple package. The code is online at github.com/pkg/errors. The errors package has two main functions:

``` go
// Wrap annotates cause with a message.
func Wrap(cause error, message string) error
```

The first function is Wrap, which takes an error, and a message and produces a new error.

``` go
// Cause unwraps an annotated error.
func Cause(err error) error
```

The second function is Cause, which takes an error that has possibly been wrapped, and unwraps it to recover the original error.

Using these two functions, we can now annotate any error, and recover the underlying error if we need to inspect it. Consider this example of a function that reads the content of a file into memory.

``` go
func ReadFile(path string) ([]byte, error) {
        f, err := os.Open(path)
        if err != nil {
                return nil, errors.Wrap(err, "open failed")
        } 
        defer f.Close()
 
        buf, err := ioutil.ReadAll(f)
        if err != nil {
                return nil, errors.Wrap(err, "read failed")
        }
        return buf, nil
}
```

We’ll use this function to write a function to read a config file, then call that from main.

``` go
func ReadConfig() ([]byte, error) {
        home := os.Getenv("HOME")
        config, err := ReadFile(filepath.Join(home, ".settings.xml"))
        return config, errors.Wrap(err, "could not read config")
}
 
func main() {
        _, err := ReadConfig()
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
```

If the ReadConfig code path fails, because we used errors.Wrap, we get a nicely annotated error in the K&D style.

```
could not read config: open failed: open /Users/dfc/.settings.xml: no such file or directory
```

Because errors.Wrap produces a stack of errors, we can inspect that stack for additional debugging information. This is the same example again, but this time we replace fmt.Println with errors.Print

``` go
func main() {
        _, err := ReadConfig()
        if err != nil {
                errors.Print(err)
                os.Exit(1)
        }
}
```

We’ll get something like this:

```
readfile.go:27: could not read config
readfile.go:14: open failed
open /Users/dfc/.settings.xml: no such file or directory
```

The first line comes from ReadConfig, the second comes from the os.Open part of ReadFile, and the remainder comes from the os package itself, which does not carry location information.

Now we’ve introduced the concept of wrapping errors to produce a stack, we need to talk about the reverse, unwrapping them. This is the domain of the errors.Cause function.

``` go
// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
        te, ok := errors.Cause(err).(temporary)
        return ok && te.Temporary()
}
```

In operation, whenever you need to check an error matches a specific value or type, you should first recover the original error using the errors.Cause function.

### Only handle errors once

Lastly, I want to mention that you should only handle errors once. Handling an error means inspecting the error value, and making a decision.

``` go
func Write(w io.Writer, buf []byte) {
        w.Write(buf)
}
```

If you make less than one decision, you’re ignoring the error. As we see here, the error from w.Write is being discarded.

But making more than one decision in response to a single error is also problematic.

``` go
func Write(w io.Writer, buf []byte) error {
        _, err := w.Write(buf)
        if err != nil {
                // annotated error goes to log file
                log.Println("unable to write:", err)
 
                // unannotated error returned to caller
                return err
        }
        return nil
}
```

In this example if an error occurs during Write, a line will be written to a log file, noting the file and line that the error occurred, and the error is also returned to the caller, who possibly will log it, and return it, all the way back up to the top of the program.

So you get a stack of duplicate lines in your log file, but at the top of the program you get the original error without any context. Java anyone?

``` go
func Write(w io.Write, buf []byte) error {
        _, err := w.Write(buf)
        return errors.Wrap(err, "write failed")
}
```

Using the errors package gives you the ability to add context to error values, in a way that is inspectable by both a human and a machine.

## Conclusion

In conclusion, errors are part of your package’s public API, treat them with as much care as you would any other part of your public API.

For maximum flexibility I recommend that you try to treat all errors as opaque. In the situations where you cannot do that, assert errors for behaviour, not type or value.

Minimise the number of sentinel error values in your program and convert errors to opaque errors by wrapping them with errors.Wrap as soon as they occur.

Finally, use errors.Cause to recover the underlying error if you need to inspect it.
