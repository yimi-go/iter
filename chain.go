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
