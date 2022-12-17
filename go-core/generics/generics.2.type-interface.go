package main

func main() {

}

type minTypes interface {
	~float64 | int
}

func min[T minTypes](a T, b T) T {
	if a < b {
		return a
	}

	return b
}
