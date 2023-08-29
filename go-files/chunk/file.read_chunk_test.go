package os_txt_files

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileByChunk(t *testing.T) {
	f, err := os.Open("_sample_large.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	chunk_size := 16
	chunk_list := []string{}
	buf := make([]byte, chunk_size)

	for {
		n, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		chunk_list = append(chunk_list, string(buf[0:n]))
	}

	expected := []string{"In the above exa", "mple, we have op", "ened the file an", "d loaded the con", "tent into the f ", "variable. the co", "ntents are read ", "with the help of", " the NewReader f", "unction which re", "turns a reader o", "bject which furt", "her can be used ", "to read contents", " into chunks of ", "bytes. The chunk", "_size defines th", "e size we want t", "o use for readin", "g the content, c", "hunk_list as a s", "lice of strings ", "which will hold ", "the slice of chu", "nks/bytes as a t", "ype caste into a", " slice of string", "s. With the Read", " function, the b", "ytes are read in", "to the function,", " and the buffer ", "is split as per ", "the chunk size o", "btained in the R", "ead function. We", " append the slic", "e of bytes into ", "the sliced array", " and thereby we ", "obtain the slice", " of strings."}

	assert.Equal(t, expected, chunk_list, "read a file by chunk")
}
