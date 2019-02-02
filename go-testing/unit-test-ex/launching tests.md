# go unit test

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

