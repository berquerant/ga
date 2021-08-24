package ga_test

import (
	"context"
	"sync"

	"github.com/berquerant/ga"
)

type mockRotateFuncEvaluator struct {
	f   []func() (float64, error)
	i   int
	mux sync.Mutex
}

func (s *mockRotateFuncEvaluator) Eval(ctx context.Context, gene ga.Gene) (float64, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	s.mux.Lock()
	defer func() {
		s.i = (s.i + 1) % len(s.f)
		s.mux.Unlock()
	}()
	type result struct {
		f   float64
		err error
	}
	resultC := make(chan *result)
	go func() {
		f, err := s.f[s.i]()
		resultC <- &result{
			f:   f,
			err: err,
		}
		close(resultC)
	}()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case r := <-resultC:
		return r.f, r.err
	}
}

type mockCountEvaluator struct{}

func countTrue(gene ga.Gene) int {
	var sum int
	for i := 0; i < gene.Len(); i++ {
		if gene.Get(i) {
			sum++
		}
	}
	return sum
}

func (s *mockCountEvaluator) Eval(ctx context.Context, gene ga.Gene) (float64, error) {
	return float64(countTrue(gene)), nil
}

type mockRandom struct {
	f      []float64
	i      []int
	fi, ii int
}

func (s *mockRandom) Float64(_, _ float64) float64 {
	defer func() {
		s.fi = (s.fi + 1) % len(s.f)
	}()
	return s.f[s.fi]
}

func (s *mockRandom) Int(_, _ int) int {
	defer func() {
		s.ii = (s.ii + 1) % len(s.i)
	}()
	return s.i[s.ii]
}
