package main

import (
	"strings"
	"path/filepath"
	"image/color"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/parallel"
	"math"
	"gopkg.in/cheggaaa/pb.v1"
)

func main() {
	factor := int64(4)

	filename := "mandelbrot.iter"
	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + "_downsampled.ppm"

	floats, err := memory_mapped.OpenMmap2dArrayFloat64(filename)
	g.Die(err)
	factor2 := float64(factor * factor)
	width2 := floats.Width() / factor
	height2 := floats.Height() / factor

	outppm, err := memory_mapped.CreatePPM(width2, height2, newFilename)
	g.Die(err)

	if floats.Width()%factor != 0 {
		panic("bad width")
	}
	if floats.Height()%factor != 0 {
		panic("bad height")
	}

	cmap := colormap.NewInterpColormap(colormap.JetColorMap())

	MaxVal := 2560.0
	Wrap := 8.0
	gamma := 1.5

	bar := pb.New64(height2)
	bar.Start()
	parallel.ParallelFor(0, int(height2), func(y_ int) {
		y := int64(y_) * factor
		for x_ := int64(0); x_ < width2; x_++ {
			x := int64(x_) * factor

			// RGB with gamma correction correction

			var rs float64
			var gs float64
			var bs float64

			for j := y; j < y+factor; j++ {
				for i := x; i < x+factor; i++ {
					val := floats.Get(i, j)
					var c color.RGBA
					if val != 0.0 {
						val = math.Log2(val+1) / math.Log2(MaxVal+1) * Wrap
						//val = math.Log(100*val + 1)
						//val = math.Sqrt(val * Wrap)
						val = math.Sin(val*2*math.Pi)/2.0 + 0.5
						c = cmap.GetColor(val)
					}
					rs += math.Pow(float64(c.R)/256, gamma)
					gs += math.Pow(float64(c.G)/256, gamma)
					bs += math.Pow(float64(c.B)/256, gamma)
				}
			}
			outppm.Set64RGB(x/factor, y/factor,
				uint8(math.Pow(rs/factor2, 1/gamma)*256),
				uint8(math.Pow(gs/factor2, 1/gamma)*256),
				uint8(math.Pow(bs/factor2, 1/gamma)*256),
			)

			// XYZ linear

			//var xs float64
			//var ys float64
			//var zs float64
			//
			//for j := y; j < y+factor; j++ {
			//	for i := x; i < x+factor; i++ {
			//		val := floats.Get(i, j)
			//		if val != 0.0 {
			//			val = math.Log2(val+1) / math.Log2(MaxVal+1) * Wrap
			//			val = math.Sin(val*2*math.Pi)/2.0 + 0.5
			//			x, y, z := colorful.MakeColor(cmap.GetColor(val)).Xyz()
			//			xs += x
			//			ys += y
			//			zs += z
			//		}
			//	}
			//}
			//xs /= factor2
			//ys /= factor2
			//zs /= factor2
			//r, g, b := colorful.Xyz(xs, ys, zs).RGB255()
			//outppm.Set64RGB(x/factor, y/factor, r, g, b)
		}
		bar.Increment()
	})
	bar.Finish()
}
