package mandelbrot_fixed_point

import (
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"strconv"
	"math"
	"fmt"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot"
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/gmp_fixnum"
)

func NewMandelbrotPerPixel(width, height, maxIter int64,
	centerR, centerI, zoom, threshold string,
	Wrap, MaxVal float64, cmap colormap.ColorMap,
	OutImage *memory_mapped.PPMFile,
	OutIter *memory_mapped.Array2dFloat64) per_pixel_image.PixelFunction {

	// calculate how many bits of precision we need
	zoomf, err := strconv.ParseFloat(zoom, 64)
	g.Die(err)
	bits := uint(math.Log2(zoomf*math.Max(float64(width), float64(height)))) + 8 // +8 just in case

	if bits < 53 {
		fmt.Println("using machine precision")

		cr, err := strconv.ParseFloat(centerR, 64)
		g.Die(err)
		ci, err := strconv.ParseFloat(centerI, 64)
		g.Die(err)
		zoom, err := strconv.ParseFloat(zoom, 64)
		g.Die(err)

		topLeft, dr, di := mandelbrot.MandelbrotBounds(width, height, complex(cr, ci), zoom)
		return &mandelbrot.MandelbrotPerPixel{
			TopLeft:  topLeft,
			Dr:       dr,
			Di:       di,
			Width:    width,
			Height:   height,
			MaxIter:  maxIter,
			MaxVal:   float64(maxIter),
			Wrap:     Wrap,
			Cmap:     cmap,
			OutImage: OutImage,
			OutIter:  OutIter,
		}
	}

	// try to use naive fixnum
	words := (bits+31)/32 + 1
	c, ok := naive_fixnum.FixnumConstructorMap[words]
	if ok {
		fmt.Println("usng", words, "words (need", bits, "bits)")
		return c(width, height, maxIter, centerR, centerI, zoom, threshold, Wrap, MaxVal, cmap, OutImage, OutIter)
	}

	// use gmp
	return gmp_fixnum.NewMandelbrotPerPixel(
		width, height, maxIter, centerR, centerI, zoom, threshold, uint(bits), Wrap, MaxVal, cmap, OutImage, OutIter)
}
