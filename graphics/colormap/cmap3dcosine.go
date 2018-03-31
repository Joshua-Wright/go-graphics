package colormap

import (
	"image/color"
	"math"
)

var DefaultCosine3D = &Cosine3D{
	0.5, 0.5, 3.0, 2 * math.Pi, 0.0,
	0.5, 0.5, 3.0, 2 * math.Pi, 0.6,
	0.5, 0.5, 3.0, 2 * math.Pi, 1.0,
}

type Cosine3D struct {
	r0, r1, r2, r3, r4 float64
	g0, g1, g2, g3, g4 float64
	b0, b1, b2, b3, b4 float64
}

func (c *Cosine3D) GetColor(x float64) color.RGBA {
	r := c.r0 + c.r1*math.Cos(c.r2+x*c.r3+c.r4);
	g := c.g0 + c.g1*math.Cos(c.g2+x*c.g3+c.g4);
	b := c.b0 + c.b1*math.Cos(c.b2+x*c.b3+c.b4);
	return color.RGBA{
		R: uint8(255 * r),
		G: uint8(255 * g),
		B: uint8(255 * b),
		A: 255,
	}
}
