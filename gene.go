package ga

// Gene is a binary gene, the target of optimization.
type Gene interface {
	// Get returns a section of the gene.
	// Returns false if out of range.
	Get(i int) bool
	// Set sets a section of the gene.
	// No operation if out of range.
	Set(i int, v bool)
	// Reverse reverses a section of the gene.
	// Returns the section.
	// No operation if out of range.
	Reverse(i int)
	// Len returns the length of the gene.
	Len() int
	// Clone returns a duplicated gene.
	Clone() Gene
}

// NewGene returns a new gene with false sections.
func NewGene(size int) Gene {
	return &gene{
		v: make([]bool, size),
	}
}

// NewRandomGene returns a new gene with random sections.
func NewRandomGene(size int) Gene {
	r := NewRandom()
	g := NewGene(size)
	for i := 0; i < g.Len(); i++ {
		if r.Int(0, 2) == 0 {
			g.Reverse(i)
		}
	}
	return g
}

type gene struct {
	v []bool
}

func (s *gene) in(i int) bool { return i >= 0 && i < len(s.v) }
func (s *gene) Len() int      { return len(s.v) }
func (s *gene) Get(i int) bool {
	if s.in(i) {
		return s.v[i]
	}
	return false
}
func (s *gene) Set(i int, v bool) {
	if s.in(i) {
		s.v[i] = v
	}
}
func (s *gene) Reverse(i int) {
	if s.in(i) {
		s.v[i] = !s.v[i]
	}
}
func (s *gene) Clone() Gene {
	v := make([]bool, len(s.v))
	for i, x := range s.v {
		v[i] = x
	}
	return &gene{
		v: v,
	}
}
func (s *gene) String() string {
	v := make([]rune, len(s.v))
	for i := 0; i < len(s.v); i++ {
		if s.v[i] {
			v[i] = '1'
		} else {
			v[i] = '0'
		}
	}
	return string(v)
}
