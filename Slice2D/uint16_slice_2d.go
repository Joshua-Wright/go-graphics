// generated code, do not edit!
package Slice2D


// Uint16Slice2D is a 2-dimensional slice of int
type Uint16Slice2D struct {
	W, H int
	data []uint16
}

func NewUint16Slice2D(x, y int) Uint16Slice2D {
	return Uint16Slice2D{
		W:    x,
		H:    y,
		data: make([]uint16, x*y),
	}
}

func (s *Uint16Slice2D) Get(x, y int) uint16 {
	return s.data[s.W * y + x]
}

func (s *Uint16Slice2D) Set(x, y int, val uint16) {
	s.data[s.W * y + x] = val
}

func (s *Uint16Slice2D) At(x, y int) *uint16 {
	return &s.data[s.W * y + x]
}
