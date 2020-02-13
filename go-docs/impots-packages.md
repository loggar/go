# Importing Packages in Go

## Standard Library Packages

```go
package main

import "math/rand"

func main() {
  for i := 0; i < 10; i++ {
    println(rand.Intn(25))
  }
}
```

```go
package main

import (
  "fmt"
  "math/rand"
)

func main() {
  for i := 0; i < 10; i++ {
    fmt.Printf("%d) %d\n", i, rand.Intn(25))
  }
}
```

## Installing Packages

```
$ go get github.com/gobuffalo/flect
```

The go get tool will find the package, on GitHub in this case, and install it into your `$GOPATH`. `$GOPATH/src/github.com/gobuffalo/flect`

### Update packages

```
$ go get -u github.com/gobuffalo/flect
```

## Aliasing

```go
package main

import (
 f "fmt"
  "math/rand"
)

func main() {
  for i := 0; i < 10; i++ {
    f.Printf("%d) %d\n", i, rand.Intn(25))
  }
}
```
