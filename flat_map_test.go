package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	ss := []string{"abc", "def"}
	it := Slice(ss)
	fm := FlatMap(it, func(t string) Iterator[rune] {
		return Slice([]rune(t))
	})
	for i := 'a'; i <= 'f'; i++ {
		v, ok := fm.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	_, ok := fm.Next()
	assert.False(t, ok)
}
