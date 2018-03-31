package main

import (
	"os"
	"strings"
	"path/filepath"
	"image/png"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	g "github.com/joshua-wright/go-graphics/graphics"
	"image/color"
	"strconv"
	"image"
	"github.com/joshua-wright/go-graphics/parallel"
)

type downsampledPpm struct {
	outWidth, outHeight  int64
	pixelRequest         chan int64
	pixelRequestResponse chan color.Color
	pool                 *parallel.StreamingWorkPool
}

func (img *downsampledPpm) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *downsampledPpm) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(img.outWidth), int(img.outHeight))
}

func (img *downsampledPpm) At(x_, y_ int) color.Color {
	x := int64(x_)
	y := int64(y_)
	idx := y*img.outWidth + x
	return img.pool.Get(idx).(color.Color)
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
	filename := os.Args[1]
	factor, err := strconv.Atoi(os.Args[2])
	g.Die(err)

	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".png"

	ppmImage, err := memory_mapped.OpenPPM(filename)
	g.Die(err)

	if ppmImage.W%int64(factor) != 0 {
		panic("bad width")
	}
	if ppmImage.H%int64(factor) != 0 {
		panic("bad height")
	}

	sampled := &downsampledPpm{
		outWidth:  ppmImage.W / int64(factor),
		outHeight: ppmImage.H / int64(factor),
		pool: parallel.MakeStreamingWorkPool(ppmImage.W/int64(factor)*ppmImage.H/int64(factor),
			func(i int64) (item parallel.WorkItem, err error) {
				px := DownsamplePixel(ppmImage, i, int64(factor))
				return px, nil
			}),
	}

	file, err := os.Create(newFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, sampled)
	if err != nil {
		panic(err)
	}
}
