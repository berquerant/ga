package ga

import (
	"context"
	"fmt"

	"github.com/berquerant/ga/internal"
	"golang.org/x/sync/errgroup"
)

// Evaluator is an interface for defining gene's score.
type Evaluator interface {
	// Eval calculates gene's score.
	Eval(ctx context.Context, gene Gene) (float64, error)
}

// IndividualsEvaluator is an interface for calculating individual's score.
type IndividualsEvaluator interface {
	// Eval calculates individuals' score and write them into the individuals.
	Eval(ctx context.Context, individuals []Individual) error
}

// NewIndividualsEvaluator returns a new IndividualsEvaluator.
// Evaluate individuals by evaluator concurrently.
func NewIndividualsEvaluator(evaluator Evaluator, concurrency int) (IndividualsEvaluator, error) {
	if evaluator == nil {
		return nil, NewGAError(nil, InvalidArgument, "need evaluator")
	}
	if concurrency < 1 {
		return nil, NewGAError(nil, InvalidArgument, "concurrency must be positive")
	}
	return &individualsEvaluator{
		evaluator:   evaluator,
		concurrency: concurrency,
	}, nil
}

type individualsEvaluator struct {
	evaluator   Evaluator
	concurrency int
}

func (s *individualsEvaluator) Eval(ctx context.Context, individuals []Individual) error {
	if internal.IsDone(ctx) {
		return NewGAError(ctx.Err(), EvalError, "timeout")
	}
	var (
		eg, eCtx     = errgroup.WithContext(ctx)
		iCtx, cancel = context.WithCancel(eCtx)
		requestC     = make(chan Individual, 2*s.concurrency)
	)
	defer cancel()
	for i := 0; i < s.concurrency; i++ {
		eg.Go(func() error {
			for r := range requestC {
				score, err := s.evaluator.Eval(iCtx, r.Gene())
				if err != nil {
					cancel()
					return NewGAError(err, EvalError, fmt.Sprintf("cannot evaluate %s", r.Gene()))
				}
				r.SetScore(score)
			}
			return nil
		})
	}
	for _, x := range individuals {
		requestC <- x
	}
	close(requestC)
	if err := eg.Wait(); err != nil {
		return NewGAError(err, EvalError, "cannot evaluate some genes")
	}
	return nil
}
