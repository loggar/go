Charly

# Golang

## build and run

### build

```
C:\Users\webnl\Documents\_workspace_go\src\practice-go\stringutil (master)
λ go build
```

### run

```
C:\Users\webnl\Documents\_workspace\loggar-go\src\practice-go\helloWorld (master)
λ go run helloWorld.go
hello, world
```

## install

### set go env

> windows

- this pc - user system variables - set GOPATH, GOBIN

> bash

- export GOPATH=\$HOME/\_workspace_go
- export GOBIN=\$HOME/\_workspace_go/bin

```
set GOARCH=amd64
set GOBIN=C:\Users\webnl\Documents\_workspace_go\bin
set GOEXE=.exe
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOOS=windows
set GOPATH=C:\Users\webnl\Documents\_workspace_go
set GORACE=
set GOROOT=C:\_dev\Go
set GOTOOLDIR=C:\_dev\Go\pkg\tool\windows_amd64
set GCCGO=gccgo
set CC=gcc
set GOGCCFLAGS=-m64 -mthreads -fmessage-length=0 -fdebug-prefix-map=C:\Users\webnl\AppData\Local\Temp\go-build799376399=/tmp/go-build -gno-record-gcc-switches
set CXX=g++
set CGO_ENABLED=1
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
```

### install

```
C:\Users\webnl\Documents\_workspace_go\src\practice-go (master)
λ go install practice-go/hello
```

or

```
C:\Users\webnl\Documents\_workspace_go\src\practice-go (master)
λ cd hello
C:\Users\webnl\Documents\_workspace_go\src\practice-go\hello (master)
λ go install
```

```
C:\Users\webnl\Documents\_workspace_go\bin
λ hello
hello, world
```

## testing

```
C:\Users\webnl\Documents\_workspace_go\src\practice-go\stringutil (master)
λ go build

C:\Users\webnl\Documents\_workspace_go\src\practice-go\stringutil (master)
λ go test
PASS
ok      practice-go/stringutil  0.740s
```

### Commentary

```
$ godoc regexp | grep -i parse
$ godoc regexp | grep parse
```

## Remote packages

```
C:\Users\webnl\Documents\_workspace_go
λ go get github.com/loggar/go/hello

C:\Users\webnl\Documents\_workspace_go
λ C:\Users\webnl\Documents\_workspace_go\bin\hello
hello, world
```

```
go get github.com/satori/go.uuid
```

## import packages

```
import "github.com/loggar/go/stringutil"
```

## Creating a custom package

```
> mkdir custom_package
> cd custom_package
```

To create a custom package we need to first create a folder with the package name we need. Let’s say we are building a package person. For that let’s create a folder named person inside custom_package folder:

```
> mkdir person
> cd person
```

Now let’s create a file person.go inside this folder.

```go
// package.go

package person
func Description(name string) string {
  return "The person name is: " + name
}
func secretName(name string) string {
  return "Do not share"
}
```

We now need to install the package so that it can be imported and used. So let’s install it:

```
> go install
```

Now let’s go back to the custom_package folder and create a main.go file

```go
// usepackage.go


package main
import(
  "custom_package/person"
  "fmt"
)
func main(){
  p := person.Description("Milap")
  fmt.Println(p)
}
// => The person name is: Milap
```

Here we can now import the package person we created and use the function Description. Note that the function secretName we created in the package will not be accessible. In Go, the method name starting without a capital letter will be private.

## Packages Documentation

Go has built-in support for documentation for packages. Run the following command to generate documentation:

```
godoc person Description
```

This will generate documentation for the Description function inside our package person. To see the documentation run a web server using the following command:

```
godoc -http=":8080"
```

Now go to the URL http://localhost:8080/pkg/ and see the documentation of the package we just created.
