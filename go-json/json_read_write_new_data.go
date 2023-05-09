package main

import (
	"fmt"
	"os"
	"path/filepath"

	json_sample "github.com/loggar/go/go-json/json-read-write"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	dataFilePath := filepath.Join(cwd, "./go-json/json-read-write", "data.ex.data.json")
	jsonTarget1 := filepath.Join(cwd, "./go-json/json-read-write", "data.ex.target.json")
	outFilePath := filepath.Join(cwd, "./_out/json", "data.out.1.json")

	dataArray, err := json_sample.ReadArray(dataFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	data1, err := json_sample.ReadStruct1(jsonTarget1)
	if err != nil {
		fmt.Println(err)
		return
	}

	topArray := data1.TopArrayElement

	for i, item := range topArray {
		newPathValue := json_sample.FindFirstPathMatch(dataArray, item.Path)
		fmt.Printf("topArrayItem %d path=%s, newPath=%s, fromData=%s\n", i, item.Path, item.NewPath, newPathValue)
		if item.NewPath != newPathValue {
			data1.TopArrayElement[i].NewPath = newPathValue
		}
		for j, subItem := range item.SubArr {
			newSubPathValue := json_sample.FindFirstPathMatch(dataArray, subItem.Path)
			fmt.Printf("subArrayItem %d,%d path=%s, newPath=%s, fromData=%s\n", i, j, subItem.Path, subItem.NewPath, newSubPathValue)
			if subItem.NewPath != newSubPathValue {
				data1.TopArrayElement[i].SubArr[j].NewPath = newSubPathValue
			}
		}
	}

	jsonStr, err := json_sample.ConvertDataToJsonString(data1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("NEW JSON - %s:\n%s\n", jsonTarget1, jsonStr)

	err = json_sample.WriteDataToJsonFile(data1, outFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

}
