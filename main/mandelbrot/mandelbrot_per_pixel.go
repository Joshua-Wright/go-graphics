package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot"
	"image/color"
	"math"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
)

type Mandelbrot struct {
	TopLeft       complex128
	Dr, Di        float64
	Width, Height int
	MaxIter       int
	MaxVal        float64
	Wrap          float64
}

func (m *Mandelbrot) GetPixel(i, j int) color.RGBA {
	z := complex(0, 0)
	c := m.TopLeft + complex(m.Dr*float64(i), m.Di*float64(j))
	val := mandelbrot.IterateMandelbrot(z, c, 1000, m.MaxIter)

	if val == 0.0 {
		return color.RGBA{0, 0, 0, 255}
	} else {
		val = math.Log2(val+1) / math.Log2(m.MaxVal+1) * m.Wrap
		val = math.Sin(val*2*math.Pi)/2.0 + 0.5

		return colormap.HotColormap(val)
	}

}

func (m *Mandelbrot) Bounds() (w int, h int) {
	return m.Width, m.Height
}

func main() {
	//Height := 1440*2
	//Width := int(Height*16.0/9.0) + 2*int(Height*4.0/5.0)
	upscaleFactor := 15
	width := 1920 * upscaleFactor
	height := 1080 * upscaleFactor
	maxIter := 25600

	bound_width := 2.0
	bounds := [4]g.Float{
		-bound_width, bound_width,
		//-bound_width, bound_width,
		-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	topLeft := complex(bounds[0], bounds[2])
	dr := (bounds[1] - bounds[0]) / float64(width)
	di := (bounds[3] - bounds[2]) / float64(height)

	err := per_pixel_image.PerPixelImage(
		&Mandelbrot{
			TopLeft: topLeft,
			Dr:      dr,
			Di:      di,
			Width:   width,
			Height:  height,
			MaxIter: maxIter,
			MaxVal:  float64(maxIter),
			Wrap:    2.0,
		},
		"mandelbrot")
	g.Die(err)
}
