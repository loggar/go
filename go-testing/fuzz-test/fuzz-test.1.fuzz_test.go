package fuzztestingingo

import (
	"log"
	"testing"

	bu "github.com/loggar/go/go-lib-src/byteutil"
)

// go test -fuzz=FuzzEqual -v go-testing/fuzz-test/fuzz-test.1.fuzz_test.go
// go test -fuzz=FuzzEqual -fuzztime=10s -v go-testing/fuzz-test/fuzz-test.1.fuzz_test.go
// go test -fuzz=FuzzEqual -fuzztime=10m -v go-testing/fuzz-test/fuzz-test.1.fuzz_test.go

func FuzzEqual(f *testing.F) {
	// target, can be only one per test
	// values of a and b will be auto-generated
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		log.Printf("Fuzzing with a: %v, b: %v\n", a, b) // Logging each input
		bu.Equal(a, b)
	})
}
