package os_txt_files

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteSliceString(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
	}
	assert.NoError(t, err, "get current directory")

	outFilePath := filepath.Join(cwd, "../../_out/go-files/txt", "data.out.1.txt")

	// Open the file if it exists, otherwise create a new one
	f, err := os.OpenFile(outFilePath, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Error opening or creating file:", err)
	}
	assert.NoError(t, err, "open file")

	defer f.Close() // Don't forget to close the file when done

	langs := []string{"golang", "python", "rust", "javascript", "ruby"}
	for _, lang := range langs {
		_, err := f.WriteString(lang + "\n")
		assert.NoError(t, err, "write file")
	}
	defer f.Close()
}
