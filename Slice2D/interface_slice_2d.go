// generated code, do not edit!
package Slice2D


// InterfaceSlice2D is a 2-dimensional slice of int
type InterfaceSlice2D struct {
	W, H int
	Data []interface{}
}

func NewInterfaceSlice2D(x, y int) InterfaceSlice2D {
	return InterfaceSlice2D{
		W:    x,
		H:    y,
		Data: make([]interface{}, x*y),
	}
}

func (s *InterfaceSlice2D) Get(x, y int) interface{} {
	return s.Data[s.W * y + x]
}

func (s *InterfaceSlice2D) Set(x, y int, val interface{}) {
	s.Data[s.W * y + x] = val
}

func (s *InterfaceSlice2D) At(x, y int) *interface{} {
	return &s.Data[s.W * y + x]
}
