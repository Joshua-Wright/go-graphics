// generated code, do not edit!
package graphics


// Vec3Slice2D is a 2-dimensional slice of int
type Vec3Slice2D struct {
	W, H int
	Data []Vec3
}

func NewVec3Slice2D(x, y int) Vec3Slice2D {
	return Vec3Slice2D{
		W:    x,
		H:    y,
		Data: make([]Vec3, x*y),
	}
}

func (s *Vec3Slice2D) Get(x, y int) Vec3 {
	return s.Data[s.W * y + x]
}

func (s *Vec3Slice2D) Set(x, y int, val Vec3) {
	s.Data[s.W * y + x] = val
}

func (s *Vec3Slice2D) At(x, y int) *Vec3 {
	return &s.Data[s.W * y + x]
}
