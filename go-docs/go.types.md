# types

## pointer: assignment

Literal pointer value assignment (without generics)

```
// with a one-liner anonymous function
func(i int64) *int64 { return &i }(1)
```

```
// with a type-casting function
func toInt64Ptr(i int64) *int64 {
	return &i
}
```
