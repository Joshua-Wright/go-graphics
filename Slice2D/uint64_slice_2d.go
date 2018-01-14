// generated code, do not edit!
package Slice2D


// Uint64Slice2D is a 2-dimensional slice of int
type Uint64Slice2D struct {
	W, H int
	Data []uint64
}

func NewUint64Slice2D(x, y int) Uint64Slice2D {
	return Uint64Slice2D{
		W:    x,
		H:    y,
		Data: make([]uint64, x*y),
	}
}

func (s *Uint64Slice2D) Get(x, y int) uint64 {
	return s.Data[s.W * y + x]
}

func (s *Uint64Slice2D) Set(x, y int, val uint64) {
	s.Data[s.W * y + x] = val
}

func (s *Uint64Slice2D) At(x, y int) *uint64 {
	return &s.Data[s.W * y + x]
}
