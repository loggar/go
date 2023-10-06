package customerrors

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
With TDD we have the benefit of getting into the mindset of:

"How would I want to use this code?"

What we could do for DumbGetter is provide a way for users to use the type system to understand what kind of error has happened.

What if DumbGetter could return us something like

type BadStatusError struct {
    URL    string
    Status int
}

Rather than a magical string, we have actual data to work with.
*/

// BadStatusError can be used like an actual data to work with.
type BadStatusError struct {
	URL    string
	Status int
}

// We'll have to make BadStatusError implement the error interface.
func (b BadStatusError) Error() string {
	return fmt.Sprintf("did not get 200 from %s, got %d", b.URL, b.Status)
}

// DumbGetter2 now can return a custom error
func DumbGetter2(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("problem fetching from %s, %v", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return "", BadStatusError{URL: url, Status: res.StatusCode}
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body) // ignoring err for brevity

	return string(body), nil
}
