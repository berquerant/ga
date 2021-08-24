package internal_test

import (
	"testing"

	"github.com/berquerant/ga/internal"
	"github.com/stretchr/testify/assert"
)

func TestIntSet(t *testing.T) {
	s := internal.NewIntSet()
	assert.Equal(t, 0, s.Size())
	assert.Empty(t, s.Slice())
	s.Set(0)
	assert.Equal(t, 1, s.Size())
	assert.True(t, s.Get(0))
	assert.False(t, s.Get(1))
	{
		v := s.Slice()
		assert.Equal(t, 1, len(v))
		assert.Equal(t, 0, v[0])
	}
	s.Set(0)
	assert.Equal(t, 1, s.Size())
	assert.True(t, s.Get(0))
	assert.False(t, s.Get(1))
	{
		v := s.Slice()
		assert.Equal(t, 1, len(v))
		assert.Equal(t, 0, v[0])
	}
}
