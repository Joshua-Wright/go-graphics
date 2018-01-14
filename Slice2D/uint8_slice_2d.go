// generated code, do not edit!
package Slice2D


// Uint8Slice2D is a 2-dimensional slice of int
type Uint8Slice2D struct {
	W, H int
	data []uint8
}

func NewUint8Slice2D(x, y int) Uint8Slice2D {
	return Uint8Slice2D{
		W:    x,
		H:    y,
		data: make([]uint8, x*y),
	}
}

func (s *Uint8Slice2D) Get(x, y int) uint8 {
	return s.data[s.W * y + x]
}

func (s *Uint8Slice2D) Set(x, y int, val uint8) {
	s.data[s.W * y + x] = val
}

func (s *Uint8Slice2D) At(x, y int) *uint8 {
	return &s.data[s.W * y + x]
}
