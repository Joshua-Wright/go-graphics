// generated code, do not edit!
package Slice2D


// Float64Slice2D is a 2-dimensional slice of int
type Float64Slice2D struct {
	W, H int
	Data []float64
}

func NewFloat64Slice2D(x, y int) Float64Slice2D {
	return Float64Slice2D{
		W:    x,
		H:    y,
		Data: make([]float64, x*y),
	}
}

func (s *Float64Slice2D) Get(x, y int) float64 {
	return s.Data[s.W * y + x]
}

func (s *Float64Slice2D) Set(x, y int, val float64) {
	s.Data[s.W * y + x] = val
}

func (s *Float64Slice2D) At(x, y int) *float64 {
	return &s.Data[s.W * y + x]
}
