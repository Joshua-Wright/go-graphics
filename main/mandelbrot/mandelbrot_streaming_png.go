package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"github.com/joshua-wright/go-graphics/parallel"
	"image/color"
	"image"
	"math"
)

type streamingMandelbrot struct {
	outWidth, outHeight  int64
	pixelRequest         chan int64
	pixelRequestResponse chan color.Color
	pool                 *parallel.StreamingWorkPool
}

func (img *streamingMandelbrot) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *streamingMandelbrot) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(img.outWidth), int(img.outHeight))
}

func (img *streamingMandelbrot) At(x_, y_ int) color.Color {
	x := int64(x_)
	y := int64(y_)
	idx := y*img.outWidth + x
	return img.pool.Get(idx).(color.Color)
}

// if we tell the PNG encoder that we are opaque, it will traverse the image exactly once
func (*streamingMandelbrot) Opaque() bool {
	return true
}

func DownsamplePixel(ppm *memory_mapped.PPMFile, index, factor int64) color.RGBA {
	w := ppm.W / factor
	x := (index % w) * factor
	y := (index / w) * factor

	var rs float64
	var gs float64
	var bs float64

	for j := y; j < y+factor; j++ {
		for i := x; i < x+factor; i++ {
			r, g, b := ppm.At64(i, j)
			rs += float64(r)
			gs += float64(g)
			bs += float64(b)
		}
	}

	d := float64(factor * factor)

	return color.RGBA{
		R: uint8(rs / d),
		G: uint8(gs / d),
		B: uint8(bs / d),
		A: 255,
	}
}

func main() {
	factor := int64(4)
	width := int64(2560)
	height := int64(1440)
	maxIter := int64(25600)
	//factor := int64(2)
	//width := int64(192)
	//height := int64(108)
	//maxIter := int64(2560)

	wrap := 8.0
	maxVal := float64(maxIter)

	cmap := colormap.NewInterpColormap(colormap.InfernoColorMap())
	//cmap := colormap.NewInterpColormap(colormap.UltraFractalColors16)
	//bounds := MandelbrotBounds(width, height, complex(-0.7435669, 0.1314023), 1344.9)
	topLeft, dr, di := mandelbrot.MandelbrotBounds(width*factor, height*factor, complex(-0.74364085, 0.13182733), 25497*1.1)

	sampled := &streamingMandelbrot{
		outWidth:  width,
		outHeight: height,
		pool: parallel.MakeStreamingWorkPool(width*height*factor*factor,
			func(idx int64) (item parallel.WorkItem, err error) {
				x := (idx % width) * factor
				y := (idx / width) * factor

				var rs float64
				var gs float64
				var bs float64

				for j := y; j < y+factor; j++ {
					for i := x; i < x+factor; i++ {
						z := complex(0, 0)
						c := topLeft + complex(dr*float64(i), -di*float64(j))
						val := mandelbrot.IterateMandelbrot(z, c, 1000, maxIter)

						if val != 0.0 {
							val = math.Log2(val+1) / math.Log2(maxVal+1) * wrap
							val = math.Sin(val*2*math.Pi)/2.0 + 0.5
							c := cmap.GetColor(val)
							rs += float64(c.R)
							gs += float64(c.G)
							bs += float64(c.B)
						}
					}
				}

				d := float64(factor * factor)

				return color.RGBA{
					R: uint8(rs / d),
					G: uint8(gs / d),
					B: uint8(bs / d),
					A: 255,
				}, nil
			}),
	}
	g.SaveAsPNG(sampled, g.ExecutableNamePng())
}
