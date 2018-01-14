package distance_transform

import (
	g "github.com/joshua-wright/go-graphics/graphics"
)

func Make2DSliceVec2(x, y int, val g.Vec2) [][]g.Vec2 {
	base := make([]g.Vec2, x*y)
	out := make([][]g.Vec2, x)

	for i := 0; i < x*y; i++ {
		base[i] = val
	}

	for i := 0; i < x; i++ {
		out[i] = base[i*x:(i+1)*x]
	}
	return out
}

func Make2DSliceFloat64(x, y int, val float64) [][]float64 {
	base := make([]float64, x*y)
	out := make([][]float64, x)

	for i := 0; i < x*y; i++ {
		base[i] = val
	}

	for i := 0; i < x; i++ {
		out[i] = base[i*x:(i+1)*x]
	}
	return out
}

func Make2DSlicePixelParallel(x, y int, val pixelParallel) [][]pixelParallel {
	base := make([]pixelParallel, x*y)
	out := make([][]pixelParallel, x)

	for i := 0; i < x*y; i++ {
		base[i] = val
	}

	for i := 0; i < x; i++ {
		out[i] = base[i*x:(i+1)*x]
	}
	return out
}
