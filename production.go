package iter

import "fmt"

// Countable are types that can be operated by '+' and their values are discrete.
type Countable interface {
	~int64 | ~int16 | ~int8 | ~int |
		~uint64 | ~uint32 | ~uint16 | ~uint8 | uint |
		~rune
}

type counter[E Countable] struct {
	step  E
	cur   E
	until func(v E) bool
}

func (s *counter[E]) Next() (v E, ok bool) {
	if s.until(s.cur) {
		return
	}
	v, ok = s.cur, true
	s.cur += s.step
	return
}

// CountTo creates an Iterator that generates values in range of specified
// start, step and end (included).
func CountTo[E Countable](from, step, to E) Iterator[E] {
	var zero E
	if step == zero {
		panic(fmt.Errorf("iter: counter: step can not be zero"))
	}
	res := &counter[E]{
		step: step,
		cur:  from,
	}
	if step > zero {
		res.until = func(v E) bool {
			return v > to
		}
	} else {
		res.until = func(v E) bool {
			return v < to
		}
	}
	return res
}

// CountUntil creates an Iterator that generates values in range of specified
// start, step and end (excluded).
func CountUntil[E Countable](from, step, notTo E) Iterator[E] {
	var zero E
	if step == zero {
		panic(fmt.Errorf("iter: counter: step can not be zero"))
	}
	res := &counter[E]{
		step: step,
		cur:  from,
	}
	if step > zero {
		res.until = func(v E) bool {
			return v >= notTo
		}
	} else {
		res.until = func(v E) bool {
			return v <= notTo
		}
	}
	return res
}

type fromFunc[E any] func() (E, bool)

func (f fromFunc[E]) Next() (E, bool) {
	return f()
}

// FromFunc creates an Iterator that produces value from specified closure.
func FromFunc[E any](fn func() (E, bool)) Iterator[E] {
	return fromFunc[E](fn)
}

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

// Slice creates an Iterator of the slice.
func Slice[E any](s []E) Iterator[E] {
	return &slice[E]{
		cur:   0,
		slice: s,
	}
}
