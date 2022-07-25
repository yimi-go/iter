package iter

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
