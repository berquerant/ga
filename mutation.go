package ga

import (
	"github.com/berquerant/ga/internal"
)

// Mutation is an interface for mutating genes.
type Mutation interface {
	// Mutate reverses some sections of the gene.
	Mutate(gene Gene) Gene
}

type mutation struct {
	n    int
	rand Random
}

// NewMutation returns new mutation.
// It reverses n sections of a gene.
func NewMutation(n int) (Mutation, error) {
	return NewMutationWithRandom(n, NewRandom())
}

// NewMutationWithRandom returns a new mutation that uses rand to select sections of the gene to reverse.
func NewMutationWithRandom(n int, rand Random) (Mutation, error) {
	if n < 0 {
		return nil, NewGAError(nil, InvalidArgument, "mutation number cannot be negative")
	}
	if rand == nil {
		return nil, NewGAError(nil, InvalidArgument, "need rand")
	}
	return &mutation{
		n:    n,
		rand: rand,
	}, nil
}

func (s *mutation) Mutate(gene Gene) Gene {
	g := gene.Clone()
	if g.Len() <= s.n {
		for i := 0; i < g.Len(); i++ {
			g.Reverse(i)
		}
		return g
	}
	set := internal.NewIntSet()
	for set.Size() < s.n {
		set.Set(s.rand.Int(0, g.Len()))
	}
	for _, i := range set.Slice() {
		g.Reverse(i)
	}
	return g
}
