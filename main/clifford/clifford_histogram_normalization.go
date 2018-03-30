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

	//overbleed := 8.0

	nPreIter := 400
	nIter := 1<<27

	width := 2560
	height := 1440

	bound_width := 3.0
	bounds := [4]g.Float{
		-bound_width, bound_width,
		-bound_width, bound_width,
		//-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	nProcs := runtime.GOMAXPROCS(-1)

	buffers := make([]Slice2D.Uint16Slice2D, nProcs)
	for i := 0; i < len(buffers); i++ {
		buffers[i] = Slice2D.NewUint16Slice2D(width, height)
	}

	println("iterate points")
	parallel.ParallelFor(0, nProcs, func(workerIndex int) {
		counts := buffers[workerIndex]

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
			if newCount == math.MaxUint16 {
				panic("too big")
			}
			counts.Set(x, y, newCount)
		}
	})

	println("histogram normalize")
	cdfs := make([][]uint64, nProcs)
	parallel.ParallelFor(0, nProcs, func(i int) {
		cdfs[i] = make([]uint64, math.MaxUint16+1)
		cdf := cdfs[i]

		// count non-zeroes
		for _, e := range buffers[i].Data {
			if e != 0 {
				cdf[e]++
			}
		}

		// prefix sum to make it a cdf
		for i := 1; i < math.MaxUint16+1; i++ {
			cdf[i] += cdf[i-1]
		}
	})
	// reduce to single cdf
	cdf := make([]uint64, math.MaxUint16+1)
	for i := 0; i < math.MaxUint16+1; i++ {
		cdf[i] = 0
		for j := 0; j < nProcs; j++ {
			cdf[i] += cdfs[j][i]
		}
	}
	numNonEmptyPixels := cdf[math.MaxUint16]

	out := image.NewNRGBA(image.Rect(0, 0, width, height))

	println("add and normalize points")
	parallel.ParallelForAdaptive(0, width, func(startInclusive, endExclusive int) {
		for x := startInclusive; x < endExclusive; x++ {
			for y := 0; y < height; y++ {

				// sum contribution from each processor
				counti := uint16(0)
				for i := 0; i < nProcs; i++ {
					counti += buffers[i].Get(x, y)
				}
				if counti == 0 {
					out.Set(x, y, color.White)
				} else {
					count := float64(cdf[counti]-1) / float64(numNonEmptyPixels-1)
					// squish the colors a lot
					count = 1 - math.Pow(1-count, 1.0/6.0)
					if count < 0 || count > 1.0 {
						// bad
						println(count)
					}
					count *= 255
					//println(int(count))

					// normalize
					//count = count / maxCount * 255 * overbleed
					//if count > 255 {
					//	count = 255
					//}

					// set element in output image
					out.Set(x, y, color.Gray{255 - uint8(count)})
				}

			}
		}
	})

	println("save png")
	g.SaveAsPNG(out, g.ExecutableNamePng())
}
