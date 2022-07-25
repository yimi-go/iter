package iter

type fromFunc[E any] func() (E, bool)

func (f fromFunc[E]) Next() (E, bool) {
	return f()
}

func FromFunc[E any](fn func() (E, bool)) Iterator[E] {
	return fromFunc[E](fn)
}
