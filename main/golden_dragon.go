package main

import (
	m "github.com/joshua-wright/image_stuff/go"
)

func main() {

	mats := []m.Matrix3{
		{
			{0.62367, -0.40337, 0},
			{0.40337, 0.62367, 0},
			{0, 0, 1},
		},
		{
			{-0.37633, -0.40337, 1},
			{0.40337, -0.37633, 0},
			{0, 0, 1},
		},
	}

	width := 1000
	xmid := 0.4
	ymid := 0.0
	dw := 0.75

	m.RenderFractal0(mats, "golden_dragon.png", 25, width, [4]m.Float{xmid - dw, xmid + dw, ymid - dw, ymid + dw})
}
