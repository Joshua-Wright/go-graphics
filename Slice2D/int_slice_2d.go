// generated code, do not edit!
package Slice2D


// IntSlice2D is a 2-dimensional slice of int
type IntSlice2D struct {
	W, H int
	data []int
}

func NewIntSlice2D(x, y int) IntSlice2D {
	return IntSlice2D{
		W:    x,
		H:    y,
		data: make([]int, x*y),
	}
}

func (s *IntSlice2D) Get(x, y int) int {
	return s.data[s.W * y + x]
}

func (s *IntSlice2D) Set(x, y int, val int) {
	s.data[s.W * y + x] = val
}

func (s *IntSlice2D) At(x, y int) *int {
	return &s.data[s.W * y + x]
}
