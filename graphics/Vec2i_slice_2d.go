// generated code, do not edit!
package graphics


// Vec2iSlice2D is a 2-dimensional slice of int
type Vec2iSlice2D struct {
	W, H int
	Data []Vec2i
}

func NewVec2iSlice2D(x, y int) Vec2iSlice2D {
	return Vec2iSlice2D{
		W:    x,
		H:    y,
		Data: make([]Vec2i, x*y),
	}
}

func (s *Vec2iSlice2D) Get(x, y int) Vec2i {
	return s.Data[s.W * y + x]
}

func (s *Vec2iSlice2D) Set(x, y int, val Vec2i) {
	s.Data[s.W * y + x] = val
}

func (s *Vec2iSlice2D) At(x, y int) *Vec2i {
	return &s.Data[s.W * y + x]
}
