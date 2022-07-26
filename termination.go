package iter

// All returns whether all elements of the Iterator match the provided predicate.
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

// Any returns whether any elements of the Iterator match the provided predicate.
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

// Each performs an action for each element of the Iterator.
func Each[E any](it Iterator[E], fn func(v E)) {
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		fn(v)
	}
}

// Reduce performs a reduction on the elements of the Iterator,
// using an initial accumulator value and an accumulation func.
//
// If the iterator has none elements, the function returns the initial accumulator value and false.
// Or it returns the accumulation result and true.
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

// Last returns the last element of the Iterator.
//
// If the Iterator has none elements, this function returns the zero value and false.
// Or it returns the last element and true.
func Last[E any](it Iterator[E]) (v E, ok bool) {
	return Reduce(it, v, func(accum, e E) E {
		return e
	})
}

// Count returns the elements count of the Iterator.
func Count[E any](it Iterator[E]) uint64 {
	r, _ := Reduce(it, uint64(0), func(accum uint64, e E) uint64 {
		return accum + 1
	})
	return r
}

// Sortable are types that values of that type can be sorted by value.
type Sortable interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
		~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~float64 | ~float32 | ~string
}

// MaxBy returns the maximum element of an Iterator.
// The Sortable value of each element will be calculated via the provided value func.
//
// If the Iterator has none elements, this function returns the zero value and false.
// Or it returns the maximum element and true.
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

// Max returns the maximum element of an Iterator of Sortable elements.
//
// If the Iterator has none elements, this function returns the zero value and false.
// Or it returns the maximum element and true.
func Max[E Sortable](it Iterator[E]) (v E, ok bool) {
	return MaxBy(it, func(e E) E {
		return e
	})
}

// MinBy returns the minimum element of an Iterator.
// The Sortable value of each element will be calculated via the provided value func.
//
// If the Iterator has none elements, this function returns the zero value and false.
// Or it returns the minimum element and true.
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

// Min returns the minimum element of an Iterator of Sortable elements.
//
// If the Iterator has none elements, this function returns the zero value and false.
// Or it returns the minimum element and true.
func Min[E Sortable](it Iterator[E]) (v E, ok bool) {
	return MinBy(it, func(e E) E {
		return e
	})
}

// Arithmetical are types that can be used in arithmetic operations.
// Note that complex types are excluded.
type Arithmetical interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
		~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~float64 | ~float32
}

// Mean returns the mean value of an Iterator of Arithmetical elements.
//
// If the Iterator has none elements, this function returns the zero value and false.
// Or it returns the mean value and true.
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

// Sum returns the sum value of an Iterator of Arithmetical elements.
//
// If the Iterator has none elements, this function returns the zero value and false.
// Or it returns the sum value and true.
func Sum[E Arithmetical](it Iterator[E]) (v E, ok bool) {
	return Reduce(it, v, func(accum, e E) E { return accum + e })
}
