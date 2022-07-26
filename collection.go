package iter

// ToSlice collects all elements of the Iterator to a new slice.
func ToSlice[E any](it Iterator[E]) []E {
	slice := make([]E, 0)
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		slice = append(slice, v)
	}
	return slice
}
