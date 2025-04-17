# Go dev tools

## Finding unreachable functions with `deadcode`

https://go.dev/blog/deadcode

```
$ go install golang.org/x/tools/cmd/deadcode@latest
$ deadcode -help
The deadcode command reports unreachable functions in Go programs.

Usage: deadcode [flags] package...

$ deadcode .
$ deadcode -whylive=example.com/greet.hello .
$ deadcode -test -filter=encoding/json encoding/json
```

## Coverage

```
go tool cover -html=coverage.out
```
