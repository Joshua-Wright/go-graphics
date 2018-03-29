package main

import (
	"os"
	"strings"
	"path/filepath"
	"image/png"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	g "github.com/joshua-wright/go-graphics/graphics"
	"image/color"
	"image"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/parallel"
	"math"
)

type downsampledColorizedIterations struct {
	factor int64
	floats *memory_mapped.Array2dFloat64
	cmap   colormap.ColorMap
	MaxVal float64
	Wrap   float64
}

func (img *downsampledColorizedIterations) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *downsampledColorizedIterations) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(img.floats.Width()/img.factor), int(img.floats.Height()/img.factor))
}

func (img *downsampledColorizedIterations) At(x_, y_ int) color.Color {
	x := int64(x_) * img.factor
	y := int64(y_) * img.factor

	var rs float64
	var gs float64
	var bs float64

	parallel.ParallelFor(0, int(img.factor), func(offset int) {
		j := y + int64(offset)
		for i := x; i < x+img.factor; i++ {
			val := img.floats.Get(i, j)
			var c color.RGBA
			if val != 0.0 {
				val = math.Log2(val+1) / math.Log2(img.MaxVal+1) * img.Wrap
				val = math.Sin(val*2*math.Pi)/2.0 + 0.5
				c = img.cmap.GetColor(val)
			}
			rs += float64(c.R)
			gs += float64(c.G)
			bs += float64(c.B)
		}
	})

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

	iterations, err := memory_mapped.OpenMmap2dArrayFloat64(filename)
	g.Die(err)

	//factor, err := strconv.Atoi(os.Args[2])
	//g.Die(err)
	factor := 2

	if iterations.Width()%int64(factor) != 0 {
		panic("bad width")
	}
	if iterations.Height()%int64(factor) != 0 {
		panic("bad height")
	}

	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())

	sampled := downsampledColorizedIterations{
		factor: int64(factor),
		floats: iterations,
		cmap:   cmap,
		MaxVal: 2560,
		Wrap:   4.0,
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
