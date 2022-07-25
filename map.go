package iter

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
