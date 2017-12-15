package golang_raytracer

import (
	"os"
	"image/png"
)

func TransformPoints(pts []Vec2, mats []Matrix3, depth int) []Vec2 {
	for {
		var newpts []Vec2

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
	pts := TransformPoints([]Vec2{Vec2Zero}, mats, depth)
	img := RasterizePoints(800, pts)
	f, err := os.Create(filename)
	die(err)
	die(png.Encode(f, img))
	die(f.Close())
}
