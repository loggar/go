package basics

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	app := &App{Message: "hello"}
	s := httptest.NewServer(app) // http provides its own test objects
	defer s.Close()

	resp, err := http.Get(s.URL) // test from the outside
	if err != nil {
		t.Error(err)
	}
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Error(err)
	} else if string(body) != "hello" {
		t.Error("expected", "hello", "got", body)
	}
}
