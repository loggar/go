package constraint

// Example of channel conversion using generics
func ToAnyChan[T any](in <-chan T) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)

		for s := range in {
			out <- s
		}
	}()
	return out
}
