package mandelbrot_fixed_point

import "github.com/ncw/gmp"

type MandelbrotSinglePixelConfig struct {
	I          int64  `json:"i"`
	J          int64  `json:"j"`
	Width      int64  `json:"width"`
	Height     int64  `json:"height"`
	MaxIter    int64  `json:"maxIter"`
	CenterR    string `json:"centerR"`
	CenterI    string `json:"centerI"`
	Zoom       string `json:"zoom"`
	Threshold2 string `json:"threshold2"`
	BasePower2 uint   `json:"basePower2"`
}

type MandelbrotSinglePixelResponse struct {
	I         int64   `json:"i"`
	J         int64   `json:"j"`
	Iteration int64   `json:"iteration"`
	Val       float64 `json:"val"`
	Mag2      string  `json:"mag2"`
}

type MandelbrotPixelRangeConfig struct {
	// TODO put in separate object so goroutines can share
	Imin int64
	Imax int64
	Jmin int64
	Jmax int64

	Width           int64
	Height          int64
	MaxIter         int64
	CenterR         *gmp.Int
	CenterI         *gmp.Int
	Zoom            *gmp.Int
	Threshold2      *gmp.Int
	BasePower2      uint
	ReturnIteration bool
	ReturnVal       bool
	ReturnMag2      bool
}

// if one of these wasn't requested, return empty array instead
type MandelbrotPixelRangeResponse struct {
	Iteration []int64
	Val       []float64
	Mag2      []*gmp.Int
}
