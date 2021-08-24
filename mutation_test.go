package ga_test

import (
	"testing"

	"github.com/berquerant/ga"
	"github.com/stretchr/testify/assert"
)

func TestMutation(t *testing.T) {
	t.Run("negative mutations", func(t *testing.T) {
		_, err := ga.NewMutationWithRandom(-1, nil)
		assert.NotNil(t, err)
	})
	t.Run("no rand", func(t *testing.T) {
		_, err := ga.NewMutationWithRandom(1, nil)
		assert.NotNil(t, err)
	})
	for _, tc := range []struct {
		name string
		n    int
		idx  []int
	}{
		{
			name: "no mutations",
		},
		{
			name: "1 mutation",
			n:    1,
			idx:  []int{1},
		},
		{
			name: "2 mutations",
			n:    2,
			idx:  []int{1, 2},
		},
		{
			name: "2 mutations with random number duplicated",
			n:    2,
			idx:  []int{1, 1, 2},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s, err := ga.NewMutationWithRandom(tc.n, &mockRandom{
				i: tc.idx,
			})
			assert.Nil(t, err)
			g := s.Mutate(ga.NewGene(5))
			idx := map[int]bool{}
			for _, i := range tc.idx {
				idx[i] = true
			}
			for i := 0; i < g.Len(); i++ {
				assert.Equal(t, idx[i], g.Get(i))
			}
		})
	}
}
