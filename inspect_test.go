package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInspect(t *testing.T) {
	it := CountUntil(0, 1, 9)
	ch := make(chan int, 10)
	it = Inspect(it, func(v int) { ch <- v })
	for i := 0; i < 9; i++ {
		v, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	close(ch)
	_, ok := it.Next()
	assert.False(t, ok)
	for i := 0; i < 9; i++ {
		v, ok := <-ch
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	_, ok = <-ch
	assert.False(t, ok)
}
