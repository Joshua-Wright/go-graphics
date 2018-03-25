package main

import (
	"os"
	"strings"
	"path/filepath"
	"image/png"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"strconv"
	g "github.com/joshua-wright/go-graphics/graphics"
	"image/color"
	"image"
)

type downsampledPpm struct {
	factor int64
	ppm    *memory_mapped.PPMFile
}

func (img *downsampledPpm) ColorModel() color.Model {
	return img.ppm.ColorModel()
}

func (img *downsampledPpm) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(img.ppm.W/img.factor), int(img.ppm.H/img.factor))
}

func (img *downsampledPpm) At(x_, y_ int) color.Color {
	x := int64(x_) * img.factor
	y := int64(y_) * img.factor

	var rs float64
	var gs float64
	var bs float64

	for j := y; j < y+img.factor; j++ {
		for i := x; i < x+img.factor; i++ {
			r, g, b := img.ppm.At64(i, j)
			rs += float64(r)
			gs += float64(g)
			bs += float64(b)
		}
	}

	d := float64(img.factor * img.factor)
	return color.RGBA{
		R: uint8(rs / d),
		G: uint8(gs / d),
		B: uint8(bs / d),
		A: 255,
	}
}

func main() {
	filename := os.Args[1]
	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".png"

	ppmImage, err := memory_mapped.OpenPPM(filename)
	g.Die(err)

	factor, err := strconv.Atoi(os.Args[2])
	g.Die(err)

	if ppmImage.W%int64(factor) != 0 {
		panic("bad width")
	}
	if ppmImage.H%int64(factor) != 0 {
		panic("bad height")
	}

	sampled := downsampledPpm{
		factor: int64(factor),
		ppm:    ppmImage,
	}

	file, err := os.Create(newFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, &sampled)
	if err != nil {
		panic(err)
	}
}
