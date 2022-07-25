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

func Reduce[E any, R any](it Iterator[E], init R, fn func(accum R, e E) R) (r R, ok bool) {
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
	return Reduce(it, v, func(accum, e E) E {
		return e
	})
}

func Count[E any](it Iterator[E]) uint64 {
	r, _ := Reduce(it, uint64(0), func(accum uint64, e E) uint64 {
		return accum + 1
	})
	return r
}

type Sortable interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
		~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~float64 | ~float32 | ~string
}

func MaxBy[E any, S Sortable](it Iterator[E], valueFn func(e E) S) (v E, found bool) {
	var p *E
	p, found = Reduce(it, p, func(accum *E, e E) *E {
		if accum == nil {
			return &e
		}
		if valueFn(*accum) < valueFn(e) {
			return &e
		}
		return accum
	})
	if found {
		v = *p
	}
	return
}

func Max[E Sortable](it Iterator[E]) (v E, ok bool) {
	return MaxBy(it, func(e E) E {
		return e
	})
}

func MinBy[E any, S Sortable](it Iterator[E], valueFn func(e E) S) (v E, found bool) {
	var p *E
	p, found = Reduce(it, p, func(accum *E, e E) *E {
		if accum == nil {
			return &e
		}
		if valueFn(*accum) > valueFn(e) {
			return &e
		}
		return accum
	})
	if found {
		v = *p
	}
	return
}

func Min[E Sortable](it Iterator[E]) (v E, ok bool) {
	return MinBy(it, func(e E) E {
		return e
	})
}

type Arithmetical interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
		~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~float64 | ~float32
}

func Mean[E Arithmetical](it Iterator[E]) (v E, ok bool) {
	count := uint64(0)
	sum, ok := Reduce(it, v, func(accum, e E) E {
		count++
		return accum + e
	})
	if !ok {
		return
	}
	v = sum / E(count)
	return
}
