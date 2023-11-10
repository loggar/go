package constraint

type AllowedData interface {
	// constraints.Ordered is a constraint that allows any ordered types which use comparison operators, such as ≤=≥===.
	Ordered
}

type Filter[T AllowedData] map[T]bool

func (r Filter[T]) add(datum T) {
	r[datum] = true
}

func (r Filter[T]) has(datum T) bool {
	_, ok := r[datum]
	return ok
}

func FindDuplicate[T AllowedData](data []T) bool {
	inArray := Filter[T]{}
	for _, datum := range data {
		if inArray.has(datum) {
			return true
		}
		inArray.add(datum)
	}
	return false
}
