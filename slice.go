package iter

type slice[E any] struct {
	cur   int
	slice []E
}

func (s *slice[E]) Next() (v E, ok bool) {
	if s.cur >= len(s.slice) {
		return
	}
	v, ok = s.slice[s.cur], true
	s.cur++
	return
}

func Slice[E any](s []E) Iterator[E] {
	return &slice[E]{
		cur:   0,
		slice: s,
	}
}
