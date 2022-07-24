package iter

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
