// generated code, do not edit!
package Slice2D


// Int16Slice2D is a 2-dimensional slice of int
type Int16Slice2D struct {
	W, H int
	data []int16
}

func NewInt16Slice2D(x, y int) Int16Slice2D {
	return Int16Slice2D{
		W:    x,
		H:    y,
		data: make([]int16, x*y),
	}
}

func (s *Int16Slice2D) Get(x, y int) int16 {
	return s.data[s.W * y + x]
}

func (s *Int16Slice2D) Set(x, y int, val int16) {
	s.data[s.W * y + x] = val
}

func (s *Int16Slice2D) At(x, y int) *int16 {
	return &s.data[s.W * y + x]
}
