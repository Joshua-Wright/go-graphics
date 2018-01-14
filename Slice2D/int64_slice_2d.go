// generated code, do not edit!
package Slice2D


// Int64Slice2D is a 2-dimensional slice of int
type Int64Slice2D struct {
	W, H int
	Data []int64
}

func NewInt64Slice2D(x, y int) Int64Slice2D {
	return Int64Slice2D{
		W:    x,
		H:    y,
		Data: make([]int64, x*y),
	}
}

func (s *Int64Slice2D) Get(x, y int) int64 {
	return s.Data[s.W * y + x]
}

func (s *Int64Slice2D) Set(x, y int, val int64) {
	s.Data[s.W * y + x] = val
}

func (s *Int64Slice2D) At(x, y int) *int64 {
	return &s.Data[s.W * y + x]
}
