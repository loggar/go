package os_csv_files

import (
	"encoding/csv"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileFromUrl(t *testing.T) {
	url := "https://github.com/loggar/go/raw/master/_in/go-files/csv/data.ex.1.csv"
	response, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	reader := csv.NewReader(response.Body)
	n, err := reader.ReadAll()

	if err != nil {
		log.Println(err)
	}

	for _, line := range n {
		for _, text := range line {
			assert.NotEmpty(t, text, "items in csv")
		}
	}

	assert.Equal(t, 5, len(n), "read csv from url")
	assert.Equal(t, 4, len(n[0]), "read csv from url")
}
