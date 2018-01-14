// generated code, do not edit!
package Slice2D
import "github.com/joshua-wright/go-graphics/graphics"

// Vec4Slice2D is a 2-dimensional slice of int
type Vec4Slice2D struct {
	W, H int
	Data []graphics.Vec4
}

func NewVec4Slice2D(x, y int) Vec4Slice2D {
	return Vec4Slice2D{
		W:    x,
		H:    y,
		Data: make([]graphics.Vec4, x*y),
	}
}

func (s *Vec4Slice2D) Get(x, y int) graphics.Vec4 {
	return s.Data[s.W * y + x]
}

func (s *Vec4Slice2D) Set(x, y int, val graphics.Vec4) {
	s.Data[s.W * y + x] = val
}

func (s *Vec4Slice2D) At(x, y int) *graphics.Vec4 {
	return &s.Data[s.W * y + x]
}
