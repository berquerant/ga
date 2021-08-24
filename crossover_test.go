package ga_test

import (
	"testing"

	"github.com/berquerant/ga"
	"github.com/stretchr/testify/assert"
)

func TestUniformCrossover(t *testing.T) {
	t.Run("need random", func(t *testing.T) {
		_, err := ga.NewUniformCrossoverWithRandom(nil)
		assert.NotNil(t, err)
	})
	t.Run("need parents", func(t *testing.T) {
		s, err := ga.NewUniformCrossoverWithRandom(&mockRandom{})
		assert.Nil(t, err)
		_, err = s.Cross(nil, nil)
		assert.NotNil(t, err)
	})
	t.Run("parents must have the same length", func(t *testing.T) {
		s, err := ga.NewUniformCrossoverWithRandom(&mockRandom{})
		assert.Nil(t, err)
		_, err = s.Cross(ga.NewGene(3), ga.NewGene(2))
		assert.NotNil(t, err)
	})
	t.Run("not enough gene length", func(t *testing.T) {
		s, err := ga.NewUniformCrossoverWithRandom(&mockRandom{})
		assert.Nil(t, err)
		_, err = s.Cross(ga.NewGene(0), ga.NewGene(0))
		assert.NotNil(t, err)
	})
	t.Run("cross", func(t *testing.T) {
		// 11010
		a := ga.NewGene(5)
		a.Reverse(0)
		a.Reverse(1)
		a.Reverse(3)
		// 00101
		b := ga.NewGene(5)
		b.Reverse(2)
		b.Reverse(4)
		s, err := ga.NewUniformCrossoverWithRandom(&mockRandom{
			i: []int{0, 0, 1, 0, 1},
		})
		assert.Nil(t, err)
		g, err := s.Cross(a, b)
		assert.Nil(t, err)
		assert.Equal(t, 5, g.Len())
		for i := 0; i < g.Len(); i++ {
			assert.True(t, g.Get(i))
		}
	})
}

func TestTwoPointCrossover(t *testing.T) {
	t.Run("need random", func(t *testing.T) {
		_, err := ga.NewTwoPointCrossoverWithRandom(nil)
		assert.NotNil(t, err)
	})
	t.Run("need parents", func(t *testing.T) {
		s, err := ga.NewTwoPointCrossoverWithRandom(&mockRandom{})
		assert.Nil(t, err)
		_, err = s.Cross(nil, nil)
		assert.NotNil(t, err)
	})
	t.Run("parents must have the same length", func(t *testing.T) {
		s, err := ga.NewTwoPointCrossoverWithRandom(&mockRandom{})
		assert.Nil(t, err)
		_, err = s.Cross(ga.NewGene(3), ga.NewGene(2))
		assert.NotNil(t, err)
	})
	t.Run("not enough gene length", func(t *testing.T) {
		s, err := ga.NewTwoPointCrossoverWithRandom(&mockRandom{})
		assert.Nil(t, err)
		_, err = s.Cross(ga.NewGene(2), ga.NewGene(2))
		assert.NotNil(t, err)
	})
	t.Run("cross", func(t *testing.T) {
		// 11010
		a := ga.NewGene(5)
		a.Reverse(0)
		a.Reverse(1)
		a.Reverse(3)
		// 00101
		b := ga.NewGene(5)
		b.Reverse(2)
		b.Reverse(4)
		s, err := ga.NewTwoPointCrossoverWithRandom(&mockRandom{
			i: []int{1, 2},
		})
		assert.Nil(t, err)
		g, err := s.Cross(a, b)
		assert.Nil(t, err)
		assert.Equal(t, 5, g.Len())
		idx := map[int]bool{
			0: true,
			1: true,
			2: true,
			3: true,
		}
		for i := 0; i < g.Len(); i++ {
			assert.Equal(t, idx[i], g.Get(i))
		}
	})
}

func TestOnePointCrossover(t *testing.T) {
	t.Run("need random", func(t *testing.T) {
		_, err := ga.NewOnePointCrossoverWithRandom(nil)
		assert.NotNil(t, err)
	})
	t.Run("need parents", func(t *testing.T) {
		s, err := ga.NewOnePointCrossoverWithRandom(&mockRandom{})
		assert.Nil(t, err)
		_, err = s.Cross(nil, nil)
		assert.NotNil(t, err)
	})
	t.Run("parents must have the same length", func(t *testing.T) {
		s, err := ga.NewOnePointCrossoverWithRandom(&mockRandom{})
		assert.Nil(t, err)
		_, err = s.Cross(ga.NewGene(1), ga.NewGene(2))
		assert.NotNil(t, err)
	})
	t.Run("not enough gene length", func(t *testing.T) {
		s, err := ga.NewOnePointCrossoverWithRandom(&mockRandom{})
		assert.Nil(t, err)
		_, err = s.Cross(ga.NewGene(1), ga.NewGene(1))
		assert.NotNil(t, err)
	})
	// 11010
	a := ga.NewGene(5)
	a.Reverse(0)
	a.Reverse(1)
	a.Reverse(3)
	// 00101
	b := ga.NewGene(5)
	b.Reverse(2)
	b.Reverse(4)
	for _, tc := range []struct {
		name string
		p    int
		want []int
	}{
		{
			name: "0",
			p:    0,
			want: []int{0, 2, 4},
		},
		{
			name: "1",
			p:    1,
			want: []int{0, 1, 2, 4},
		},
		{
			name: "2",
			p:    2,
			want: []int{0, 1, 4},
		},
		{
			name: "3",
			p:    3,
			want: []int{0, 1, 3, 4},
		},
	} {
		t.Run(string(tc.name), func(t *testing.T) {
			s, err := ga.NewOnePointCrossoverWithRandom(&mockRandom{
				i: []int{tc.p},
			})
			g, err := s.Cross(a, b)
			assert.Nil(t, err)
			assert.Equal(t, 5, g.Len())
			idx := map[int]bool{}
			for _, i := range tc.want {
				idx[i] = true
			}
			for i := 0; i < g.Len(); i++ {
				assert.Equal(t, idx[i], g.Get(i))
			}
		})
	}
}
