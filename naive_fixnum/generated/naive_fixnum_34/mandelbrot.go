package naive_fixnum_34

import (
	"math"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"math/big"
	"image/color"
)

var two = new(Fixnum).FromInt32(2)

type MandelbrotPerPixel struct {
	Width, Height   int64
	MaxIter         int64
	centerR         Fixnum
	centerI         Fixnum
	dr              Fixnum
	threshold2      Fixnum
	Wrap            float64
	MaxVal          float64
	Cmap            colormap.ColorMap
	OutImage        *memory_mapped.PPMFile
	OutIter         *memory_mapped.Array2dFloat64
	primaryDelegate *MandelbrotPerPixelDelegate
}

func (m *MandelbrotPerPixel) GetPixel(i, j int64) {
	if m.primaryDelegate == nil {
		println("don't use this MandelbrotPerPixel directly!")
		m.primaryDelegate = new(MandelbrotPerPixelDelegate)
	}
	m.primaryDelegate.GetPixel(i, j)
}

func (m *MandelbrotPerPixel) GetCachedWorker() per_pixel_image.PixelFunction {
	return &MandelbrotPerPixelDelegate{
		parent: m,
	}
}

// assume that arbitrary precision integers are very expensive (because they usually are)
func (m *MandelbrotPerPixel) GetJobSize() int64 { return 128 }

func NewMandelbrotPerPixel(width, height, maxIter int64,
	centerR, centerI, zoom, threshold string,
	Wrap, MaxVal float64, cmap colormap.ColorMap,
	OutImage *memory_mapped.PPMFile,
	OutIter *memory_mapped.Array2dFloat64) *MandelbrotPerPixel {

	m := &MandelbrotPerPixel{
		Width:      width,
		Height:     height,
		MaxIter:    maxIter,
		Wrap:       Wrap,
		MaxVal:     MaxVal,
		Cmap:       cmap,
		OutImage:   OutImage,
		OutIter:    OutIter,
	}

	prec := uint(FpWords * 32 * 2)
	zoomBf, _, err := big.ParseFloat(zoom, 10, prec, big.ToNearestEven)
	if err != nil {
		panic(err)
	}
	fourBf := big.NewFloat(4.0).SetPrec(prec)
	widthBf := big.NewFloat(float64(width)).SetPrec(prec)
	fourBf.Quo(fourBf, zoomBf.Mul(zoomBf, widthBf))
	m.dr.FromBigFloat(fourBf)

	_, err = m.centerR.FromString(centerR)
	if err != nil {
		panic(err)
	}
	_, err = m.centerI.FromString(centerI)
	if err != nil {
		panic(err)
	}
	_, err = m.threshold2.FromString(threshold)
	if err != nil {
		panic(err)
	}
	m.threshold2.Mul(&m.threshold2, &m.threshold2)

	return m
}

type MandelbrotPerPixelDelegate struct {
	parent *MandelbrotPerPixel

	/// used to find coordinate
	r, i   Fixnum
	cr, ci Fixnum

	/// used by mandelbrot kernel
	zr         Fixnum
	zi         Fixnum
	zr2        Fixnum
	zi2        Fixnum
	magnitude2 Fixnum
	zri        Fixnum
	val        float64
}

func (m *MandelbrotPerPixelDelegate) SetMandelbrotCoordinate(x, y int64) {
	m.r.FromInt32(int32(x - m.parent.Width/2))
	m.i.FromInt32(int32(m.parent.Height/2 - y))

	m.cr.Mul(&m.r, &m.parent.dr)
	m.cr.Add(&m.cr, &m.parent.centerR)

	m.ci.Mul(&m.i, &m.parent.dr)
	m.ci.Add(&m.ci, &m.parent.centerI)

	//fmt.Println(m.cr.Float64(), m.ci.Float64())
}

func (m *MandelbrotPerPixelDelegate) MandelbrotKernel() {
	m.zr.SetZero()
	m.zi.SetZero()
	m.zr2.SetZero()
	m.zi2.SetZero()
	m.magnitude2.SetZero()
	m.zri.SetZero()
	for i := int64(0); i < m.parent.MaxIter; i++ {
		m.magnitude2.Add(&m.zr2, &m.zi2)
		if m.magnitude2.cmpWords(&m.parent.threshold2) > 0 {
			z2 := m.magnitude2.Float64()
			v := float64(i-1) - math.Log2(math.Log2(z2)) + 1
			if math.IsNaN(v) {
				v = float64(i - 1)
			}
			m.val = v
			return
		}
		// otherwise re-use those values

		// calculate 2ab/B + m.ci
		// (calculate out of place to not disturb &m.zi for the next calculation)
		m.zri.Mul(&m.zr, &m.zi)
		m.zri.Mul(&m.zri, two).Add(&m.zri, &m.ci)

		// calculate (a^2 - b^2)/B + m.cr (in place this time)
		m.zr.Sub(&m.zr2, &m.zi2).Add(&m.zr, &m.cr)

		m.zi = m.zri
		m.zr2.Mul(&m.zr, &m.zr)
		m.zi2.Mul(&m.zi, &m.zi)
	}
	// explicit not in set sentinel value
	m.val = -1.0
}

func (m *MandelbrotPerPixelDelegate) GetPixel(i, j int64) {
	m.SetMandelbrotCoordinate(i, j)

	m.MandelbrotKernel()

	if m.parent.OutIter != nil {
		m.parent.OutIter.Set(i, j, m.val)
	}

	if m.parent.OutImage != nil {
		if m.val >= 0.0 {
			m.val = math.Log2(m.val+1) / math.Log2(m.parent.MaxVal+1) * m.parent.Wrap
			m.val = math.Sin(m.val*2*math.Pi)/2.0 + 0.5
			m.parent.OutImage.Set64(i, j, m.parent.Cmap.GetColor(m.val))
		} else {
			m.parent.OutImage.Set64(i, j, color.RGBA{0, 0, 0, 255})
			return
		}
	}
}

func (m *MandelbrotPerPixel) Bounds() (w, h int64)         { return m.Width, m.Height }
func (m *MandelbrotPerPixelDelegate) Bounds() (w, h int64) { return m.parent.Bounds() }
