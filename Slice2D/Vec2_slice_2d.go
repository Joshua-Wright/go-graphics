// generated code, do not edit!
package Slice2D
import "github.com/joshua-wright/go-graphics/graphics"

// Vec2Slice2D is a 2-dimensional slice of int
type Vec2Slice2D struct {
	W, H int
	data []graphics.Vec2
}

func NewVec2Slice2D(x, y int) Vec2Slice2D {
	return Vec2Slice2D{
		W:    x,
		H:    y,
		data: make([]graphics.Vec2, x*y),
	}
}

func (s *Vec2Slice2D) Get(x, y int) graphics.Vec2 {
	return s.data[s.W * y + x]
}

func (s *Vec2Slice2D) Set(x, y int, val graphics.Vec2) {
	s.data[s.W * y + x] = val
}

func (s *Vec2Slice2D) At(x, y int) *graphics.Vec2 {
	return &s.data[s.W * y + x]
}
