// generated code, do not edit!
package Slice2D


// Int8Slice2D is a 2-dimensional slice of int
type Int8Slice2D struct {
	W, H int
	data []int8
}

func NewInt8Slice2D(x, y int) Int8Slice2D {
	return Int8Slice2D{
		W:    x,
		H:    y,
		data: make([]int8, x*y),
	}
}

func (s *Int8Slice2D) Get(x, y int) int8 {
	return s.data[s.W * y + x]
}

func (s *Int8Slice2D) Set(x, y int, val int8) {
	s.data[s.W * y + x] = val
}

func (s *Int8Slice2D) At(x, y int) *int8 {
	return &s.data[s.W * y + x]
}
