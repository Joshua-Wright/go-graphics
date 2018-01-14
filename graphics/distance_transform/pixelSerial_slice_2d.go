// generated code, do not edit!
package distance_transform


// PixelSerialSlice2D is a 2-dimensional slice of int
type PixelSerialSlice2D struct {
	W, H int
	Data []pixelSerial
}

func NewPixelSerialSlice2D(x, y int) PixelSerialSlice2D {
	return PixelSerialSlice2D{
		W:    x,
		H:    y,
		Data: make([]pixelSerial, x*y),
	}
}

func (s *PixelSerialSlice2D) Get(x, y int) pixelSerial {
	return s.Data[s.W * y + x]
}

func (s *PixelSerialSlice2D) Set(x, y int, val pixelSerial) {
	s.Data[s.W * y + x] = val
}

func (s *PixelSerialSlice2D) At(x, y int) *pixelSerial {
	return &s.Data[s.W * y + x]
}
