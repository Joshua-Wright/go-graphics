// generated code, do not edit!
package Slice2D


// Float32Slice2D is a 2-dimensional slice of int
type Float32Slice2D struct {
	W, H int
	data []float32
}

func NewFloat32Slice2D(x, y int) Float32Slice2D {
	return Float32Slice2D{
		W:    x,
		H:    y,
		data: make([]float32, x*y),
	}
}

func (s *Float32Slice2D) Get(x, y int) float32 {
	return s.data[s.W * y + x]
}

func (s *Float32Slice2D) Set(x, y int, val float32) {
	s.data[s.W * y + x] = val
}

func (s *Float32Slice2D) At(x, y int) *float32 {
	return &s.data[s.W * y + x]
}
