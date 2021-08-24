package internal

type IntSet interface {
	Set(v int)
	Get(v int) bool
	Size() int
	Slice() []int
}

func newInterfaceSet() *interfaceSet {
	return &interfaceSet{
		d: map[interface{}]bool{},
	}
}

type interfaceSet struct {
	d map[interface{}]bool
}

func (s *interfaceSet) get(v interface{}) bool { return s.d[v] }
func (s *interfaceSet) set(v interface{})      { s.d[v] = true }
func (s *interfaceSet) size() int              { return len(s.d) }
func (s *interfaceSet) slice() []interface{} {
	var i int
	v := make([]interface{}, len(s.d))
	for x := range s.d {
		v[i] = x
		i++
	}
	return v
}

func NewIntSet() IntSet {
	return &intSet{
		d: newInterfaceSet(),
	}
}

type intSet struct {
	d *interfaceSet
}

func (s *intSet) Get(v int) bool { return s.d.get(v) }
func (s *intSet) Set(v int)      { s.d.set(v) }
func (s *intSet) Size() int      { return s.d.size() }
func (s *intSet) Slice() []int {
	v := make([]int, s.d.size())
	for i, x := range s.d.slice() {
		v[i] = x.(int)
	}
	return v
}
