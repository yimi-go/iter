package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	a, b := Slice([]int{1, 2, 3}), Slice([]int{4, 5, 6})
	it := Chain(a, b)
	for i := 0; i < 6; i++ {
		v, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, i+1, v)
	}
	_, ok := it.Next()
	assert.False(t, ok)
}
