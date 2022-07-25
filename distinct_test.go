package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	it := CountUntil(0, 1, 8)
	it = Chain(it, CountUntil(2, 1, 9))
	it = Distinct(it)
	for i := 0; i < 9; i++ {
		v, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	_, ok := it.Next()
	assert.False(t, ok)
}
