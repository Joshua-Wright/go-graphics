// generated code, do not edit!
package Slice2D


// StringSlice2D is a 2-dimensional slice of int
type StringSlice2D struct {
	W, H int
	data []string
}

func NewStringSlice2D(x, y int) StringSlice2D {
	return StringSlice2D{
		W:    x,
		H:    y,
		data: make([]string, x*y),
	}
}

func (s *StringSlice2D) Get(x, y int) string {
	return s.data[s.W * y + x]
}

func (s *StringSlice2D) Set(x, y int, val string) {
	s.data[s.W * y + x] = val
}

func (s *StringSlice2D) At(x, y int) *string {
	return &s.data[s.W * y + x]
}
