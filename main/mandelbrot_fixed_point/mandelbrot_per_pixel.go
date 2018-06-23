package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	mandelbrot "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
)

func main() {
	width := int64(512)
	height := int64(512)
	maxIter := int64(512)
	wrap := 8.0
	centerR := "-0.74364085"
	centerI := "0.13182733"
	zoom := "30047"
	threshold := "32.0"
	basePower2 := uint(128)

	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())

	///

	outImage, err := memory_mapped.OpenOrCreatePPM(width, height, "mandelbrot.ppm")
	g.Die(err)

	outIter, err := memory_mapped.OpenOrCreateMmap2dArrayFloat64(width, height, "mandelbrot.iter")
	g.Die(err)

	err = per_pixel_image.PerPixelImage(
		mandelbrot.NewMandelbrotPerPixel(
			width, height,
			maxIter,
			centerR,
			centerI,
			zoom,
			threshold,
			basePower2,
			wrap,
			float64(maxIter),
			cmap,
			outImage,
			outIter,
		),
		"mandelbrot.bitmap")
	g.Die(err)
}
