package basics

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func fakeApp(msg string) *httptest.Server {
	app := &App{Message: msg}
	return httptest.NewServer(app)
}

func get(t *testing.T, s *httptest.Server, path string) string {
	resp, err := http.Get(s.URL + path)
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	return string(body)
}

/*
In this case, I wrote two helpers to make this easier on me. This is my fake app, which is going to run an HTTP test server. I’ll initialize the app with whatever message I passed to the server, and I’ll return an HTTP test version of the server.

The second one I wrote is get. You give it your testing instance and a server and a path, and it makes a get request to that server and path, checks for error. It reads the body and checks for an error and returns the string version of that body. So this is going to do the whole get for me and send it back for any path on my server.
*/

func TestHelpedHello(t *testing.T) {
	s := fakeApp("hello")
	defer s.Close()
	if body := get(t, s, "/"); body != "hello" {
		t.Error("expected", "hello", "got", body)
	}
}
