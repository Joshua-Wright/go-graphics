// generated code, do not edit!
package Slice2D


// Int8Slice2D is a 2-dimensional slice of int
type Int8Slice2D struct {
	W, H int
	Data []int8
}

func NewInt8Slice2D(x, y int) Int8Slice2D {
	return Int8Slice2D{
		W:    x,
		H:    y,
		Data: make([]int8, x*y),
	}
}

func (s *Int8Slice2D) Get(x, y int) int8 {
	return s.Data[s.W * y + x]
}

func (s *Int8Slice2D) Set(x, y int, val int8) {
	s.Data[s.W * y + x] = val
}

func (s *Int8Slice2D) At(x, y int) *int8 {
	return &s.Data[s.W * y + x]
}
