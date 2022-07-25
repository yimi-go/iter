package iter

func All[E any](it Iterator[E], predicate func(v E) bool) bool {
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		if !predicate(v) {
			return false
		}
	}
	return true
}

func Any[E any](it Iterator[E], predicate func(v E) bool) bool {
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		if predicate(v) {
			return true
		}
	}
	return false
}

func Count[E any](it Iterator[E]) uint64 {
	count := uint64(0)
	for {
		_, ok := it.Next()
		if !ok {
			return count
		}
		count++
	}
}
