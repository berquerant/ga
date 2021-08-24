package ga

import "fmt"

// Individual is a wrapper of a gene to optimize.
type Individual interface {
	// Gene returns the gene.
	Gene() Gene
	// GetScore returns the score.
	GetScore() float64
	// SetScore sets score to this.
	SetScore(score float64)
	// Clone returns a duplicated individual.
	Clone() Individual
}

// NewIndividual returns a new individual with score zero.
func NewIndividual(gene Gene) Individual {
	return &individual{
		gene: gene,
	}
}

type individual struct {
	gene  Gene
	score float64
}

func (s *individual) Clone() Individual {
	return &individual{
		score: s.score,
		gene:  s.gene.Clone(),
	}
}
func (s *individual) Gene() Gene             { return s.gene }
func (s *individual) GetScore() float64      { return s.score }
func (s *individual) SetScore(score float64) { s.score = score }
func (s *individual) String() string         { return fmt.Sprintf("score: %f gene: %s", s.score, s.gene) }
