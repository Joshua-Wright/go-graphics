package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/parallel"
	"github.com/joshua-wright/go-graphics/Slice2D"
	"math"
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
	//width := 1920
	//height := 1080
	//blackThreshold := 256
	//nPtsPerX := 50000
	//nIter := 5000
	width := 1280
	height := 720
	blackThreshold := 256
	nPtsPerX := 8000
	nIter := 500

	//rMin := 3.0
	//rMin := 3.54409
	// do the full range, and then squish it sideways later
	rMin := 1.0
	rMax := 4.0

	grid := Slice2D.NewUint32Slice2D(width, height)
	parallel.ParallelFor(0, width, func(i int) {
		t := float64(i) / float64(width)
		r := g.Lerp(rMin, rMax, math.Pow(t, 1.0/8.0))
		for j := 0; j < nPtsPerX; j++ {
			// locate and iterate point
			x_init := float64(j) / float64(nPtsPerX)
			x_final := logistic_map_iter(r, x_init, nIter)

			// map to y-value, flipping
			x_mapped := height - 1 - int(x_final*float64(height))
			// increment count on grid
			grid.Set(i, x_mapped, grid.Get(i, x_mapped)+1)
		}
	})

	maxCount := uint32(0)
	for _, e := range grid.Data {
		if e > maxCount {
			maxCount = e
		}
	}

	out := g.CountBoxImage1(grid, uint32(blackThreshold), 1.0)

	g.SaveAsPNG(out, g.ExecutableNamePng())

}
