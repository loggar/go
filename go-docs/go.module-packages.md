# Golang modules and packages

## GOPATH

GOPATH is the path or the location in your system's disk where all the packages and modules are stored. You can get the default location of your GOPATH environment variable from the simple shell command.

```
$ echo $GOPATH
/Users/<user>/go
```

## modules

Init go mod:

```
go mod init github.com/loggar/go
```

Install go modules:

```
go get github.com/stretchr/testify/assert
```

Update a module to the recent version:

```
go get -u google.golang.org/api
```

## `go mod tidy`

```
Add missing dependencies: The command analyzes the import statements in your Go source code, fetches any missing dependencies, and adds them to your go.mod file. This is helpful when you've imported a new package in your code, and you want to update your module's dependency list.

Remove unused dependencies: The command also identifies any dependencies listed in your go.mod file that are no longer needed or used by your code. It removes those unnecessary dependencies from your go.mod file, keeping the file clean and up to date.

Update transitive dependencies: When a direct dependency of your module has its own set of dependencies (transitive dependencies), the go mod tidy command ensures that the minimum required versions of those transitive dependencies are included in your go.mod file. It helps in resolving version conflicts and avoiding dependency issues.

Prune the go.sum file: The go.sum file contains the cryptographic checksums of all the dependencies used in your project. The go mod tidy command removes the checksums of any unused dependencies, ensuring that the go.sum file is in sync with the actual dependencies listed in your go.mod file.
```

## go list go modules

```
go list -m all
```

## go module upgrade

https://go.dev/wiki/Modules#how-to-upgrade-and-downgrade-dependencies

Module upgrade util: https://github.com/oligot/go-mod-upgrade

Module upgrade example:

```
$ go list -m all | grep github.com/gorilla/handlers
github.com/gorilla/handlers v0.0.0-20161206055144-3a5767ca75ec

$ go get github.com/gorilla/handlers@v1.3.0
go: upgraded github.com/gorilla/handlers v0.0.0-20161206055144-3a5767ca75ec => v1.3.0

$ go list -m all | grep github.com/gorilla/handlers
github.com/gorilla/handlers v1.3.0
```

## clean

```sh
go clean -cache && go clean -modcache
```

if you're using a go.sum file and/or a vendor directory for vendoring dependencies, you should remove these to ensure they're freshly generated:

```sh
rm go.sum
rm -rf vendor
```

Tidying Your Module

```sh
go mod tidy
```

## Check dependencies

Check all modules in your dependency graph:

```sh
go list -m all | grep protobuf
```

For more detailed information about the specific module:

```sh
go mod why -m github.com/golang/protobuf
```

To see which of your packages import it (directly or indirectly):

```sh
go list -f '{{if .Imports}}{{.ImportPath}} imports:{{join .Imports "\n  "}}{{end}}' ./... | grep -A1 protobuf
```

Check if it's in your go.sum file:

```sh
grep github.com/golang/protobuf go.sum
```
