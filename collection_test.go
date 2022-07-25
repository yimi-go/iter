package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	it := CountUntil(0, 1, 9)
	slice := ToSlice(it)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, slice)
}
