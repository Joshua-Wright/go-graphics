package main

import (
	m ".."
	"math"
	"fmt"
)

func main() {

	// golden ratio
	phi := 1.618033988749894848204586834365638117720309179805762862135
	r := math.Pow(1.0/phi, 1.0/phi)
	fmt.Println("r =", r)

	//mats := []m.Matrix3{
	//	{
	//		{0.62367, -0.40337, 0},
	//		{0.40337, 0.62367, 0},
	//		{0, 0, 1},
	//	},
	//	{
	//		{-0.37633, -0.40337, 1},
	//		{0.40337, -0.37633, 0},
	//		{0, 0, 1},
	//	},
	//}
	mats := []m.Matrix3{
		m.Scale2D(0, 0, r).Compose(m.Rotate2D(0.0, 0.0, 32.893818*math.Pi/180.0)),
		m.Scale2D(1, 0, r*r).Compose(m.Rotate2D(1.0, 0.0, 133.0140178*math.Pi/180.0)),
	}

	width := 1000
	//xmid := 0.4
	//ymid := 0.0
	//dw := 0.75
	xmid := 0.6
	ymid := 0.0
	dw := 1.0

	m.RenderFractal0(mats, "golden_dragon.png", 25, width, [4]m.Float{xmid - dw, xmid + dw, ymid - dw, ymid + dw})
	//m.RenderFractal0(mats, "golden_dragon.png", 25, width, m.DefaultFractalBounds)
}
