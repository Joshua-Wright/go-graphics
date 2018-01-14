// generated code, do not edit!
package Slice2D


// Int32Slice2D is a 2-dimensional slice of int
type Int32Slice2D struct {
	W, H int
	data []int32
}

func NewInt32Slice2D(x, y int) Int32Slice2D {
	return Int32Slice2D{
		W:    x,
		H:    y,
		data: make([]int32, x*y),
	}
}

func (s *Int32Slice2D) Get(x, y int) int32 {
	return s.data[s.W * y + x]
}

func (s *Int32Slice2D) Set(x, y int, val int32) {
	s.data[s.W * y + x] = val
}

func (s *Int32Slice2D) At(x, y int) *int32 {
	return &s.data[s.W * y + x]
}
