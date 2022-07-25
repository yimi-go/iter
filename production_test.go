package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_counter_Next(t *testing.T) {
	type fields struct {
		step  int64
		cur   int64
		until func(v int64) bool
	}
	tests := []struct {
		name   string
		fields fields
		wantV  int64
		wantOk bool
	}{
		{
			name: "done",
			fields: fields{
				step: 1,
				cur:  1,
				until: func(v int64) bool {
					return v <= 1
				},
			},
			wantOk: false,
		},
		{
			name: "not_done",
			fields: fields{
				step: 1,
				cur:  1,
				until: func(v int64) bool {
					return v >= 10
				},
			},
			wantOk: true,
			wantV:  1,
		},
		{
			name: "infinity",
			fields: fields{
				step: 1,
				cur:  1,
				until: func(v int64) bool {
					return false
				},
			},
			wantOk: true,
			wantV:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &counter[int64]{
				step:  tt.fields.step,
				cur:   tt.fields.cur,
				until: tt.fields.until,
			}
			gotV, gotOk := s.Next()
			assert.Equal(t, tt.wantOk, gotOk)
			if !tt.wantOk {
				return
			}
			assert.Equal(t, tt.wantV, gotV)
		})
	}
}

func TestCountTo(t *testing.T) {
	t.Run("step_zero", func(t *testing.T) {
		defer func() {
			assert.NotNil(t, recover())
		}()
		_ = CountTo(0, 0, 9)
	})
	t.Run("count_up", func(t *testing.T) {
		it := CountTo(0, 1, 9)
		for i := 0; i <= 9; i++ {
			v, ok := it.Next()
			assert.True(t, ok)
			assert.Equal(t, i, v)
		}
		_, ok := it.Next()
		assert.False(t, ok)
	})
	t.Run("count_down", func(t *testing.T) {
		it := CountTo(9, -1, 0)
		for i := 0; i <= 9; i++ {
			v, ok := it.Next()
			assert.True(t, ok)
			assert.Equal(t, 9-i, v)
		}
		_, ok := it.Next()
		assert.False(t, ok)
	})
}

func TestCountUntil(t *testing.T) {
	t.Run("step_zero", func(t *testing.T) {
		defer func() {
			assert.NotNil(t, recover())
		}()
		_ = CountUntil(0, 0, 9)
	})
	t.Run("count_up", func(t *testing.T) {
		it := CountUntil(0, 1, 9)
		for i := 0; i < 9; i++ {
			v, ok := it.Next()
			assert.True(t, ok)
			assert.Equal(t, i, v)
		}
		_, ok := it.Next()
		assert.False(t, ok)
	})
	t.Run("count_down", func(t *testing.T) {
		it := CountUntil(9, -1, 0)
		for i := 0; i < 9; i++ {
			v, ok := it.Next()
			assert.True(t, ok)
			assert.Equal(t, 9-i, v)
		}
		_, ok := it.Next()
		assert.False(t, ok)
	})
}

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
