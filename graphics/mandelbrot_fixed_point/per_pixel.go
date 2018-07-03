package mandelbrot_fixed_point

import (
	"github.com/ncw/gmp"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"image/color"
	"math"
)

type MandelbrotPerPixel struct {
	Width, Height    int64
	MaxIter          int64
	centerR, centerI *gmp.Int
	zoom             *gmp.Int
	threshold2       *gmp.Int
	basePower2       uint
	Wrap             float64
	MaxVal           float64
	Cmap             colormap.ColorMap
	OutImage         *memory_mapped.PPMFile
	OutIter          *memory_mapped.Array2dFloat64
}

// assume that arbitrary precision integers are very expensive (because they usually are)
func (m *MandelbrotPerPixel) GetJobSize() int64 { return 128}

// TODO: builder pattern to reduce parameters?
func NewMandelbrotPerPixel(width, height, maxIter int64,
	centerR, centerI, zoom, threshold string, basePower2 uint,
	Wrap, MaxVal float64, cmap colormap.ColorMap,
	OutImage *memory_mapped.PPMFile,
	OutIter *memory_mapped.Array2dFloat64) *MandelbrotPerPixel {

	zoom_, success := new(gmp.Int).SetString(zoom, 10)
	if !success {
		panic("bad zoom string")
	}
	centerR_ := ParseFixnum(centerR, basePower2)
	centerI_ := ParseFixnum(centerI, basePower2)
	threshold2 := ParseFixnum(threshold, basePower2)
	threshold2.Mul(threshold2, threshold2)
	return &MandelbrotPerPixel{
		Width:      width,
		Height:     height,
		MaxIter:    maxIter,
		centerR:    centerR_,
		centerI:    centerI_,
		zoom:       zoom_,
		threshold2: threshold2,
		basePower2: basePower2,
		Wrap:       Wrap,
		MaxVal:     MaxVal,
		Cmap:       cmap,
		OutImage:   OutImage,
		OutIter:    OutIter,
	}
}

func (m *MandelbrotPerPixel) GetPixel(i, j int64) {
	cr, ci := MandelbrotCoordinate(i, j, m.Width, m.Height, m.centerR, m.centerI, m.zoom, m.basePower2)

	_, val, _ := MandelbrotKernel(cr, ci, m.threshold2, m.MaxIter, m.basePower2)

	if m.OutIter != nil {
		m.OutIter.Set(i, j, val)
	}

	if m.OutImage != nil {
		if val >= 0.0 {
			val = math.Log2(val+1) / math.Log2(m.MaxVal+1) * m.Wrap
			val = math.Sin(val*2*math.Pi)/2.0 + 0.5
			m.OutImage.Set64(i, j, m.Cmap.GetColor(val))
		} else {
			m.OutImage.Set64(i, j, color.RGBA{0, 0, 0, 255})
			return
		}
	}
}

func (m *MandelbrotPerPixel) Bounds() (w, h int64) { return m.Width, m.Height }
