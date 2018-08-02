package main

import "fmt"

func runF() {
	f()
	fmt.Println("Returned normally from f.")
}

/*
Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
*/

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

/*
If we remove the deferred function from f the panic is not recovered and reaches the top of the goroutine's call stack, terminating the program. This modified program will output:

Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
panic: 4

panic PC=0x2a9cd8
[stack trace omitted]
*/

/*
For a real-world example of panic and recover, see the json package from the Go standard library. It decodes JSON-encoded data with a set of recursive functions. When malformed JSON is encountered, the parser calls panic to unwind the stack to the top-level function call, which recovers from the panic and returns an appropriate error value (see the 'error' and 'unmarshal' methods of the decodeState type in decode.go).

The convention in the Go libraries is that even when a package uses panic internally, its external API still presents explicit error return values.

Other uses of defer (beyond the file.Close example given earlier) include releasing a mutex:

mu.Lock()
defer mu.Unlock()

printing a footer:

printHeader()
defer printFooter()

and more.
*/
