package iter

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
