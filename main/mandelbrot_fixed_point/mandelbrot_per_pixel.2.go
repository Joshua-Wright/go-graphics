package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
)

func main() {
	width := int64(512)
	height := int64(512)
	//centerR := "-1.7490868103567926827901785714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714286"
	//centerI := "0.00035485293981043127901785714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714"
	//zoom := "35840000000000000"
	centerR := "0.0"
	centerI := "0.0"
	zoom := "1.0"

	maxIter := int64(5120)
	threshold := "32.0"
	wrap := 12.0

	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())

	///

	outImage, err := memory_mapped.OpenOrCreatePPM(width, height, "mandelbrot.ppm")
	g.Die(err)

	outIter, err := memory_mapped.OpenOrCreateMmap2dArrayFloat64(width, height, "mandelbrot.iter")
	g.Die(err)

	err = per_pixel_image.PerPixelImage(
		mandelbrot_fixed_point.NewMandelbrotPerPixel(
			width, height,
			maxIter,
			centerR,
			centerI,
			zoom,
			threshold,
			wrap,
			float64(maxIter),
			cmap,
			outImage,
			outIter,
		),
		"mandelbrot.bitmap")
	g.Die(err)
}
