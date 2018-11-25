# Stack traces and the errors package

ref) https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package

## Wrapping and stack traces

``` go
package main

import "fmt"
import "github.com/pkg/errors"

func main() {
        err := errors.New("error")
        err = errors.Wrap(err, "open failed")
        err = errors.Wrap(err, "read config failed")

        fmt.Println(err) // read config failed: open failed: error
}
```

Wraping an error added context to the underlying error and recorded the file and line that the error occurred. This file and line information could be retrieved via a helper function, Fprint, to give a trace of the execution path leading away from the error. More on that later.

However, when I came to integrate the errors package into my own projects, I found that using Wrap at each call site in the return path often felt redundant. For example:

``` go
func readconfig(file string) {
        if err := openfile(file); err != nil {
                return errors.Wrap(err, "read config failed")
        }
        // ...
}
```

If openfile failed it would likely annotate the error it returned with open failed, and that error would also include the file and line of the openfile function. Similarly, readconfig‘s wrapped error would be annotated with read config failed as well as the file and line of the call to errors.Wrap inside the readconfig function.

I realised that, at least in my own code, it is likely that the name of the function contains sufficient information to frequently make the additional context passed to Wrap redundant. But as Wrap requires a message, even if I had nothing useful to add, I’d still have to pass something:

``` go
if err != nil {
        return errors.Wrap(err, "") // ewww
}
```

I briefly considered making Wrap variadic–to make the second parameter optional–before realising that rather than forcing the user to manually annotate each stack frame in the return path, I can just record the entire stack trace at the point that an error is created by the errors package.

I believe that for 90% of the use cases, this natural stack trace–that is the trace collected at the point New or Errorf are called–is correct with respect to the information required to investigate the error’s cause. In the other cases, Wrap and Wrapf can be used to add context when needed.

This lead to a large internal refactor of the package to collect and expose this natural stack trace.

## Fprint and Print have been removed

As mentioned earlier, the mechanism for printing not just the err.Error() text of an error, but also its stack trace, has also changed with feedback from early users.

The first attempts were a pair of functions; Print(err error), which printed the detailed error to os.Stderr, and Fprint(w io.Writer, err error) which did the same but allowed the caller to control the destination. Neither were very popular.

Print was removed in version 0.4.0 because it was just a wrapper around Fprint(os.Stderr, err) and was hard to test, harder to write an example test for, and didn’t feel like its three lines paid their way. However, with Print gone, users were unhappy that Fprint required you to pass an io.Writer, usually a bytes.Buffer, just to retrieve a string form of the error’s trace.

So, Print and Fprint were the wrong API. They were too opinionated, without it being a useful opinion. Fprint has been slowly gutted over the period of 0.5, 0.6 and now has been replaced with a much more powerful facility inspired by Chris Hines’ go-stack/stack package.

The errors package now leverages the powerful fmt.Formatter interface to allow it to customise its output when any error generated, or wrapped by this package, is passed to fmt.Printf. This extended format is activated by the %+v verb. For example,

``` go
func main() {
        err := parseArgs(os.Args[1:])
        fmt.Printf("%v\n", err)
}
```

Prints, as expected,

```
not enough arguments, expected at least 3, got 0
```

However if we change the formatting verb to `%+v`,

``` go
func main() {
        err := parseArgs(os.Args[1:])
        fmt.Printf("%+v\n", err)
}
```

the same error value now results in

```
not enough arguments, expected at least 3, got 0
main.parseArgs
        /home/dfc/src/github.com/pkg/errors/_examples/wrap/main.go:12
main.main
        /home/dfc/src/github.com/pkg/errors/_examples/wrap/main.go:18
runtime.main
        /home/dfc/go/src/runtime/proc.go:183
runtime.goexit
        /home/dfc/go/src/runtime/asm_amd64.s:2059
```

For those that need more control the Cause and StackTrace behaviours return values who have their own fmt.Formatter implementations. The latter is alias for a slice of Frame values which represent each frame in a call stack. Again, Frame implements several fmt.Formatter verbs that allow its output to be customised as required.

## Putting it all together

With the changes to the errors package, some guidelines on how to use the package are in order.

* In your own code, use errors.New or errors.Errorf at the point an error occurs.

``` go
func parseArgs(args []string) error {
        if len(args) < 3 {
                return errors.Errorf("not enough arguments, expected at least 3, got %d", len(args))
        }
        // ...
}
```

* If you receive an error from another function, it is often sufficient to simply return it.

``` go
if err != nil {
       return err
}
```

* If you interact with a package from another repository, consider using errors.Wrap or errors.Wrapf to establish a stack trace at that point. This advice also applies when interacting with the standard library.

``` go
f, err := os.Open(path)
if err != nil {
        return errors.Wrapf(err, "failed to open %q", path)
}
```

* Always return errors to their caller rather than logging them throughout your program.
* At the top level of your program, or worker goroutine, use %+v to print the error with sufficient detail.

``` go
func main() {
        err := app.Run()
        if err != nil {
                fmt.Printf("FATAL: %+v\n", err)
                os.Exit(1)
        }
}
```

* If you want to exclude some classes of error from printing, use errors.Cause to unwrap errors before inspecting them.

## Conclusion

The errors package, from the point of view of the four package level functions, New, Errorf, Wrap, and Wrapf, is done. Their API signatures are well tested, and now this package has been integrated into over 100 other packages, are unlikely to change at this point.

The extended stack trace format, %+v, is still very new and I encourage you to try it and leave feedback via an issue.
