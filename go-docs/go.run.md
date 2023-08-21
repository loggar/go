# go run

Run a go file:

```
go run main.go
```

In Go, when you run a specific file (like `go run ./main.go`), the Go compiler only considers that file. It doesn't automatically include other files from the same package. You have two options:

Use the package directory with `go run`:

```
go run .
```

Specify all the files:

```
go run main.go same-package.go
```
