package colormap

import (
	"image/color"
	"github.com/lucasb-eyer/go-colorful"
	"math"
)

type ColorMap interface {
	// 0 <= x < 1
	GetColor(x float64) color.RGBA
}

var UltraFractalColors16 = []color.Color{
	color.RGBA{66, 30, 15, 255},
	color.RGBA{25, 7, 26, 255},
	color.RGBA{9, 1, 47, 255},
	color.RGBA{4, 4, 73, 255},
	color.RGBA{0, 7, 100, 255},
	color.RGBA{12, 44, 138, 255},
	color.RGBA{24, 82, 177, 255},
	color.RGBA{57, 125, 209, 255},
	color.RGBA{134, 181, 229, 255},
	color.RGBA{211, 236, 248, 255},
	color.RGBA{241, 233, 191, 255},
	color.RGBA{248, 201, 95, 255},
	color.RGBA{255, 170, 0, 255},
	color.RGBA{204, 128, 0, 255},
	color.RGBA{153, 87, 0, 255},
	color.RGBA{106, 52, 3, 255},
}

type Hsv struct{}

func (Hsv) GetColor(x float64) color.RGBA {
	r, g, b := colorful.Hsv(360*x, 1.0, 1.0).RGB255()
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

type Sinebow struct{}

func (Sinebow) GetColor(h float64) color.RGBA {
	h = -(h+1.0/2.0)
	r := math.Sin(math.Pi * h)
	g := math.Sin(math.Pi * (h + 1.0/3.0))
	b := math.Sin(math.Pi * (h + 2.0/3.0))
	return color.RGBA{
		R: uint8(r * r * 255),
		G: uint8(g * g * 255),
		B: uint8(b * b * 255),
		A: 255,
	}

}
