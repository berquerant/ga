package ga

import (
	"math"
	"math/rand"
)

// Distribution is a probability distribution.
type Distribution interface {
	// Generate generates a random number.
	Generate() float64
}

// NewStandardUniformDistribution returns a new standard uniform distribution.
// Call rand.Seed if you want to get random values every executions.
func NewStandardUniformDistribution() Distribution {
	return &standardUniformDistribution{}
}

type standardUniformDistribution struct {
}

func (s *standardUniformDistribution) Generate() float64 { return rand.Float64() }

// Random is a random number generator.
type Random interface {
	// Float64 returns a random number in [min, max).
	Float64(min, max float64) float64
	// Int returns a random number in [min, max).
	Int(min, max int) int
}

// NewRandom returns a new Random.
// Call rand.Seed if you want to get random values every executions.
func NewRandom() Random { return &random{dist: &standardUniformDistribution{}} }

type random struct {
	dist *standardUniformDistribution
}

func (s *random) Float64(min, max float64) float64 {
	return min + (max-min)*s.dist.Generate()
}

func (s *random) Int(min, max int) int {
	x := int(math.Floor(s.dist.Generate() * float64(math.MaxInt32)))
	return min + x%(max-min)
}
