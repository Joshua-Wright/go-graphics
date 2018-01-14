// generated code, do not edit!
package Slice2D
import "github.com/joshua-wright/go-graphics/graphics"

// Vec2iSlice2D is a 2-dimensional slice of int
type Vec2iSlice2D struct {
	W, H int
	data []graphics.Vec2i
}

func NewVec2iSlice2D(x, y int) Vec2iSlice2D {
	return Vec2iSlice2D{
		W:    x,
		H:    y,
		data: make([]graphics.Vec2i, x*y),
	}
}

func (s *Vec2iSlice2D) Get(x, y int) graphics.Vec2i {
	return s.data[s.W * y + x]
}

func (s *Vec2iSlice2D) Set(x, y int, val graphics.Vec2i) {
	s.data[s.W * y + x] = val
}

func (s *Vec2iSlice2D) At(x, y int) *graphics.Vec2i {
	return &s.data[s.W * y + x]
}
