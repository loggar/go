# 5 Advanced Go Testing Techniques

* https://hackernoon.com/5-advanced-go-testing-techniques-7199b686b6c1

Go has a robust built-in testing library. If you write Go, you already know this. In this post we will discuss a handful of strategies to level up your Go testing. We have learned from experience on our large Go codebase that these strategies work to save time and effort maintaining the code.

## Use test suites

If you take away just one thing from this post, it should be: use test suites. For those unfamiliar with the pattern, suite testing is the process of developing a test against a common interface which can be used against multiple implementations of that interface. Below, you can see how we can pass in multiple different `Thinger` implementations and have them run against the same tests.

```go
type Thinger interface {
    DoThing(input string) (Result, error)
}
// Suite tests all the functionality that Thingers should implement
func Suite(t *testing.T, impl Thinger) {
    res, _ := impl.DoThing("thing")
    if res != expected {
        t.Fail("unexpected result")
    }
}
// TestOne tests the first implementation of Thinger
func TestOne(t *testing.T) {
    one := one.NewOne()
    Suite(t, one)
}
// TestOne tests another implementation of Thinger
func TestTwo(t *testing.T) {
    two := two.NewTwo()
    Suite(t, two)
}
```

Fortunate readers may have worked with codebases which use this technique. Frequently used in plugin-based systems, tests which are written against the interface are usable by all implementations of that interface to determine if the behavior requirements are met.

Using this strategy will save potentially hours, days, or perhaps even enough time to solve the P versus NP problem. Also when swapping two underlying systems you don’t have to write (many) additional tests and it provides confidence that doing so won’t break your application. This implicitly requires that you create an interface defining the surface area of that which you’re testing. Using dependency injection, you set up the suite from your package passing in the implementation for the package.

