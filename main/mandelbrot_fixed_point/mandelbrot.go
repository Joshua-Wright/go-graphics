package main

import (
	"image"
	"github.com/ncw/gmp"
	"github.com/joshua-wright/go-graphics/parallel"
	m "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"image/color"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"gopkg.in/cheggaaa/pb.v1"
)

func main() {
	width := int64(1024)
	height := int64(512)
	maxIter := int64(512)
	wrap := 8.0
	base_power_2 := uint(64)

	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	threshold2 := gmp.NewInt(2)
	threshold2.Lsh(threshold2, base_power_2)
	threshold2.Mul(threshold2, threshold2)

	zoom := gmp.NewInt(1)
	centerR := gmp.NewInt(0)
	centerI := gmp.NewInt(0)
	//zoom := gmp.NewInt(28047)
	//centerR := m.ParseFixnum("-0.74364085", base_power_2)
	//centerI := m.ParseFixnum("0.13182733", base_power_2)

	bar := pb.StartNew(int(height))
	parallel.ParallelFor(0, int(height), func(j_ int) {
		j := int64(j_)
		//for j := int64(0); j < height; j++ {
		//parallel.ParallelFor(0, int(width), func(i_ int) {
		//	i := int64(i_)
		for i := int64(0); i < width; i++ {

			cr, ci := m.MandelbrotCoordinate(i, j, width, height, centerR, centerI, zoom, base_power_2)

			val := m.MandelbrotKernel(cr, ci, threshold2, maxIter, base_power_2)
			var col color.Color
			if val >= 0 {
				val = math.Log2(val+1) / math.Log2(float64(maxIter)+1) * wrap
				val = math.Sin(val*2*math.Pi)/2.0 + 0.5
				col = cmap.GetColor(val)
			} else {
				col = color.Black
			}
			img.Set(int(i), int(j), col)
			//}
		}
		bar.Increment()
	})
	bar.Finish()

	g.SaveAsPNG(img, g.ExecutableNamePng())
}
