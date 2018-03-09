package graphics

import (
	"github.com/joshua-wright/go-graphics/Slice2D"
	"github.com/joshua-wright/go-graphics/parallel"
	"image"
	"image/color"
)

// automatically find max count
func CountBoxImage0(grid Slice2D.Uint32Slice2D, overbleed float64) *image.NRGBA {
	maxCount := uint32(0)
	for _, e := range grid.Data {
		if e > maxCount {
			maxCount = e
		}
	}
	return CountBoxImage1(grid, maxCount, overbleed)
}

// using 32 bit integers to reduce memory consumption
func CountBoxImage1(grid Slice2D.Uint32Slice2D, maxCount uint32, overbleed float64) *image.NRGBA {
	width := grid.W
	height := grid.H
	out := image.NewNRGBA(image.Rect(0, 0, width, height))

	maxCountf := float64(maxCount)

	// add and normalize points
	parallel.ParallelForAdaptive(0, width, func(startInclusive, endExclusive int) {
		for x := startInclusive; x < endExclusive; x++ {
			for y := 0; y < height; y++ {

				count := float64(grid.Get(x, y))

				// normalize
				count = count / maxCountf * 255 * overbleed
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
