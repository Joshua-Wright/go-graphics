package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/parallel"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"gopkg.in/cheggaaa/pb.v1"
	"image"
	"math"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot"
)

func main() {
	//height := 1440*2
	//width := int(height*16.0/9.0) + 2*int(height*4.0/5.0)
	width := 1920 * 2
	height := 1080 * 2
	//darkPerPoint := uint16(32)
	maxIter := 256
	wrap := 2.0

	bound_width := 2.0
	bounds := [4]g.Float{
		-bound_width, bound_width,
		//-bound_width, bound_width,
		-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	buf := mandelbrot.Mandelbrot(bounds, width, height, maxIter)

	println("total")
	maxVal := 0.0
	for _, v := range buf.Data {
		if v > maxVal {
			maxVal = v
		}
	}

	println("map/colorize")
	reduceBar := pb.StartNew(width)
	img := image.NewPaletted(image.Rect(0, 0, width, height), colormap.InfernoColorMap())
	parallel.ParallelFor(0, width, func(i int) {
		for j := 0; j < height; j++ {
			val := buf.Get(i, j)
			if val == 0.0 {
				img.SetColorIndex(i, j, uint8(val))
			} else {
				val = math.Log2(val+1) / math.Log2(maxVal+1) * wrap
				val = math.Sin(val * 2 * math.Pi)/2.0 + 0.5
				val = val * 255
				if val > 255 {
					val = 255
				}
				img.SetColorIndex(i, j, uint8(val))
			}
		}
		reduceBar.Increment()
	})
	reduceBar.Finish()

	println("write output")
	g.SaveAsPNG(img, g.ExecutableNamePng())

}
