package mandelbrot_fixed_point

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
