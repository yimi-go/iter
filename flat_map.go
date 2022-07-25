package iter

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
