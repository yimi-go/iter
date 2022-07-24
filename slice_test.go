package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	it := Slice(s)
	for i := 0; i < 5; i++ {
		v, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, i+1, v)
	}
	_, ok := it.Next()
	assert.False(t, ok)
}
