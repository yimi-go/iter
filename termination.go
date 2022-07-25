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

func Each[E any](it Iterator[E], fn func(v E)) {
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		fn(v)
	}
}

func Reduce[E any, R any](it Iterator[E], init R, fn func(accum R, v E) R) (r R, ok bool) {
	r = init
	for {
		next, nextOk := it.Next()
		if !nextOk {
			break
		}
		r = fn(r, next)
		ok = true
	}
	if !ok {
		r = init
	}
	return
}

func Last[E any](it Iterator[E]) (v E, ok bool) {
	return Reduce(it, v, func(accum, v E) E {
		return v
	})
}

func Count[E any](it Iterator[E]) uint64 {
	r, _ := Reduce(it, uint64(0), func(accum uint64, v E) uint64 {
		return accum + 1
	})
	return r
}
