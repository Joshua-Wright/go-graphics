package main

import (
	"github.com/joshua-wright/go-graphics/graphics/fractals"
	g "github.com/joshua-wright/go-graphics/graphics"
	"fmt"
	"github.com/joshua-wright/go-graphics/parallel"
	"math"
)

func main() {
	a := -1.24458046630025
	b := -1.25191834103316
	c := -1.81590817030519
	d := -1.90866735205054

	//maxCount := 17050 / 12.0
	//maxCount := 17050.0 / 12.0 / 8.0
	maxCount := 300.0

	nPreIter := 400
	nIter := 60000000

	width := 800
	height := 800

	//nFrames := 12
	nFrames := 1200

	bound_width := 3.0
	bounds := [4]g.Float{
		-bound_width, bound_width,
		-bound_width, bound_width,
		//-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	parallel.ParallelFor(0, nFrames, func(frame int) {
		a2 := math.Sin(float64(frame)/float64(nFrames)*2*math.Pi) * 0.01
		b2 := math.Cos(float64(frame)/float64(nFrames)*2*math.Pi) * 0.01
		out := fractals.CliffordFractalSerial(width, height, a+a2, b+b2, c, d, maxCount, bounds, nPreIter, nIter)
		g.SaveAsPNG(out, g.ExecutableFolderFileName(fmt.Sprint(frame)+".png"))
		println(frame)
	})

}
