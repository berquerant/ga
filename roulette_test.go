package ga_test

import (
	"fmt"
	"testing"

	"github.com/berquerant/ga"
	"github.com/stretchr/testify/assert"
)

func TestRoulette(t *testing.T) {
	t.Run("no weights", func(t *testing.T) {
		_, err := ga.NewRouletteWithRandom(nil, nil)
		assert.NotNil(t, err)
	})
	t.Run("no rand", func(t *testing.T) {
		_, err := ga.NewRouletteWithRandom([]float64{1}, nil)
		assert.NotNil(t, err)
	})
	t.Run("negative weight", func(t *testing.T) {
		_, err := ga.NewRouletteWithRandom([]float64{-1}, &mockRandom{})
		assert.NotNil(t, err)
	})

	// Spin roulette
	// 0-1 => 0
	// 1-3 => 1
	// 3-6 => 2
	weights := []float64{1, 2, 3}
	for _, tc := range []struct {
		p    float64
		want int
	}{
		{
			p:    0,
			want: 0,
		},
		{
			p:    0.5,
			want: 0,
		},
		{
			p:    0.99,
			want: 0,
		},
		{
			p:    1.1,
			want: 1,
		},
		{
			p:    3.1,
			want: 2,
		},
		{
			p:    6,
			want: 2,
		},
		{
			p:    1000,
			want: 2,
		},
	} {
		t.Run(fmt.Sprintf("gen %f got %d", tc.p, tc.want), func(t *testing.T) {
			r, err := ga.NewRouletteWithRandom(weights, &mockRandom{
				f: []float64{tc.p},
			})
			assert.Nil(t, err)
			assert.Equal(t, tc.want, r.Spin())
		})
	}
}
