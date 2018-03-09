// generated code, do not edit!
package graphics


// Vec3iSlice2D is a 2-dimensional slice of int
type Vec3iSlice2D struct {
	W, H int
	Data []Vec3i
}

func NewVec3iSlice2D(x, y int) Vec3iSlice2D {
	return Vec3iSlice2D{
		W:    x,
		H:    y,
		Data: make([]Vec3i, x*y),
	}
}

func (s *Vec3iSlice2D) Get(x, y int) Vec3i {
	return s.Data[s.W * y + x]
}

func (s *Vec3iSlice2D) Set(x, y int, val Vec3i) {
	s.Data[s.W * y + x] = val
}

func (s *Vec3iSlice2D) At(x, y int) *Vec3i {
	return &s.Data[s.W * y + x]
}
