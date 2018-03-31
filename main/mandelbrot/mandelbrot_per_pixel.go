package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
)

func main() {
	upscaleFactor := int64(8)
	width := 2560 * upscaleFactor
	height := 1440 * upscaleFactor
	maxIter := int64(25600)
	wrap := 8.0

	//rs := interpolation.NewCubicInterpolator(0, 1, []float64{0.0, 0.16, 0.42, 0.6425, 0.8575, 1}, []float64{0, 32, 237, 255, 0, 0})
	//gs := interpolation.NewCubicInterpolator(0, 1, []float64{0.0, 0.16, 0.42, 0.6425, 0.8575, 1}, []float64{7, 107, 255, 170, 2, 7})
	//bs := interpolation.NewCubicInterpolator(0, 1, []float64{0.0, 0.16, 0.42, 0.6425, 0.8575, 1}, []float64{100, 203, 255, 0, 0, 100})
	//
	//cs := make([]color.Color, 256)
	//for i := 0; i < 256; i++ {
	//	cs[i] = color.RGBA{
	//		R: uint8(rs.Get(float64(i) / float64(256))),
	//		G: uint8(gs.Get(float64(i) / float64(256))),
	//		B: uint8(bs.Get(float64(i) / float64(256))),
	//	}
	//}
	//cmap := colormap.NewInterpColormap(cs)
	//cmap := Hsv{}
	//cmap := colormap.Sinebow{}

	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())
	//cmap := colormap.NewInterpColormap(colormap.UltraFractalColors16)
	//bounds := MandelbrotBounds(width, height, complex(-0.7435669, 0.1314023), 1344.9)
	topLeft, dr, di := mandelbrot.MandelbrotBounds(width, height, complex(-0.74364085, 0.13182733), 25497*1.1)

	outImage, err := memory_mapped.OpenOrCreatePPM(width, height, "mandelbrot.ppm")
	g.Die(err)

	outIter, err := memory_mapped.OpenOrCreateMmap2dArrayFloat64(width, height, "mandelbrot.iter")
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
			OutIter:  outIter,
		},
		"mandelbrot.bitmap")
	g.Die(err)
}
