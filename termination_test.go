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
