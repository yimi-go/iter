package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFunc(t *testing.T) {
	count := 0
	it := FromFunc(func() (v int, ok bool) {
		if count >= 10 {
			return
		}
		v, ok = count, true
		count++
		return
	})
	for i := 0; i < 10; i++ {
		v, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	_, ok := it.Next()
	assert.False(t, ok)
}
