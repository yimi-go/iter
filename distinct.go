package iter

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
