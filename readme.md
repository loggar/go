#loggar golang

##build
###build
```
C:\Users\webnl\Documents\_workspace_go\src\practice-go\stringutil (master)
λ go build
```

###run
```
C:\Users\webnl\Documents\_workspace\loggar-go\src\practice-go\helloWorld (master)
λ go run helloWorld.go
hello, world
```

##install
###set go env
>windows 
* this pc - user system variables - set GOPATH, GOBIN

>bash
* export GOPATH=$HOME/_workspace_go
* export GOBIN=$HOME/_workspace_go/bin

```
λ go env
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

###install
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

##testing
```
C:\Users\webnl\Documents\_workspace_go\src\practice-go\stringutil (master)
λ go build

C:\Users\webnl\Documents\_workspace_go\src\practice-go\stringutil (master)
λ go test
PASS
ok      practice-go/stringutil  0.740s
```

##Remote packages
```
C:\Users\webnl\Documents\_workspace_go
λ go get github.com/loggar/go/hello

C:\Users\webnl\Documents\_workspace_go
λ C:\Users\webnl\Documents\_workspace_go\bin\hello
hello, world
```
```
import "github.com/loggar/go/stringutil"
```

###Commentary
```
$ godoc regexp | grep -i parse
$ godoc regexp | grep parse
```