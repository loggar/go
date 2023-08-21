# go test

## test all files

```
go test ./...
go test -v ./...
```

## clean test-cache and test:

```
GOCACHE=off go test ./...
go clean -testcache && go test ./...
go test -count=1 ./...
```

## launching tests (vervose)

```
λ go test
PASS
ok      github.com/loggar/go/go-testing/unit-test-ex    0.120s

λ go test -v
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestSumTable
--- PASS: TestSumTable (0.00s)
PASS
ok      github.com/loggar/go/go-testing/unit-test-ex    0.118s

λ go test github.com\loggar\go\go-testing\unit-test-ex
ok      github.com/loggar/go/go-testing/unit-test-ex    0.107s

λ go test github.com\loggar\go\go-testing\unit-test-ex -v
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestSumTable
--- PASS: TestSumTable (0.00s)
PASS
ok      github.com/loggar/go/go-testing/unit-test-ex    0.102s
```

## statement coverage

```
λ go test -cover
PASS
coverage: 50.0% of statements
ok      github.com/loggar/go/go-testing/unit-test-ex    0.112s
```

with coverage output

```
λ go test -cover -coverpofile=c.out
```

## Log

If you want to print some debug information when running tests, you should use the Log function from the testing package. Here's how you can use it:

```go
func TestSomething(t *testing.T) {
    fmt.Println("Pi = ", math.Pi)

    t.Log("This will be displayed when running tests")
    t.Logf("This will be displayed when running tests %d", 1)
}

func TestSomething(t *testing.T) {
    t.Error("This will be displayed only if the test fails")
}
```

You can use the -v flag when running go test to see all t.Log statements, like this:

```
go test -v ./...
```

To run Go tests with the -v (verbose) flag in Visual Studio Code, you will need to modify the settings for your Go extension.

```json
{
  "go.testFlags": ["-v"]
}
```

## Run a fuzz test

```
go test -fuzz=FuzzEqual -v go-testing/fuzz-test/fuzz-test.1.fuzz_test.go
go test -fuzz=FuzzEqual -fuzztime=10s -v go-testing/fuzz-test/fuzz-test.1.fuzz_test.go
go test -fuzz=FuzzEqual -fuzztime=10m -v go-testing/fuzz-test/fuzz-test.1.fuzz_test.go
```
