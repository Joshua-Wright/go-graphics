//+build cgo

package mandelbrot_fixed_point

import (
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/gmp_fixnum"
)

func NewMandelbrotPerPixelGmp(width, height, maxIter int64,
	centerR, centerI, zoom, threshold string, bits uint,
	Wrap, MaxVal float64, cmap colormap.ColorMap,
	OutImage *memory_mapped.PPMFile,
	OutIter *memory_mapped.Array2dFloat64) per_pixel_image.PixelFunction {

	// use gmp
	return gmp_fixnum.NewMandelbrotPerPixel(
		width, height, maxIter, centerR, centerI, zoom, threshold, bits, Wrap, MaxVal, cmap, OutImage, OutIter)
}
