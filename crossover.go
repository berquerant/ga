package ga

// Crossover crosses genes.
type Crossover interface {
	// Cross generates a new gene from two genes.
	Cross(a, b Gene) (Gene, error)
}

func checkParents(a, b Gene) error {
	if a == nil || b == nil {
		return NewGAError(nil, CannotCrossover, "parents cannot be nil")
	}
	if a.Len() != b.Len() {
		return NewGAError(nil, CannotCrossover, "parents must have the same length")
	}
	if a.Len() == 0 {
		return NewGAError(nil, CannotCrossover, "parents must be not empty")
	}
	return nil
}

type uniformCrossover struct {
	rand Random
}

func NewUniformCrossover() Crossover {
	x, _ := NewUniformCrossoverWithRandom(NewRandom())
	return x
}

func NewUniformCrossoverWithRandom(rand Random) (Crossover, error) {
	if rand == nil {
		return nil, NewGAError(nil, InvalidArgument, "need rand")
	}
	return &uniformCrossover{
		rand: rand,
	}, nil
}

func (s *uniformCrossover) Cross(a, b Gene) (Gene, error) {
	if err := checkParents(a, b); err != nil {
		return nil, err
	}
	g := NewGene(a.Len())
	for i := 0; i < g.Len(); i++ {
		if s.rand.Int(0, 2) == 0 {
			g.Set(i, a.Get(i))
		} else {
			g.Set(i, b.Get(i))
		}
	}
	return g, nil
}

type twoPointCrossover struct {
	rand Random
}

func NewTwoPointCrossover() Crossover {
	x, _ := NewTwoPointCrossoverWithRandom(NewRandom())
	return x
}

func NewTwoPointCrossoverWithRandom(rand Random) (Crossover, error) {
	if rand == nil {
		return nil, NewGAError(nil, InvalidArgument, "need rand")
	}
	return &twoPointCrossover{
		rand: rand,
	}, nil
}

func (s *twoPointCrossover) Cross(a, b Gene) (Gene, error) {
	if err := checkParents(a, b); err != nil {
		return nil, err
	}
	if a.Len() < 3 {
		return nil, NewGAError(nil, CannotCrossover, "not enough gene length")
	}
	g := NewGene(a.Len())
	l := s.rand.Int(0, g.Len()-2)
	r := s.rand.Int(l+1, g.Len()-1)
	//  a b a
	// 0 l r len
	for i := 0; i < g.Len(); i++ {
		if i <= l || i > r {
			g.Set(i, a.Get(i))
		} else {
			g.Set(i, b.Get(i))
		}
	}
	return g, nil
}

type onePointCrossover struct {
	rand Random
}

func NewOnePointCrossover() Crossover {
	x, _ := NewOnePointCrossoverWithRandom(NewRandom())
	return x
}

func NewOnePointCrossoverWithRandom(rand Random) (Crossover, error) {
	if rand == nil {
		return nil, NewGAError(nil, InvalidArgument, "need rand")
	}
	return &onePointCrossover{
		rand: rand,
	}, nil
}

func (s *onePointCrossover) Cross(a, b Gene) (Gene, error) {
	if err := checkParents(a, b); err != nil {
		return nil, err
	}
	if a.Len() < 2 {
		return nil, NewGAError(nil, CannotCrossover, "not enough gene length")
	}
	g := NewGene(a.Len())
	p := s.rand.Int(0, g.Len()-1)
	for i := 0; i < g.Len(); i++ {
		if i <= p {
			g.Set(i, a.Get(i))
		} else {
			g.Set(i, b.Get(i))
		}
	}
	return g, nil
}
