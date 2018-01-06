package graphics

import (
	"os"
	"image/png"
	"math"
	"sync"
)

func TransformPointsSerial(pts []Vec2, mats []Matrix3, max_depth int) []Vec2 { return TransformPointsImpl(pts, mats, max_depth, false) }
func TransformPoints(pts []Vec2, mats []Matrix3, max_depth int) []Vec2       { return TransformPointsImpl(pts, mats, max_depth, true) }

func TransformPointsImpl(pts []Vec2, mats []Matrix3, max_depth int, parallel bool) []Vec2 {
	// make sure that we don't do parallel computation unless the data size
	// is larger than some threshold
	threshold := 256

	// FIXME: this is not particularly efficient
	if max_depth == 0 {
		// do not return input array
		tmp := make([]Vec2, len(pts))
		copy(tmp, pts)
		return tmp
	}

	for {
		if max_depth == 0 {
			// do not return input array
			tmp := make([]Vec2, len(pts))
			copy(tmp, pts)
			return tmp
		}

		newpts := make([]Vec2, 0, len(mats)*len(pts))

		for _, m := range mats {
			for _, p := range pts {
				newpts = append(newpts, m.TransformPoint(&p))
			}
		}
		if max_depth == 0 {
			return newpts
		} else if len(newpts) > threshold && parallel {
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
		if d == max_depth {
			wg.Done()
			return
		}
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
	RenderFractal0(mats, filename, depth, 800, 800, DefaultFractalBounds)
}
func RenderFractal0(mats []Matrix3, filename string, depth int, width, height int, bounds [4]Float) {
	pts := TransformPoints([]Vec2{Vec2Zero}, mats, depth)
	img := RasterizePoints0(width, height, pts, bounds)
	file, err := os.Create(filename)
	Die(err)
	Die(png.Encode(file, img))
	Die(file.Close())
}

var DefaultFractalBounds = [4]Float{-1.0, 1.0, -1.0, 1.0}

func BoundsForResolution(xmid, ymid, xwidth Float, w_res, h_res int) [4]Float {
	// half width to get distance from center to edge
	xwidth = xwidth / 2
	aspect := Float(w_res) / Float(h_res)
	ywidth := Float(xwidth) / aspect
	bounds := [4]Float{xmid - xwidth, xmid + xwidth, ymid - ywidth, ymid + ywidth}
	return bounds
}
