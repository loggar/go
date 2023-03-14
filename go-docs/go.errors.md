# errors

error msg:

```
return errors.Errorf("%s", key)
```

error type check:

```
err := doSomething()

if err != nil {
    switch err.(type) {
        case MyErrorType1:
            // handle error of type MyErrorType1
        case MyErrorType2:
            // handle error of type MyErrorType2
        default:
            // handle other error types
    }
}
```
