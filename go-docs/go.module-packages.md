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