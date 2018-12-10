package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// GetAstronauts fectch people from http request
func GetAstronauts(getWebRequest GetWebRequest) int {
	url := "https://loggar.github.io/note/sample-res/sample.astros.json"
	bodyBytes := getWebRequest.FetchBytes(url)
	peopleResult := people{}
	jsonErr := json.Unmarshal(bodyBytes, &peopleResult)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return peopleResult.Number
}

func main() {
	liveClient := LiveGetWebRequest{}
	number := GetAstronauts(liveClient)

	fmt.Println(number)
}
