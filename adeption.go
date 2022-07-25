package iter

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

func DistinctBy[E any, C comparable](it Iterator[E], fn func(v E) C) Iterator[E] {
	return &distinct[E, C]{
		iter:  it,
		set:   map[C]bool{},
		keyFn: fn,
	}
}

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

func Map[T any, R any](it Iterator[T], fn func(t T) R) Iterator[R] {
	return &mapping[T, R]{
		iter:  it,
		mapFn: fn,
	}
}
