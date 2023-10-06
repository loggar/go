package customerrors

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestDumbGetter creates a server which always returns StatusTeapot and then we use its URL as the argument to DumbGetter so we can see it handles non 200 responses correctly.
func TestDumbGetter2(t *testing.T) {
	t.Run("when you dont get a 200 you get a status error", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusTeapot)
		}))
		defer svr.Close()

		_, err := DumbGetter2(svr.URL)

		if err == nil {
			t.Fatal("expected an error")
		}

		got, isStatusErr := err.(BadStatusError)

		if !isStatusErr {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}

		want := BadStatusError{URL: svr.URL, Status: http.StatusTeapot}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

/*
This change has had some real positive effects:

Our DumbGetter function has become simper, it's no longer concerned with the intricacies of an error string, it just creates a BadStatusError.
Our tests now reflect (and document) what a user of our code could do if they decided they wanted to do some more sophisticated error handling than just logging. Just do a type assertion and then you get easy access to the properties of the error.
It is still "just" an error, so if they choose to they can pass it up the call stack or log it like any other error.
*/

/*
If you find yourself testing for multiple error conditions dont fall in to the trap of comparing the error messages.

This leads to flaky and difficult to read/write tests and it reflects the difficulties the users of your code will have if they also need to start doing things differently depending on the kind of errors that have occurred.

Always make sure your tests reflect how you'd like to use your code, so in this respect consider creating an error type to encapsulate your kinds of errors. This makes handling different kinds of errors easier for users of your code and also makes writing your error handling code simpler and easier to read.
*/
