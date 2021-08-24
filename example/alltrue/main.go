package main

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/berquerant/ga"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

type Eval struct {
}

func countTrue(gene ga.Gene) int {
	var c int
	for i := 0; i < gene.Len(); i++ {
		if gene.Get(i) {
			c++
		}
	}
	return c
}

func (s *Eval) Eval(ctx context.Context, gene ga.Gene) (float64, error) {
	sum := float64(countTrue(gene))
	if sum > 20 {
		sum += (sum - 20) * (sum - 20) / 2 // encourage to set more bits
	}
	return sum, nil
}

func initialIndividuals(iSize, gSize int) []ga.Individual {
	v := make([]ga.Individual, iSize)
	for i := 0; i < len(v); i++ {
		v[i] = ga.NewIndividual(ga.NewRandomGene(gSize))
	}
	return v
}

func printGeneration(n int, g ga.Generation) {
	individuals := g.Individuals()
	sort.Slice(individuals, func(i, j int) bool { return individuals[i].GetScore() > individuals[j].GetScore() })
	best := individuals[0]
	worst := individuals[len(individuals)-1]
	fmt.Printf("generation: %3d best %2d worst %2d best gene %s\n", n, countTrue(best.Gene()), countTrue(worst.Gene()), best.Gene())
}

func main() {
	rand.Seed(time.Now().UnixNano()) // because ga.Random created by ga.NewRandom depends on math/rand
	const (
		individualsSize = 30
		geneSize        = 32
		generationLimit = 100
		crossoverp      = 0.9
		mutationSize    = 1
		mutationp       = 0.02
		evalConcurrency = 4
	)
	// generation settings
	var b ga.GenerationBuilder
	b.SelectorGenerator(func(v []ga.Individual) ga.Selector {
		weights := make([]float64, len(v))
		for i, x := range v {
			weights[i] = x.GetScore()
		}
		r, err := ga.NewRoulette(weights)
		panicOnError(err)
		return r.Spin
	})
	b.Evaluator(&Eval{})
	b.Crossover(ga.NewUniformCrossover())
	m, err := ga.NewMutation(mutationSize)
	panicOnError(err)
	b.Mutation(m)
	b.Individuals(initialIndividuals(individualsSize, geneSize))
	b.EvalConcurrency(evalConcurrency)
	b.Mutationp(mutationp)
	b.Crossoverp(crossoverp)
	// start optimization
	generation, err := b.Build()
	panicOnError(err)
	ctx := context.Background()
	for i := 1; i <= generationLimit; i++ {
		next, err := generation.Next(ctx)
		panicOnError(err)
		if i == 1 || i%10 == 0 {
			printGeneration(i, generation)
		}
		generation = next
	}
}
