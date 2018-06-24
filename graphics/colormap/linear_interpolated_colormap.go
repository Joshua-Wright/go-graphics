package colormap

import (
	"github.com/joshua-wright/go-graphics/graphics/interpolation"
	"image/color"
	"github.com/lucasb-eyer/go-colorful"
	"math"
)

type XyzInterpColormap struct {
	x, y, z *interpolation.CubicEquidistant
}

func NewXyzInterpColormap(colors []color.Color) *XyzInterpColormap {
	nPts := len(colors)
	xs := make([]float64, nPts)
	ys := make([]float64, nPts)
	zs := make([]float64, nPts)
	for i := 0; i < nPts; i++ {
		c, _ := colorful.MakeColor(colors[i])
		x, y, z := c.Xyz()
		xs[i] = x
		ys[i] = y
		zs[i] = z
	}

	icm := XyzInterpColormap{}
	icm.x = interpolation.NewCubicEquidistantInterpolatorFromSlices(0, math.Nextafter(1.0, 0.0), nPts, xs)
	icm.y = interpolation.NewCubicEquidistantInterpolatorFromSlices(0, math.Nextafter(1.0, 0.0), nPts, ys)
	icm.z = interpolation.NewCubicEquidistantInterpolatorFromSlices(0, math.Nextafter(1.0, 0.0), nPts, zs)
	return &icm
}

// 0 <= x < 1
func (icm *XyzInterpColormap) GetColor(x float64) color.RGBA {
	r, g, b := colorful.Xyz(
		icm.x.Get(x),
		icm.y.Get(x),
		icm.z.Get(x),
	).RGB255()
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}
