package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

var gopath = os.Getenv("GOPATH")
var rootDir = "/src/github.com/loggar/go"
var saveDir = "/_out/go-webapp/gowiki/"

// Page is an wiki page structure
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	title := p.Title
	filename := path.Join(gopath, rootDir, saveDir, title+".txt")

	// Check if file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		dir := filepath.Dir(filename)

		// If directories do not exist, create them
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}

		data := []byte("Your file content goes here") // Replace with your data
		err = ioutil.WriteFile(filename, data, 0600)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := path.Join(gopath, rootDir, saveDir, title+".txt")
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage1", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage1")
	fmt.Println(string(p2.Body))
}
