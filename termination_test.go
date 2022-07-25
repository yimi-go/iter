package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	assert.True(t, All(CountUntil(0, 1, 9), func(v int) bool {
		return v >= 0
	}))
	assert.False(t, All(CountUntil(0, 1, 9), func(v int) bool {
		return v%2 == 0
	}))
}

func TestAny(t *testing.T) {
	assert.True(t, Any(CountUntil(0, 1, 9), func(v int) bool {
		return v%2 == 0
	}))
	assert.False(t, Any(CountUntil(0, 1, 9), func(v int) bool {
		return v < 0
	}))
}

func TestEach(t *testing.T) {
	it := CountUntil(0, 1, 9)
	ch := make(chan int, 10)
	Each(it, func(v int) {
		ch <- v
	})
	close(ch)
	for i := 0; i < 9; i++ {
		v, ok := <-ch
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	_, ok := <-ch
	assert.False(t, ok)
}

func TestReduce(t *testing.T) {
	it := CountUntil(0, 1, 10)
	sumFn := func(accum, v int) int {
		return accum + v
	}
	sum, ok := Reduce(it, 0, sumFn)
	assert.True(t, ok)
	assert.Equal(t, 45, sum)
	_, ok = Reduce(it, 0, sumFn) // Reduce on an iterated iterator
	assert.False(t, ok)
}

func TestLast(t *testing.T) {
	it := CountUntil(0, 1, 9)
	v, ok := Last(it)
	assert.True(t, ok)
	assert.Equal(t, 8, v)
	_, ok = Last(it) // Last on an iterated iterator
	assert.False(t, ok)
}

func TestCount(t *testing.T) {
	assert.Equal(t, uint64(9), Count(CountUntil(0, 1, 9)))
}

func TestMax(t *testing.T) {
	it := Chain(CountUntil(5, 1, 9), CountTo(4, -1, 0))
	v, ok := Max(it)
	assert.True(t, ok)
	assert.Equal(t, 8, v)
	_, ok = Max(it) // Max on an iterated iterator
	assert.False(t, ok)
}
