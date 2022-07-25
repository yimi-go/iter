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

func TestCount(t *testing.T) {
	assert.Equal(t, uint64(9), Count(CountUntil(0, 1, 9)))
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