A complete example is provided [here](https://github.com/segmentio/testdemo) . While this example is fairly contrived, you can imagine one implementation being a remote database whilst the other being an in-memory database.

Another fantastic example of this in the standard library is the `golang.org/x/net/nettest` package. It provides the means to verify a `net.Conn` satisfies its interface.

## Avoid interface pollution

We can’t talk about testing in Go without talking about interfaces.

Interfaces are important in the context of testing because they are the most powerful tool in our test arsenal, so it’s important to use them correctly. Packages frequently export an interface for consumers to use which in turn leads to either: a. consumers implementing their own mock of the package implementation or b. the package exporting their own mock.

> The bigger the interface, the weaker the abstraction.
— Rob Pike, Go Proverbs

Interfaces should be carefully considered before exporting them. Developers are often tempted to export interfaces as a way for consumers to mock out their behavior. Instead, document what interfaces your structs satisfy such that you don’t create a hard dependency between the consumer package and your own. A great example of this is the errors package.

When we have an interface in our program which we don’t want to export can use an [internal/ package subtree](https://golang.org/doc/go1.4#internalpackages)  to keep it scoped to the package. By doing this, we remove the concern that other consumers might depend on it and therefore can be flexible in the evolution of the interface as new requirements present themselves. We usually create interfaces around external dependencies and use dependency injection so we can run tests locally.

This enables the consumer to implement small interfaces of their own, only wrapping the consumed surface of the library for their own testing.

## Don’t export concurrency primitives

Go offers easy-to-use concurrency primitives which can also sometimes lead to their overuse. We’re primarily concerned about channels and the `sync` package. It is sometimes tempting to export a channel from your package for consumers to use. Additionally, it’s a common mistake to embed `sync.Mutex` without making it private. As with anything, this isn’t always bad but it does pose challenges when testing your program.

f you’re exporting channels, you’re exposing the consumer of the package to additional complexity they shouldn’t care about. As soon as a channel is exported from a package, you’re opening up challenges in testing for the one consuming that channel. To test well, the consumer will need to know:

* When data is finished being sent on the channel
* Whether or not there were any errors receiving the data
* How does the package clean up the channel after completion, if at all?
* How can I wrap an interface around the package API so I don’t have to call it directly?

Consider an example of reading a queue. Here’s an example library which reads from the queue and exposes a channel for the consumer to read from.

```go
type Reader struct {...}
func (r *Reader) ReadChan() <-chan Msg {...}
```

Now a user of your library wants to implement a test for their consumer:

```go
func TestConsumer(t testing.T) {
    cons := &Consumer{
        r: libqueue.NewReader(),
    }
    for msg := range cons.r.ReadChan() {
        // Test thing.
    }
}
```

The user might then decide that dependency injection is a good idea here and write their own messages along the channel:

```go
func TestConsumer(t testing.T, q queueIface) {
    cons := &Consumer{
        r: q,
    }
    for msg := range cons.r.ReadChan() {
        // Test thing.
    }
}
```

But wait, what about errors?

```go
func TestConsumer(t testing.T, q queueIface) {
    cons := &Consumer{
        r: q,
    }
    for {
        select {
        case msg := <-cons.r.ReadChan():
            // Test thing.
        case err := <-cons.r.ErrChan():
            // What caused this again?
        }
    }
}
```

Now, how do we generate events to actually write into this mock which sufficiently replicate the behavior of the actual library we’re using? If the library simply wrote a synchronous API, then we could add all this concurrency in our client code and it becomes much simpler to test.

```go
func TestConsumer(t testing.T, q queueIface) {
    cons := &Consumer{
        r: q,
    }
    msg, err := cons.r.ReadMsg()
    // handle err, test thing
}
```

When in doubt, remember that it’s always easy to add concurrency in a consuming package and hard/impossible to remove once exported from a library. Finally, don’t forget to mention in package documentation whether or not a struct/package is safe for concurrent access by multiple goroutines!

Sometimes, it’s still desirable or necessary to export a channel. To alleviate some of the problems above, you can expose channels through accessors instead of directly and force them to be read-only (←chan) or write-only (chan←) channels in their declaration.

## Use net/http/httptest

`httptest` allows you to exercise your `http.Handler` code without actually spinning up a server or binding to a port. This speeds up tests and allows them to run in parallel with less effort.

Here’s an example of the same test implemented using both methods. It doesn’t look like much, but it saves you a considerable amount of code and resources.

```go
func TestServe(t *testing.T) {
    // The method to use if you want to practice typing
    s := &http.Server{
        Handler: http.HandlerFunc(ServeHTTP),
    }
    // Pick port automatically for parallel tests and to avoid conflicts
    l, err := net.Listen("tcp", ":0")
    if err != nil {
        t.Fatal(err)
    }
    defer l.Close()
    go s.Serve(l)
    res, err := http.Get("http://" + l.Addr().String() + "/?sloths=arecool")
    if err != nil {
        log.Fatal(err)
    }
    greeting, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(greeting))
}
func TestServeMemory(t *testing.T) {
    // Less verbose and more flexible way
    req := httptest.NewRequest("GET", "http://example.com/?sloths=arecool", nil)
    w := httptest.NewRecorder()
    ServeHTTP(w, req)
    greeting, err := ioutil.ReadAll(w.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(greeting))
}
```

Perhaps the biggest thing using `httptest` buys you is the ability to compartmentalize your test to just the function you want to test. No routers, middleware or any other side-effects that come from setting up servers, services, handler factories, handler factory factories or whatever abominations are thrown upon you by ideas your former self thought were good.

## Use a separate _test package

Most tests in the ecosystem are created in files pkg_test.go but still live in the same package: package pkg. A separate test package is a package you create in a new file, foo_test.go, inside the directory of the package you wish to test, foo/, with the declaration package foo_test . From there, you can import github.com/example/foo and other dependencies. This feature enables a number of things. It is the recommended workaround for cyclic dependencies in tests, it prevents brittle tests, and it allows the developer to feel what it’s like to consume their own package. If your package is hard to use, it will likely also be hard to test using this method.

This strategy prevents brittle tests by restricting access to private variables. In particular, if your tests break and you’re using a separate test packages it’s almost guaranteed that a client using the feature which broke in tests will also break when called.

Finally, this aids in avoiding import cycles in tests. Most packages likely depend on other packages you wrote aside from those being tested, so you’ll eventually run into a situation where an import cycle happens as a natural consequence. An external package sits above both packages in the package hierarchy. To take an example from The Go Programming Language (Chp. 11 Sec 2.4), net/urlimplements a URL parser which net/http imports for use. However, net/urlwould like to test using a real use case by importing net/http. Thus net/url_testwas born.

Now that you’re using a separate test package you might require access to unexported entities in your package where they were previously accessible. Most people hit this first when testing something time based (e.g. time.Now gets mocked via a function). In this situation, we can use an additional file to expose them exclusively during testing since _test.go files are excluded from regular builds.
