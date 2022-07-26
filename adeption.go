package iter

// TODO refactor chain: wrapping slice of Iterators.
type chain[E any] struct {
	a, b Iterator[E]
}

func (c *chain[E]) Next() (E, bool) {
	v, ok := c.a.Next()
	if ok {
		return v, true
	}
	return c.b.Next()
}

// Chain connects iterators into one Iterator that iterates the chained iterators
// one by one.
func Chain[E any](a, b Iterator[E]) Iterator[E] {
	return &chain[E]{a, b}
}

type distinct[E any, C comparable] struct {
	iter  Iterator[E]
	set   map[C]bool
	keyFn func(v E) C
}

func (d *distinct[E, C]) Next() (v E, ok bool) {
	for {
		v, ok = d.iter.Next()
		if !ok {
			return
		}
		key := d.keyFn(v)
		if d.set[key] {
			continue
		}
		d.set[key] = true
		return
	}
}

// DistinctBy warps an Iterator and returns a new Iterator that only returns
// the first one of duplicated elements of key which calculated by the key func.
func DistinctBy[E any, C comparable](it Iterator[E], fn func(v E) C) Iterator[E] {
	return &distinct[E, C]{
		iter:  it,
		set:   map[C]bool{},
		keyFn: fn,
	}
}

// Distinct wraps an Iterator of comparable elements and returns a new Iterator
// that only returns the first one of duplicated elements.
func Distinct[E comparable](it Iterator[E]) Iterator[E] {
	return DistinctBy(it, func(v E) E {
		return v
	})
}

type filter[E any] struct {
	iter      Iterator[E]
	predicate func(v E) bool
}

func (f *filter[E]) Next() (v E, ok bool) {
	for {
		v, ok = f.iter.Next()
		if !ok {
			return
		}
		if f.predicate(v) {
			return
		}
	}
}

// Filter wraps an Iterator and returns a new Iterator that only returns
// elements that match the given predicate.
func Filter[E any](it Iterator[E], predicate func(v E) bool) Iterator[E] {
	return &filter[E]{
		iter:      it,
		predicate: predicate,
	}
}

type flatMap[T any, R any] struct {
	iter       Iterator[T]
	mapFn      func(t T) Iterator[R]
	mappedIter Iterator[R]
}

func (f *flatMap[T, R]) Next() (v R, ok bool) {
	for {
		if f.mappedIter != nil {
			if r, rOk := f.mappedIter.Next(); rOk {
				return r, true
			}
		}
		t, tOk := f.iter.Next()
		if !tOk {
			return
		}
		f.mappedIter = f.mapFn(t)
	}
}

// FlatMap wraps an Iterator and returns a new Iterator that returns the elements
// of the Iterators that produced by the mapping func one by one.
func FlatMap[T any, R any](it Iterator[T], mapFn func(t T) Iterator[R]) Iterator[R] {
	return &flatMap[T, R]{
		iter:  it,
		mapFn: mapFn,
	}
}

type inspect[E any] struct {
	iter Iterator[E]
	fn   func(v E)
}

func (i *inspect[E]) Next() (v E, ok bool) {
	v, ok = i.iter.Next()
	if ok {
		i.fn(v)
	}
	return
}

// Inspect wraps an Iterator and returns a new Iterator that performs the provided
// action on each element iterated before return it.
func Inspect[E any](it Iterator[E], fn func(v E)) Iterator[E] {
	return &inspect[E]{
		iter: it,
		fn:   fn,
	}
}

type mapping[T any, R any] struct {
	iter  Iterator[T]
	mapFn func(t T) R
}

func (m *mapping[T, R]) Next() (v R, ok bool) {
	t, ok := m.iter.Next()
	if !ok {
		return
	}
	v = m.mapFn(t)
	return
}

// Map wraps an Iterator and returns a new Iterator that transform each element
// to new value by provided mapping func, and returns the transformed element.
func Map[T any, R any](it Iterator[T], fn func(t T) R) Iterator[R] {
	return &mapping[T, R]{
		iter:  it,
		mapFn: fn,
	}
}
