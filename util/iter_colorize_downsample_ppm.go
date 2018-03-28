package main

import (
	"os"
	"strings"
	"path/filepath"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"strconv"
	g "github.com/joshua-wright/go-graphics/graphics"
	"image/color"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/parallel"
	"math"
)

func main() {
	filename := os.Args[1]
	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + "_downsampled.ppm"

	floats, err := memory_mapped.OpenMmap2dArrayFloat64(filename)
	g.Die(err)

	factor_, err := strconv.Atoi(os.Args[2])
	g.Die(err)
	factor := int64(factor_)
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

	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())

	MaxVal := 2560.0
	Wrap := 4.0

	parallel.ParallelFor(0, int(height2), func(y_ int) {
		y := int64(y_) * factor
		for x_ := int64(0); x_ < width2; x_++ {
			x := int64(x_) * factor

			var rs float64
			var gs float64
			var bs float64

			for j := y; j < y+factor; j++ {
				for i := x; i < x+factor; i++ {
					val := floats.Get(i, j)
					var c color.RGBA
					if val != 0.0 {
						val = math.Log2(val+1) / math.Log2(MaxVal+1) * Wrap
						val = math.Sin(val*2*math.Pi)/2.0 + 0.5
						c = cmap.GetColor(val)
					}
					rs += float64(c.R)
					gs += float64(c.G)
					bs += float64(c.B)
				}
			}
			outppm.Set64RGB(x/factor, y/factor,
				uint8(rs/factor2), uint8(gs/factor2), uint8(bs/factor2))
		}
	})
}
