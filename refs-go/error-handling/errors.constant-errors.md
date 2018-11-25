# Constant Errors

ref) https://dave.cheney.net/2016/04/07/constant-errors

This is a thought experiment about sentinel error values in Go.

Sentinel errors are bad, they introduce strong source and run time coupling, but are sometimes necessary. io.EOF is one of these sentinel values. Ideally a sentinel value should behave as a constant, that is it should be immutable and fungible.

The first problem is io.EOF is a public variable–any code that imports the io package could change the value of io.EOF. It turns out that most of the time this isn’t a big deal, but it could be a very confusing problem to debug.

``` go
fmt.Println(io.EOF == io.EOF) // true
x := io.EOF
fmt.Println(io.EOF == x)      // true
	
io.EOF = fmt.Errorf("whoops")
fmt.Println(io.EOF == io.EOF) // true
fmt.Println(x == io.EOF)      // false
```

The second problem is io.EOF behaves like a singleton, not a constant. Even if we follow the exact procedure used by the io package to create our own EOF value, they are not comparable.

``` go
err := errors.New("EOF")   // io/io.go line 38
fmt.Println(io.EOF == err) // false
```

Combine these properties and you have a set of weird behaviours stemming from the fact that sentinel error values in Go, those traditionally created with errors.New or fmt.Errorf, are not constants.

## Constant errors

Before I introduce my solution, let’s recap how the error interface works in Go. Any type with an Error() string method fulfils the error interface. This includes primitive types like string, including constant strings.

With that background, consider this error implementation.

``` go
type Error string

func (e Error) Error() string { return string(e) }
```

It looks similar to the errors.errorString implementation that powers errors.New. However unlike errors.errorString this type is a constant expression.

``` go
const err = Error("EOF") 
const err2 = errorString{"EOF"} // const initializer errorString literal is not a constant
```

As constants of the Error type are not variables, they are immutable.

``` go
const err = Error("EOF") 
err = Error("not EOF") // error, cannot assign to err
```

Additionally, two constant strings are always equal if their contents are equal, which means two Error values with the same contents are equal.

``` go
const err = Error("EOF") 
fmt.Println(err == Error("EOF")) // true
```

Said another way, equal Error values are the same, in the way that the constant 1 is the same as every other constant 1.

``` go
const eof = Error("eof")

type Reader struct{}

func (r *Reader) Read([]byte) (int, error) {
        return 0, eof
}

func main() {
        var r Reader
        _, err := r.Read([]byte{})
        fmt.Println(err == eof) // true
}
```

Could we change the definition of io.EOF to be a constant?S It turns out that this compiles just fine and passes all the tests, but it’s probably a stretch for the Go 1 contract.

However this does not prevent you from using this idiom in your own code. Although, you really shouldn’t be using sentinel errors anyway.
