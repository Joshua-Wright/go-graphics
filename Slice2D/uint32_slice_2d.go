// generated code, do not edit!
package Slice2D


// Uint32Slice2D is a 2-dimensional slice of int
type Uint32Slice2D struct {
	W, H int
	Data []uint32
}

func NewUint32Slice2D(x, y int) Uint32Slice2D {
	return Uint32Slice2D{
		W:    x,
		H:    y,
		Data: make([]uint32, x*y),
	}
}

func (s *Uint32Slice2D) Get(x, y int) uint32 {
	return s.Data[s.W * y + x]
}

func (s *Uint32Slice2D) Set(x, y int, val uint32) {
	s.Data[s.W * y + x] = val
}

func (s *Uint32Slice2D) At(x, y int) *uint32 {
	return &s.Data[s.W * y + x]
}
