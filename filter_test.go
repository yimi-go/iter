package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	it := CountTo(0, 1, 9)
	it = Filter(it, func(v int) bool {
		return v%2 == 0
	})
	for i := 0; i <= 9; i += 2 {
		v, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	_, ok := it.Next()
	assert.False(t, ok)
}
