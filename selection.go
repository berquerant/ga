package ga

// Selector returns an index in the given list of individuals.
type Selector func() int
type SelectorGenerator func([]Individual) Selector

// IndividualsSelector is an interface for selecting individuals for next generation.
type IndividualsSelector interface {
	// Select generates individuals for next generation.
	// If not error, result individuals size and input size are the same.
	Select(individuals []Individual) ([]Individual, error)
}

type newGeneration int

const (
	crossover newGeneration = iota
	mutate
	regenerate
)

// NewIndividualsSelector returns a new IndividualsSelector.
//
// crossoverp chance of crossover.
// mutationp chance of mutation.
// 1 - (crossoverp + mutationp) chance of regenerate.
func NewIndividualsSelector(selectorGenrator SelectorGenerator, crossover Crossover, mutation Mutation,
	crossoverp, mutationp float64) (IndividualsSelector, error) {
	return NewIndividualsSelectorWithRandom(selectorGenrator, crossover, mutation, crossoverp, mutationp, NewRandom())
}

func NewIndividualsSelectorWithRandom(selectorGenrator SelectorGenerator, crossover Crossover, mutation Mutation,
	crossoverp, mutationp float64, rand Random) (IndividualsSelector, error) {
	if selectorGenrator == nil {
		return nil, NewGAError(nil, InvalidArgument, "need selector generator")
	}
	if crossover == nil {
		return nil, NewGAError(nil, InvalidArgument, "need crossover")
	}
	if mutation == nil {
		return nil, NewGAError(nil, InvalidArgument, "need mutation")
	}
	if rand == nil {
		return nil, NewGAError(nil, InvalidArgument, "need random")
	}
	if crossoverp < 0 {
		return nil, NewGAError(nil, InvalidArgument, "crossoverp cannot be negative")
	}
	if mutationp < 0 {
		return nil, NewGAError(nil, InvalidArgument, "mutationp cannot be negative")
	}
	if crossoverp+mutationp > 1 {
		return nil, NewGAError(nil, InvalidArgument, "crossoverp + mutationp must be less than 1")
	}
	return &individualsSelector{
		selectorGenerator: selectorGenrator,
		crossover:         crossover,
		mutation:          mutation,
		crossoverp:        crossoverp,
		mutationp:         mutationp,
		rand:              rand,
	}, nil
}

type individualsSelector struct {
	selectorGenerator     SelectorGenerator
	crossoverp, mutationp float64
	crossover             Crossover
	mutation              Mutation
	rand                  Random
}

func (s *individualsSelector) Select(individuals []Individual) ([]Individual, error) {
	var (
		size     = len(individuals)
		v        = make([]Individual, size)
		selector = s.selectorGenerator(individuals)
	)
	for i := 0; i < size; i++ {
		switch s.nextAction() {
		case crossover:
			{
				var (
					x = selector()
					y = selector()
				)
				for x == y {
					y = selector()
				}
				if x < 0 || y < 0 || x >= size || y >= size {
					return nil, NewGAError(nil, CannotSelectIndividual, "selector returned out of range index")
				}
				g, err := s.crossover.Cross(individuals[x].Gene(), individuals[y].Gene())
				if err != nil {
					return nil, NewGAError(err, CannotSelectIndividual, "crossover error")
				}
				v[i] = NewIndividual(g)
			}
		case mutate:
			{
				x := s.rand.Int(0, size)
				g := s.mutation.Mutate(individuals[x].Gene())
				v[i] = NewIndividual(g)
			}
		default:
			{
				x := s.rand.Int(0, size)
				v[i] = NewIndividual(individuals[x].Gene())
			}
		}
	}
	return v, nil
}

func (s *individualsSelector) nextAction() newGeneration {
	p := s.rand.Float64(0, 1)
	p -= s.crossoverp
	if p < 0 {
		return crossover
	}
	p -= s.mutationp
	if p < 0 {
		return mutate
	}
	return regenerate
}
