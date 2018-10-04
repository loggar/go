package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var gopath = os.Getenv("GOPATH")
var rootDir = "/src/github.com/loggar/go"
var saveDir = "/_data/practice-go-web/gowiki/"

// Page is an wiki page structure
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	title := p.Title
	filename := path.Join(gopath, rootDir, saveDir, title+".txt")
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
