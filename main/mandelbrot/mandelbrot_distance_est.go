package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"math"
)

func main() {
	upscaleFactor := int64(4)
	width := 2560 * upscaleFactor
	height := 1440 * upscaleFactor
	maxIter := int64(25600)
	wrap := 8.0

	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())
	//cmap := colormap.NewInterpColormap(colormap.UltraFractalColors16)
	//bounds := MandelbrotBounds(width, height, complex(-0.7435669, 0.1314023), 1344.9)
	//topLeft, dr, di := mandelbrot.MandelbrotBounds(width, height, complex(-0.74364085, 0.13182733), 25497*1.1)
	topLeft, dr, di := mandelbrot.MandelbrotBounds(width, height, complex(0, 0), 1)

	outImage, err := memory_mapped.OpenOrCreatePPM(width, height, "mandelbrot.ppm")
	g.Die(err)

	outIter, err := memory_mapped.OpenOrCreateMmap2dArrayFloat64(width, height, "mandelbrot.iter")
	g.Die(err)

	err = per_pixel_image.PerPixelImage(
		&MandelbrotPerPixelDistanceEst{
			TopLeft:  topLeft,
			Dr:       dr,
			Di:       di,
			Width:    width,
			Height:   height,
			MaxIter:  maxIter,
			MaxVal:   float64(maxIter),
			Wrap:     wrap,
			Cmap:     cmap,
			OutImage: outImage,
			OutIter:  outIter,
		},
		"mandelbrot.bitmap")
	g.Die(err)
}

type MandelbrotPerPixelDistanceEst struct {
	TopLeft       complex128
	Dr, Di        float64
	Width, Height int64
	MaxIter       int64
	MaxVal        float64
	Wrap          float64
	Cmap          colormap.ColorMap
	OutImage      *memory_mapped.PPMFile
	OutIter       *memory_mapped.Array2dFloat64
}

func (m *MandelbrotPerPixelDistanceEst) GetPixel(i, j int64) {
	z := complex(0, 0)
	c := m.TopLeft + complex(m.Dr*float64(i), -m.Di*float64(j))
	val, dist := mandelbrot.IterateMandelbrotDistanceEst(z, c, 1000, m.MaxIter)

	// if the pixel intersects the set, consider it inside the set
	//if dist < m.Dr/2 {
	//	val = 0
	//}

	_ = val
	x := 0.5*math.Tanh(dist/m.Dr) + 0.5
	if math.IsNaN(x) {
		x = 0
	}
	m.OutImage.Set64(i, j, m.Cmap.GetColor(g.ScaleToNotIncludeOne(x)))

	//if m.OutIter != nil {
	//	m.OutIter.Set(i, j, val)
	//}
	//
	//if m.OutImage != nil {
	//	if val == 0.0 {
	//		m.OutImage.Set64(i, j, color.RGBA{0, 0, 0, 255})
	//		return
	//	} else {
	//		val = math.Log2(val+1) / math.Log2(m.MaxVal+1) * m.Wrap
	//		val = math.Sin(val*2*math.Pi)/2.0 + 0.5
	//		m.OutImage.Set64(i, j, m.Cmap.GetColor(val))
	//	}
	//}

}

func (m *MandelbrotPerPixelDistanceEst) Bounds() (w, h int64) { return m.Width, m.Height }
