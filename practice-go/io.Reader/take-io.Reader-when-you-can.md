### Take io.Reader when you can

If you’re designing a package or utility (even if it’s an internal thing that nobody will ever see) rather than taking in strings or []byte slices, consider taking in an io.Reader if you can for data sources. Because suddenly, your code will work with every type that implements io.Reader.

So this:

```
func Reverse(s string) (string, error)
```

Could become:

```
func Reverse(r io.Reader) io.Reader
```

Then if someone wants to use it with a string, they can:

```
r = Reverse(strings.NewReader("Make me backwards"))
```

But they can also use it with a file:

```
f, err := os.Open("file.txt")
if err != nil {
  log.Fatalln(err)
}
r = Reverse(f)
```

Or a web request:

```
func handle(w http.ResponseWriter, r *http.Request) {
  rev := Reverse(r.Body)
  // etc...
}
```

Use io.Reader (and io.Writer) whenever you’re dealing with streams of data. And this goes for all single method interfaces from the standard library.