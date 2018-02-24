package fractals

import (
	"runtime"
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/Slice2D"
	"github.com/joshua-wright/go-graphics/parallel"
	"math/rand"
	"image/color"
	"image"
	"math"
)

func Clifford(a, b, c, d float64, v g.Vec2) g.Vec2 {
	return g.Vec2{
		X: math.Sin(a*v.Y) + c*math.Cos(a*v.X),
		Y: math.Sin(b*v.X) + d*math.Cos(b*v.Y),
	}
}

func CliffordFractalCounts(
	width, height int,
	a, b, c, d g.Float,
	bounds [4]g.Float,
	nPreIter, nIter int) (grid Slice2D.Uint16Slice2D, max uint16) {
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
				panic(g.Vec2i{x, y});
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
	maxCount := uint16(0)
	for _, newCount := range maxes {
		if newCount > maxCount {
			maxCount = newCount
		}
	}

	outBuffer := Slice2D.NewUint16Slice2D(width, height)

	// aggregate maximums
	parallel.ParallelForAdaptive(0, width, func(startInclusive, endExclusive int) {
		for x := startInclusive; x < endExclusive; x++ {
			for y := 0; y < height; y++ {

				// sum contribution from each processor
				count := buffers[0].Get(x, y)
				for i := 1; i < nProcs; i++ {
					count += buffers[i].Get(x, y)
				}

				outBuffer.Set(x, y, count)
			}
		}
	})

	return outBuffer, maxCount
}

func CliffordFractal(
	width, height int,
	a, b, c, d,
	overbleed g.Float,
	bounds [4]g.Float,
	nPreIter, nIter int) *image.NRGBA {

	buffer, maxCount := CliffordFractalCounts(width, height, a, b, c, d, bounds, nPreIter, nIter);
	out := image.NewNRGBA(image.Rect(0, 0, width, height))

	// add and normalize points
	parallel.ParallelForAdaptive(0, width, func(startInclusive, endExclusive int) {
		for x := startInclusive; x < endExclusive; x++ {
			for y := 0; y < height; y++ {

				// normalize
				count := g.Float(buffer.Get(x, y)) / g.Float(maxCount) * 255 * overbleed
				if count > 255 {
					count = 255
				}

				// set element in output image
				out.Set(x, y, color.Gray{255 - uint8(count)})
			}
		}
	})
	return out
}
