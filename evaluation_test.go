package ga_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/berquerant/ga"
	"github.com/stretchr/testify/assert"
)

func TestIndividualsEvaluator(t *testing.T) {
	t.Run("no evaluator", func(t *testing.T) {
		_, err := ga.NewIndividualsEvaluator(nil, 2)
		assert.NotNil(t, err)
	})
	t.Run("zero concurrency", func(t *testing.T) {
		_, err := ga.NewIndividualsEvaluator(&mockCountEvaluator{}, 0)
		assert.NotNil(t, err)
	})
	t.Run("eval", func(t *testing.T) {
		s, err := ga.NewIndividualsEvaluator(&mockCountEvaluator{}, 2)
		assert.Nil(t, err)
		for _, tc := range []struct {
			name     string
			reverses []int
		}{
			{
				name: "no individuals",
			},
			{
				name:     "an individual",
				reverses: []int{1},
			},
			{
				name:     "2 individuals",
				reverses: []int{1, 0},
			},
			{
				name:     "many",
				reverses: []int{1, 3, 2, 4, 1, 1, 0, 3},
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				v := make([]ga.Individual, len(tc.reverses))
				for i := 0; i < len(v); i++ {
					g := ga.NewGene(5)
					for j := 0; j < tc.reverses[i]; j++ {
						g.Reverse(j)
					}
					v[i] = ga.NewIndividual(g)
				}
				assert.Nil(t, s.Eval(context.TODO(), v))
				for _, x := range v {
					assert.Equal(t, float64(countTrue(x.Gene())), x.GetScore())
				}
			})
		}
	})
	t.Run("cancel", func(t *testing.T) {
		const (
			heavyJob     = 400 * time.Millisecond
			middleJob    = 200 * time.Millisecond
			earlyTimeout = 50 * time.Millisecond
		)
		var (
			heavy = func() (float64, error) {
				time.Sleep(heavyJob)
				return 1.0, nil
			}
			lightError = func() (float64, error) {
				time.Sleep(middleJob)
				return 0, errors.New("light error")
			}
			heavyError = func() (float64, error) {
				time.Sleep(heavyJob)
				return 0, errors.New("heavy error")
			}
			individuals = func(size int) []ga.Individual {
				v := make([]ga.Individual, size)
				for i := 0; i < len(v); i++ {
					v[i] = ga.NewIndividual(ga.NewGene(5))
				}
				return v
			}
		)

		type timeoutTestcase struct {
			name    string
			f       []func() (float64, error)
			timeout time.Duration
		}
		for _, tc := range []*timeoutTestcase{
			{
				name: "an error",
				f: []func() (float64, error){
					lightError,
				},
			},
			{
				name: "timeout",
				f: []func() (float64, error){
					heavy,
				},
				timeout: earlyTimeout,
			},
			{
				name: "2 errors",
				f: []func() (float64, error){
					heavyError, lightError,
				},
			},
			{
				name: "timeout many jobs",
				f: []func() (float64, error){
					heavy, heavy, heavy, heavy,
				},
				timeout: earlyTimeout,
			},
			{
				name: "an error cancels many jobs",
				f: []func() (float64, error){
					heavy, heavyError, heavy, heavy,
				},
			},
		} {
			t.Run(tc.name, func(tc *timeoutTestcase) func(*testing.T) {
				return func(t *testing.T) {
					t.Parallel()
					s, err := ga.NewIndividualsEvaluator(&mockRotateFuncEvaluator{f: tc.f}, 2)
					assert.Nil(t, err)
					ctx, cancel := context.WithTimeout(context.TODO(), func() time.Duration {
						if tc.timeout > 0 {
							return tc.timeout
						}
						return time.Hour
					}())
					defer cancel()
					assert.NotNil(t, s.Eval(ctx, individuals(len(tc.f))))
				}
			}(tc))
		}
	})
}
