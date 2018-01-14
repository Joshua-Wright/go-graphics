// generated code, do not edit!
package Slice2D


// UintSlice2D is a 2-dimensional slice of int
type UintSlice2D struct {
	W, H int
	data []uint
}

func NewUintSlice2D(x, y int) UintSlice2D {
	return UintSlice2D{
		W:    x,
		H:    y,
		data: make([]uint, x*y),
	}
}

func (s *UintSlice2D) Get(x, y int) uint {
	return s.data[s.W * y + x]
}

func (s *UintSlice2D) Set(x, y int, val uint) {
	s.data[s.W * y + x] = val
}

func (s *UintSlice2D) At(x, y int) *uint {
	return &s.data[s.W * y + x]
}
