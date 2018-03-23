package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/file_backed_image"
)

func main() {
	upscaleFactor := int64(2)
	width := 1920 * upscaleFactor
	height := 1080 * upscaleFactor
	maxIter := int64(25600)
	wrap := 4.0
	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())
	//bounds := MandelbrotBounds(width, height, complex(-0.7435669, 0.1314023), 1344.9)
	topLeft, dr, di := mandelbrot.MandelbrotBounds(width, height, complex(-0.743643900055, 0.131825890901), 62407000*1.1)


	outImage, err := file_backed_image.OpenOrCreatePPM(width, height, "mandelbrot.ppm")
	g.Die(err)

	err = per_pixel_image.PerPixelImage(
		&mandelbrot.MandelbrotPerPixel{
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
		},
		"mandelbrot.bitmap")
	g.Die(err)
}
