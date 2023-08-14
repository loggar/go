package core_files

import (
	"encoding/csv"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileByDelimiter(t *testing.T) {
	f, err := os.Open("_delimiter.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.Comma = ':'

	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	expected := [][]string{{"10", "22", "2022"}, {"golang", "21", "read"}}

	assert.Equal(t, expected, data, "read a scv by delimiter")
}
