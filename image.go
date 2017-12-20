package golang_raytracer

import (
	"image"
	"image/color"
)

//func ReRangePoints(pts []Vec2, wd int, bounds [4]Float) []Vec2 {
//	outpts := make([]Vec2, 0, len(pts))
//	xmin := bounds[0]
//	xmax := bounds[1]
//	ymin := bounds[2]
//	ymax := bounds[3]
//	wdF := Float(wd)
//	for _, p := range pts {
//		outpts = append(outpts,
//			Vec2{
//				(p.X - xmin) / (xmax - xmin) * (wdF - 1),
//				wdF - 1 - (p.Y-ymin)/(ymax-ymin)*(wdF-1),
//			})
//	}
//	return outpts
//}

func RasterizePoints1(width int, pts []Vec2) image.Image {
	return RasterizePoints0(width, pts, DefaultFractalBounds)
}

func TransformPoint(width int, p Vec2, bounds [4]Float) (int, int) {
	xmin := bounds[0]
	xmax := bounds[1]
	ymin := bounds[2]
	ymax := bounds[3]
	x := p.X
	y := p.Y
	xi := (x - xmin) / (xmax - xmin) * Float(width-1)
	yi := Float(width-1) - (y-ymin)/(ymax-ymin)*Float(width-1)
	return int(xi), int(yi)
}

func RasterizePoints0(width int, pts []Vec2, bounds [4]Float) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, width))

	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{0, 0, 0, 255})
		}
	}

	for _, p := range pts {
		xi, yi := TransformPoint(width, p, bounds)
		img.Set(xi, yi, color.NRGBA{255, 255, 255, 255})
	}

	return img
}
