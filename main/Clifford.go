package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/parallel"
	"github.com/joshua-wright/go-graphics/Slice2D"
	"image/color"
	"math/rand"
	"math"
	"runtime"
	"image"
)

func Clifford(a, b, c, d float64, v g.Vec2) g.Vec2 {
	return g.Vec2{
		X: math.Sin(a*v.Y) + c*math.Cos(a*v.X),
		Y: math.Sin(b*v.X) + d*math.Cos(b*v.Y),
	}
}

func main() {
	a := -1.33980582524272
	b := -2.11650485436893
	c := -1.84466019417476
	d := 1.12621359223301

	overbleed := 8.0

	nPreIter := 400
	nIter := 20000000

	width := 2560
	height := 1440

	bound_width := 3.0
	bounds := [4]g.Float{
		// a little extra space around the rose
		-bound_width, bound_width,
		-bound_width, bound_width,
		//-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	nProcs := runtime.GOMAXPROCS(-1)

	buffers := make([]Slice2D.Uint16Slice2D, nProcs)
	for i := 0; i < len(buffers); i++ {
		buffers[i] = Slice2D.NewUint16Slice2D(width, height)
	}
	maxes := make([]uint16, nProcs)

	// iterate points
	parallel.ParallelFor(0, nProcs, func(workerIndex int) {
		counts := buffers[workerIndex]
		var max_count uint16

		pt := g.Vec2{
			X: rand.Float64()*2 - 1,
			Y: rand.Float64()*2 - 1,
		}
		for i := 0; i < nPreIter; i++ {
			pt = Clifford(a, b, c, d, pt)
		}

		for i := 0; i < nIter; i++ {
			pt = Clifford(a, b, c, d, pt)
			x, y := g.WindowTransformPoint(width, height, pt, bounds)
			if x >= width || x < 0 || y >= height || y < 0 {
				println(x, y)
			}
			newCount := counts.Get(x, y) + 1
			counts.Set(x, y, newCount)
			if newCount > max_count {
				max_count = newCount
			}
		}

		maxes[workerIndex] = max_count
	})

	// find max point count
	maxCount := 0.0
	for _, newCount := range maxes {
		if float64(newCount) > maxCount {
			maxCount = float64(newCount)
		}
	}

	out := image.NewNRGBA(image.Rect(0, 0, width, height))

	// add and normalize points
	parallel.ParallelForAdaptive(0, width, func(startInclusive, endExclusive int) {
		for x := startInclusive; x < endExclusive; x++ {
			for y := 0; y < height; y++ {

				// sum contribution from each processor
				count := 0.0
				for i := 0; i < nProcs; i++ {
					count += float64(buffers[i].Get(x, y))
				}

				// normalize
				count = count / maxCount * 255 * overbleed
				if count > 255 {
					count = 255
				}

				// set element in output image
				out.Set(x, y, color.Gray{255 - uint8(count)})
			}
		}
	})

	g.SaveAsPNG(out, g.ExecutableNamePng())

}
