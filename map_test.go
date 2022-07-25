package iter

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	it := CountUntil(0, 1, 9)
	m := Map(it, func(t int) string {
		return strconv.Itoa(t)
	})
	for i := 0; i < 9; i++ {
		v, ok := m.Next()
		assert.True(t, ok)
		assert.Equal(t, strconv.Itoa(i), v)
	}
	_, ok := m.Next()
	assert.False(t, ok)
}
