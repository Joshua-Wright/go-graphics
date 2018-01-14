// generated code, do not edit!
package Slice2D
import "github.com/joshua-wright/go-graphics/graphics"

// Vec3iSlice2D is a 2-dimensional slice of int
type Vec3iSlice2D struct {
	W, H int
	data []graphics.Vec3i
}

func NewVec3iSlice2D(x, y int) Vec3iSlice2D {
	return Vec3iSlice2D{
		W:    x,
		H:    y,
		data: make([]graphics.Vec3i, x*y),
	}
}

func (s *Vec3iSlice2D) Get(x, y int) graphics.Vec3i {
	return s.data[s.W * y + x]
}

func (s *Vec3iSlice2D) Set(x, y int, val graphics.Vec3i) {
	s.data[s.W * y + x] = val
}

func (s *Vec3iSlice2D) At(x, y int) *graphics.Vec3i {
	return &s.data[s.W * y + x]
}
