package iter

import (
	"strconv"
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

func TestFlatMap(t *testing.T) {
	strIter := Slice([]string{"abc", "def"})
	runeIter := FlatMap(strIter, func(t string) Iterator[rune] {
		return Slice([]rune(t))
	})
	for i := 'a'; i <= 'f'; i++ {
		v, ok := runeIter.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	_, ok := runeIter.Next()
	assert.False(t, ok)
}

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

func TestMap(t *testing.T) {
	intIter := CountUntil(0, 1, 9)
	strIter := Map(intIter, func(t int) string {
		return strconv.Itoa(t)
	})
	for i := 0; i < 9; i++ {
		v, ok := strIter.Next()
		assert.True(t, ok)
		assert.Equal(t, strconv.Itoa(i), v)
	}
	_, ok := strIter.Next()
	assert.False(t, ok)
}
