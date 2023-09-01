package os_csv_files

import (
	"encoding/csv"
	"log"
	"net/http"
	"testing"
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
			log.Println(text)
		}
	}

	// expected := [][]string{{"10", "22", "2022"}, {"golang", "21", "read"}}

	// assert.Equal(t, expected, data, "read a scv by delimiter")
}
