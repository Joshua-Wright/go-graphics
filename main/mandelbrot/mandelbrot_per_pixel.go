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
	Cmap          *colormap.XyzInterpColormap
}

func (m *Mandelbrot) GetPixel(i, j int) color.RGBA {
	z := complex(0, 0)
	c := m.TopLeft + complex(m.Dr*float64(i), -m.Di*float64(j))
	val := mandelbrot.IterateMandelbrot(z, c, 1000, m.MaxIter)

	if val == 0.0 {
		return color.RGBA{0, 0, 0, 255}
	} else {
		val = math.Log2(val+1) / math.Log2(m.MaxVal+1) * m.Wrap
		val = math.Sin(val*2*math.Pi)/2.0 + 0.5

		return m.Cmap.GetColor(val)
	}

}

func (m *Mandelbrot) Bounds() (w int, h int) { return m.Width, m.Height }

func calc_pixel_widths(x, y int, zoom float64) (float64, float64) {
	dx := 2.0 / zoom;
	dy := 2.0 / zoom;
	if x > y {
		/* widescreen image */
		dx = float64(x) / float64(y) * dy;
	} else if (y > x) {
		/* portrait */
		dy = float64(y) / float64(x) * dx;
	} // otherwise square
	return dx / float64(x), dy / float64(y);
}

func calc_bounds(x, y int, center complex128, zoom float64) [4]float64 {
	dx := 2.0 / zoom;
	dy := 2.0 / zoom;
	if x > y {
		/* widescreen image */
		dx = float64(x) / float64(y) * dy;
	} else if (y > x) {
		/* portrait */
		dy = float64(y) / float64(x) * dx;
	} // otherwise square
	return [4]float64{
		real(center) - dx, real(center) + dx, imag(center) - dy, imag(center) + dy,
	};
}

func main() {
	upscaleFactor := 2
	width := 1920 * upscaleFactor
	height := 1080 * upscaleFactor
	maxIter := 25600

	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())

	bounds := calc_bounds(width, height, complex(-0.7435669, 0.1314023), 1344.9)
	topLeft := complex(bounds[0], bounds[3])
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
			Wrap:    4.0,
			Cmap:    cmap,
		},
		"mandelbrot")
	g.Die(err)
}
