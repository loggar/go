package customerrors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestDumbGetter creates a server which always returns StatusTeapot and then we use its URL as the argument to DumbGetter so we can see it handles non 200 responses correctly.
func TestDumbGetter(t *testing.T) {
	t.Run("when you dont get a 200 you get a status error", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusTeapot)
		}))
		defer svr.Close()

		_, err := DumbGetter(svr.URL)

		if err == nil {
			t.Fatal("expected an error")
		}

		want := fmt.Sprintf("did not get 200 from %s, got %d", svr.URL, http.StatusTeapot)
		got := err.Error()

		if got != want {
			t.Errorf(`got "%v", want "%v"`, got, want)
		}
	})

}

/*
Problems with this way of testing

This book tries to emphasise listen to your tests and this test doesn't feel good:

We're constructing the same string as production code does to test it
It's annoying to read and write
Is the exact error message string what we're actually concerned with ?

What does this tell us? The ergonomics of our test would be reflected on another bit of code trying to use our code.

How does a user of our code react to the specific kind of errors we return? The best they can do is look at the error string which is extremely error prone and horrible to write.
*/
