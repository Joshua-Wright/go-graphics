// generated code, do not edit!
package Slice2D
import "github.com/joshua-wright/go-graphics/graphics"

// Vec3Slice2D is a 2-dimensional slice of int
type Vec3Slice2D struct {
	W, H int
	data []graphics.Vec3
}

func NewVec3Slice2D(x, y int) Vec3Slice2D {
	return Vec3Slice2D{
		W:    x,
		H:    y,
		data: make([]graphics.Vec3, x*y),
	}
}

func (s *Vec3Slice2D) Get(x, y int) graphics.Vec3 {
	return s.data[s.W * y + x]
}

func (s *Vec3Slice2D) Set(x, y int, val graphics.Vec3) {
	s.data[s.W * y + x] = val
}

func (s *Vec3Slice2D) At(x, y int) *graphics.Vec3 {
	return &s.data[s.W * y + x]
}
