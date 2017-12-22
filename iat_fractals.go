package golang_raytracer

import (
	"os"
	"image/png"
	"math"
	"sync"
)

func TransformPoints(pts []Vec2, mats []Matrix3, max_depth int) []Vec2 {
	// make sure that we don't do parallel computation unless the data size
	// is larger than some threshold
	threshold := 256
	for {
		newpts := make([]Vec2, 0, len(mats)*len(pts))

		for _, m := range mats {
			for _, p := range pts {
				newpts = append(newpts, m.TransformPoint(&p))
			}
		}
		if max_depth == 0 {
			return newpts
		} else if len(newpts) > threshold {
			max_depth -= 1
			return transformPointsParallelImpl(newpts, mats, max_depth)
		} else {
			max_depth -= 1
			pts = newpts
		}
	}
}

func transformPointsParallelImpl(pts []Vec2, mats []Matrix3, max_depth int) []Vec2 {
	// points at each depth, pre-allocated for speed
	n_pts := len(pts)
	n_mats := len(mats)
	depths := make([][]Vec2, max_depth+1)
	for i := 1; i <= max_depth; i++ {
		depths[i] = make([]Vec2, len(pts)*int(math.Pow(float64(len(mats)), float64(i))))
	}
	depths[0] = pts

	var wg sync.WaitGroup
	wg.Add(1)

	var worker func(startInclusive, endExclusive, d int)
	worker = func(startInclusive, endExclusive, d int) {
		for midx, m := range mats {
			for idx := startInclusive; idx < endExclusive; idx++ {
				i := idx % n_pts
				b := idx / n_pts
				newidx := b*n_mats*n_pts + midx*n_pts + i
				depths[d+1][newidx] = m.TransformPoint(&depths[d][idx])
			}
		}
		if d < max_depth-1 {
			b := startInclusive / n_pts
			wg.Add(n_mats)
			for midx, _ := range mats {
				newidx := b*n_mats*n_pts + midx*n_pts
				go worker(newidx, newidx+n_pts, d+1)
			}
		}
		wg.Done()
	}
	go worker(0, len(pts), 0)

	wg.Wait()

	return depths[len(depths)-1]
}

func RenderFractal(mats []Matrix3, filename string, depth int) {
	RenderFractal0(mats, filename, depth, 800, DefaultFractalBounds)
}
func RenderFractal0(mats []Matrix3, filename string, depth int, width int, bounds [4]Float) {
	pts := TransformPoints([]Vec2{Vec2Zero}, mats, depth)
	img := RasterizePoints0(width, pts, bounds)
	file, err := os.Create(filename)
	Die(err)
	Die(png.Encode(file, img))
	Die(file.Close())
}

var DefaultFractalBounds = [4]Float{-1.0, 1.0, -1.0, 1.0}
