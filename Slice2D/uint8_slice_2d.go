// generated code, do not edit!
package Slice2D


// Uint8Slice2D is a 2-dimensional slice of int
type Uint8Slice2D struct {
	W, H int
	Data []uint8
}

func NewUint8Slice2D(x, y int) Uint8Slice2D {
	return Uint8Slice2D{
		W:    x,
		H:    y,
		Data: make([]uint8, x*y),
	}
}

func (s *Uint8Slice2D) Get(x, y int) uint8 {
	return s.Data[s.W * y + x]
}

func (s *Uint8Slice2D) Set(x, y int, val uint8) {
	s.Data[s.W * y + x] = val
}

func (s *Uint8Slice2D) At(x, y int) *uint8 {
	return &s.Data[s.W * y + x]
}
