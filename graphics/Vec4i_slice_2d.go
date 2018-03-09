// generated code, do not edit!
package graphics


// Vec4iSlice2D is a 2-dimensional slice of int
type Vec4iSlice2D struct {
	W, H int
	Data []Vec4i
}

func NewVec4iSlice2D(x, y int) Vec4iSlice2D {
	return Vec4iSlice2D{
		W:    x,
		H:    y,
		Data: make([]Vec4i, x*y),
	}
}

func (s *Vec4iSlice2D) Get(x, y int) Vec4i {
	return s.Data[s.W * y + x]
}

func (s *Vec4iSlice2D) Set(x, y int, val Vec4i) {
	s.Data[s.W * y + x] = val
}

func (s *Vec4iSlice2D) At(x, y int) *Vec4i {
	return &s.Data[s.W * y + x]
}
