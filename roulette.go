package ga

// Roulette selects an element randomly.
type Roulette interface {
	// Spin returns the index of the selected element.
	Spin() int
}

type roulette struct {
	weights []float64
	sum     float64
	rand    Random
}

// NewRoulette returns a new roulette.
// The weights must contain some elements, and they must be all non-negative.
func NewRoulette(weights []float64) (Roulette, error) {
	return NewRouletteWithRandom(weights, NewRandom())
}

// NewRouletteWithRandom returns a new roulette that uses rand to select an element.
func NewRouletteWithRandom(weights []float64, rand Random) (Roulette, error) {
	if len(weights) == 0 {
		return nil, NewGAError(nil, InvalidArgument, "need some weights")
	}
	if rand == nil {
		return nil, NewGAError(nil, InvalidArgument, "need rand")
	}
	var sum float64
	for _, x := range weights {
		if x < 0 {
			return nil, NewGAError(nil, InvalidArgument, "all weights must be non-negative")
		}
		sum += x
	}
	return &roulette{
		weights: weights,
		rand:    rand,
		sum:     sum,
	}, nil
}

func (s *roulette) Spin() int {
	p := s.rand.Float64(0, s.sum)
	for i, x := range s.weights {
		p -= x
		if p < 0 {
			return i
		}
	}
	return len(s.weights) - 1
}
