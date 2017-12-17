package golang_raytracer

import (
	"os"
	"image/png"
)

func TransformPoints(pts []Vec2, mats []Matrix3, depth int) []Vec2 {
	for {
		newpts := make([]Vec2, 0, len(mats)*len(pts))

		for _, m := range mats {
			for _, p := range pts {
				newpts = append(newpts, m.TransformPoint(&p))
			}
		}
		if depth == 0 {
			return newpts
		} else {
			depth -= 1
			pts = newpts
		}
	}
}
func RenderFractal(mats []Matrix3, filename string, depth int) {
	RenderFractal0(mats, filename, depth, 800, DefaultFractalBounds)
}
func RenderFractal0(mats []Matrix3, filename string, depth int, width int, bounds [4]Float) {
	pts := TransformPoints([]Vec2{Vec2Zero}, mats, depth)
	img := RasterizePoints0(width, pts, bounds)
	f, err := os.Create(filename)
	Die(err)
	Die(png.Encode(f, img))
	Die(f.Close())
}

var DefaultFractalBounds = [4]Float{-1.0, 1.0, -1.0, 1.0}