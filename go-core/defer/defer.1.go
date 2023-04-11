package main

import (
	"fmt"
	"io"
	"os"
)

/*
Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns.
It's an unusual but effective way to deal with situations such as resources that must be released regardless of which path a function takes to return.
The canonical examples are unlocking a mutex or closing a file.
*/

// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}

/*
Deferring a call to a function such as Close has two advantages.
First, it guarantees that you will never forget to close the file, a mistake that's easy to make if you later edit the function to add a new return path.
Second, it means that the close sits near the open, which is much clearer than placing it at the end of the function.
*/

/*
Deferred functions are executed in LIFO order, so this code will cause 4 3 2 1 0 to be printed when the function returns.
A more plausible example is a simple way to trace function execution through the program.
We could write a couple of simple tracing routines like this:
*/

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
	trace("a")
	defer untrace("a")
	// do something....
}

/*
We can do better by exploiting the fact that arguments to deferred functions are evaluated when the defer executes.
The tracing routine can set up the argument to the untracing routine.
This example prints

entering: b
in b
entering: a
in a
leaving: a
leaving: b

*/

func trace2(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a2() {
	defer un(trace2("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace2("b"))
	fmt.Println("in b")
	a()
}

func main_1() {
	b()
}

/*
For programmers accustomed to block-level resource management from other languages, defer may seem peculiar,
but its most interesting and powerful applications come precisely from the fact that it's not block-based but function-based.
In the section on panic and recover we'll see another example of its possibilities.
*/
