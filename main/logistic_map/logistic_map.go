package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/parallel"
	"gopkg.in/cheggaaa/pb.v1"
	"image"
	"image/draw"
	"math"
	"math/rand"
)

func logistic_map(r, x float64) float64 {
	return r * x * (1.0 - x)
}

func logistic_map_iter(r, x_init float64, nIter int) (x float64) {
	x = x_init
	for i := 0; i < nIter; i++ {
		x = logistic_map(r, x)
	}
	return x
}

func main() {
	height := 1440
	width := int(height*16.0/9.0) + 2*int(height*4.0/5.0)
	//width := 1280
	//height := 720
	darkPerPoint := uint8(1)
	nPtsPerX := 32000
	nIter := 5000
	exponent := 1 / 5.0

	//rMin := 3.0
	//rMin := 3.54409
	// do the full range, and then squish it sideways later
	rMin := 1.0
	rMax := 4.0

	println("allocate")
	img := image.NewGray(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Over)

	println("calculate")
	bar := pb.StartNew(width)
	parallel.ParallelFor(0, width, func(i int) {
		// choose a random t within this window
		t := float64(i) / float64(width) + (rand.Float64()-0.5)/float64(width)
		r := g.Lerp(rMin, rMax, math.Pow(t, exponent))
		for j := 0; j < nPtsPerX; j++ {
			// locate point
			x_init := float64(j) / float64(nPtsPerX)
			// iterate the point
			x_final := logistic_map_iter(r, x_init, nIter)

			yval := int(float64(x_final) * float64(height))
			c := img.GrayAt(i, height-1-yval)
			if c.Y > darkPerPoint {
				c.Y -= darkPerPoint
			}
			img.SetGray(i, height-1-yval, c)
		}
		bar.Increment()
	})
	bar.Finish()

	println("write output")
	g.SaveAsPNG(img, g.ExecutableNamePng())

}
