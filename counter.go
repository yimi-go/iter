package iter

import "fmt"

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
