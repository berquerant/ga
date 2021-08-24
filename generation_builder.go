package ga

// GenerationBuilder is a builder of Generation.
type GenerationBuilder struct {
	selectorGenerator SelectorGenerator
	evaluator         Evaluator
	crossover         Crossover
	mutation          Mutation
	individuals       []Individual
	evalConcurrency   int
	mutationp         float64
	crossoverp        float64
}

// Build returns a new Generation.
func (s *GenerationBuilder) Build() (Generation, error) {
	if err := s.validate(); err != nil {
		return nil, NewGAError(err, CannotBuildGeneration, "failed to build generation")
	}
	evaluator, err := NewIndividualsEvaluator(s.evaluator, s.evalConcurrency)
	if err != nil {
		return nil, NewGAError(err, CannotBuildGeneration, "failed to build generation")
	}
	selector, err := NewIndividualsSelector(s.selectorGenerator, s.crossover, s.mutation, s.crossoverp, s.mutationp)
	if err != nil {
		return nil, NewGAError(err, CannotBuildGeneration, "failed to build generation")
	}
	return &generation{
		selector:    selector,
		evaluator:   evaluator,
		individuals: s.individuals,
	}, nil
}

func (s *GenerationBuilder) validate() error {
	if len(s.individuals) < 2 {
		return NewGAError(nil, InvalidArgument, "size of individuals cannot be less than 2")
	}
	return nil
}

// SelectorGenerator sets a function to generate a function to select an individual.
func (s *GenerationBuilder) SelectorGenerator(v SelectorGenerator) {
	s.selectorGenerator = v
}

// Evaluator sets an evaluator of a gene.
func (s *GenerationBuilder) Evaluator(v Evaluator) {
	s.evaluator = v
}

// Crossover sets a strategy for crossover.
func (s *GenerationBuilder) Crossover(v Crossover) {
	s.crossover = v
}

// Mutation sets a strategy for mutation.
func (s *GenerationBuilder) Mutation(v Mutation) {
	s.mutation = v
}

// Individuals sets the initial individuals.
func (s *GenerationBuilder) Individuals(v []Individual) {
	s.individuals = v
}

// EvalConcurrency sets a concurrency to evaluate individuals.
func (s *GenerationBuilder) EvalConcurrency(v int) {
	s.evalConcurrency = v
}

// Mutationp sets a probability for mutation.
func (s *GenerationBuilder) Mutationp(v float64) {
	s.mutationp = v
}

// Crossoverp sets a probability for crossover.
func (s *GenerationBuilder) Crossoverp(v float64) {
	s.crossoverp = v
}
