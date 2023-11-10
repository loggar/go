package constraint

// Example of channel conversions without generics
func SToIChan(in <-chan string) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)

		for s := range in {
			out <- s
		}
	}()
	return out
}

func IntToIChan(in <-chan int) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)

		for s := range in {
			out <- s
		}
	}()
	return out
}

type MyStruct struct {
	Data int
	Tags string
}

func StrToIChan(in <-chan MyStruct) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)

		for s := range in {
			out <- s
		}
	}()
	return out
}
