# Inspecting errors

ref) https://dave.cheney.net/2014/12/24/inspecting-errors

The common contract for functions which return a value of the interface type error, is the caller should not presume anything about the state of the other values returned from that call without first checking the error.

In the majority of cases, error values returned from functions should be opaque to the caller. That is to say, a test that error is nil indicates if the call succeeded or failed, and that’s all there is to it.

A small number of cases, generally revolving around interactions with the world outside your process, like network activity, require that the caller investigate the nature of the error to decide if it is reasonable to retry the operation.

A common request for package authors is to return errors of a known public type, so the caller can type assert and inspect them. I believe this practice leads to a number of undesirable outcomes:

* Public error types increase the surface area of the package’s API.
* New implementations must only return types specified in the interface’s declaration, even if they are a poor fit.
* The error type cannot be changed or deprecated after introduction without breaking compatibility, making for a brittle API.

Callers should feel no more comfortable asserting an error is a particular type than they would be asserting the string returned from Error() matches a particular pattern.

Instead I present a suggestion that permits package authors and consumers to communicate about their intention, without having to overly couple their implementation to the caller.

## Assert errors for behaviour, not type

Don’t assert an error value is a specific type, but rather assert that the value implements a particular behaviour.

This suggestion fits the has a nature of Go’s implicit interfaces, rather than the is a [subtype of] nature of inheritance based languages. Consider this example:

``` go
func isTimeout(err error) bool {
        type timeout interface {
                Timeout() bool
        }
        te, ok := err.(timeout)
        return ok && te.Timeout()
}
```

The caller can use isTimeout() to determine if the error is related to a timeout, via its implementation of the timeout interface, and then confirm if the error was timeout related — all without knowing anything about the type, or the original source of the error value.

Gift wrapping errors, usually by libraries that annotate the error path, is enabled by this method; providing that the wrapped error types also implement the interfaces of the error they wrap.

This may seem like an insoluble problem, but in practice there are relatively few interface methods that are in common use, so Timeout() bool and Temporary() bool would cover a large set of the use cases.

## In conclusion

Don’t assert errors for type, assert for behaviour.

For package authors, if your package generates errors of a temporary nature, ensure you return error types that implement the respective interface methods. If you wrap error values on the way out, ensure that your wrappers respect the interface(s) that the underlying error value implemented.

For package users, if you need to inspect an error, use interfaces to assert the behaviour you expect, not the error’s type. Don’t ask package authors for public error types; ask that they make their types conform to common interfaces by supplying Timeout() or Temporary() methods as appropriate.
