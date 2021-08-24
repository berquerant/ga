package ga_test

import (
	"fmt"
	"testing"

	"github.com/berquerant/ga"
	"github.com/stretchr/testify/assert"
)

func ExampleGene() {
	g := ga.NewGene(4)
	fmt.Println(g)
	fmt.Println(g.Get(1))
	g.Reverse(1)
	fmt.Println(g)
	// Output:
	// 0000
	// false
	// 0100
}

func TestGene(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		x := ga.NewGene(3)
		assert.Equal(t, 3, x.Len())
		for i := -1; i < x.Len()+1; i++ {
			assert.False(t, x.Get(i))
		}
	})

	t.Run("set", func(t *testing.T) {
		x := ga.NewGene(3)
		x.Set(1, true)
		x.Set(2, false)
		x.Set(3, true) // out of range
		for i, w := range []bool{false, true, false} {
			assert.Equal(t, w, x.Get(i))
		}
	})

	t.Run("reverse", func(t *testing.T) {
		x := ga.NewGene(3)
		x.Reverse(1)
		x.Reverse(2)
		x.Reverse(3) // out of range
		x.Reverse(2) // second call
		for i, w := range []bool{false, true, false} {
			assert.Equal(t, w, x.Get(i))
		}
	})

	t.Run("clone", func(t *testing.T) {
		x := ga.NewGene(3)
		x.Reverse(2)
		y := x.Clone()
		assert.Equal(t, x.Len(), y.Len())
		for i := 0; i < x.Len(); i++ {
			assert.Equal(t, x.Get(i), y.Get(i))
		}
		y.Reverse(2)
		assert.True(t, x.Get(2))
		assert.False(t, y.Get(2))
	})
}
