# Error Handling

Errors are the undesired and unexpected result of a program. Let’s say we are making an API call to an external service. This API call may be successful or could fail. An error in a Go program can be recognized when an error type is present. Let’s see the example:

```go
resp, err := http.Get("http://example.com/")
```

Here the API call to the error object may pass or could fail. We can check if the error is nil or present and handle the response accordingly:

```go
package main

import (
  "fmt"
  "net/http"
)

func main(){
  resp, err := http.Get("http://example.com/")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(resp)
}
```

## Returning custom error from a function

When we are writing a function of our own, there are cases when we have errors. These errors can be returned with the help of the error object:

```go
func Increment(n int) (int, error) {
  if n < 0 {
    // return error object
    return nil, errors.New("math: cannot process negative number")
  }
  return (n + 1), nil
}

func main() {
  num := 5

  if inc, err := Increment(num); err != nil {
    fmt.Printf("Failed Number: %v, error message: %v", num, err)
  } else {
    fmt.Printf("Incremented Number: %v", inc)
  }
}
```

Most of the packages that are built in Go, or external packages we use, have a mechanism for error handling. So any function we call could have possible errors. These errors should never be ignored and always handled gracefully in the place we call these functions, as we have done in the above example.

## Panic

Panic is something that is unhandled and is suddenly encountered during a program execution. In Go, panic is not the ideal way to handle exceptions in a program. It is recommended to use an error object instead. When a panic occurs, the program execution get’s halted. The thing that gets executed after a panic is the defer.

## Defer

Defer is something that will always get executed at the end of a function.

```go
package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
```

In the above example, we panic the execution of the program using panic(). As you notice, there is a defer statement which will make the program execute the line at the end of the execution of the program. Defer can also be used when we need something to be executed at the end of the function, for example closing a file.
