package ga

import (
	"context"

	"github.com/berquerant/ga/internal"
)

// Generation is a set of individuals.
type Generation interface {
	// Len returns the size of the individuals.
	Len() int
	// Individuals returns the copy of the individuals.
	Individuals() []Individual
	// Next calculates the next generation.
	Next(ctx context.Context) (Generation, error)
}

type generation struct {
	selector    IndividualsSelector
	evaluator   IndividualsEvaluator
	individuals []Individual
}

func (s *generation) Len() int { return len(s.individuals) }
func (s *generation) Individuals() []Individual {
	v := make([]Individual, len(s.individuals))
	for i, x := range s.individuals {
		v[i] = x.Clone()
	}
	return v
}

func (s *generation) Next(ctx context.Context) (Generation, error) {
	if internal.IsDone(ctx) {
		return nil, NewGAError(ctx.Err(), CannotMakeNextGeneration, "timeout")
	}
	if err := s.evaluator.Eval(ctx, s.individuals); err != nil {
		return nil, NewGAError(err, CannotMakeNextGeneration, "eval error")
	}
	individuals, err := s.selector.Select(s.individuals)
	if err != nil {
		return nil, NewGAError(err, CannotMakeNextGeneration, "select error")
	}
	return &generation{
		selector:    s.selector,
		evaluator:   s.evaluator,
		individuals: individuals,
	}, nil
}
