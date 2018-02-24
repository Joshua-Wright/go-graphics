package main

import (
	"github.com/joshua-wright/go-graphics/graphics/fractals"
	g "github.com/joshua-wright/go-graphics/graphics"
)

func main() {
	a := -1.33980582524272
	b := -2.11650485436893
	c := -1.84466019417476
	d := 1.12621359223301

	overbleed := 8.0

	nPreIter := 400
	nIter := 20000000

	width := 1280
	height := 820

	bound_width := 3.0
	bounds := [4]g.Float{
		// a little extra space around the rose
		-bound_width, bound_width,
		-bound_width, bound_width,
		//-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	out := fractals.CliffordFractal(width, height, a, b, c, d, overbleed, bounds, nPreIter, nIter)

	g.SaveAsPNG(out, g.ExecutableNamePng())

}
