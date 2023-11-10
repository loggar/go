package constraint

type Example interface {
	int | int16 | int32 | int64
}

func SampleWithExample[T Example](data T) bool {
	return true
}
