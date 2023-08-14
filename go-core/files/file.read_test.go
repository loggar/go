package core_files

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTxtAsString(t *testing.T) {
	text, err := os.ReadFile("_sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	assert.NoError(t, err, "read file")

	assert.Equal(t, "", text, "read a text file as a string")
}
