// generated code, do not edit!
package Slice2D
import "github.com/joshua-wright/go-graphics/graphics"

// Vec4iSlice2D is a 2-dimensional slice of int
type Vec4iSlice2D struct {
	W, H int
	data []graphics.Vec4i
}

func NewVec4iSlice2D(x, y int) Vec4iSlice2D {
	return Vec4iSlice2D{
		W:    x,
		H:    y,
		data: make([]graphics.Vec4i, x*y),
	}
}

func (s *Vec4iSlice2D) Get(x, y int) graphics.Vec4i {
	return s.data[s.W * y + x]
}

func (s *Vec4iSlice2D) Set(x, y int, val graphics.Vec4i) {
	s.data[s.W * y + x] = val
}

func (s *Vec4iSlice2D) At(x, y int) *graphics.Vec4i {
	return &s.data[s.W * y + x]
}
