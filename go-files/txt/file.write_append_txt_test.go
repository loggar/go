package os_txt_files

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteAppend(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
	}
	assert.NoError(t, err, "get current directory")

	outFilePath := filepath.Join(cwd, "../../_out/go-files/txt", "data.out.2.txt")

	s := "Hello"
	err = os.WriteFile(outFilePath, []byte(s), 0660)
	assert.NoError(t, err, "write file")

	s = " World!"
	f, err := os.OpenFile(outFilePath, os.O_APPEND|os.O_WRONLY, 0660)
	assert.NoError(t, err, "open file")

	_, err = f.WriteString(s)
	assert.NoError(t, err, "append string to file")

	defer f.Close()
}
