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

func RasterizePoints1(width, height int, pts []Vec2) image.Image {
	return RasterizePoints0(width, height, pts, DefaultFractalBounds)
}

func WindowTransformPoint(width, height int, p Vec2, bounds [4]Float) (int, int) {
	xmin := bounds[0]
	xmax := bounds[1]
	ymin := bounds[2]
	ymax := bounds[3]
	x := p.X
	y := p.Y
	max_dim := 0
	//min_dim := 0
	if width > height {
		max_dim = width
		//min_dim = height
	} else {
		max_dim = height
		//min_dim = width
	}
	xi := (x - xmin) / (xmax - xmin) * Float(max_dim-1)
	yi := Float(height-1) - (y-ymin)/(ymax-ymin)*Float(height-1)
	return int(xi), int(yi)
}

func RasterizePoints0(width, height int, pts []Vec2, bounds [4]Float) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			img.Set(x, y, color.NRGBA{0, 0, 0, 255})
		}
	}

	for _, p := range pts {
		xi, yi := WindowTransformPoint(width, height, p, bounds)
		img.Set(xi, yi, color.NRGBA{255, 255, 255, 255})
	}

	return img
}

func RasterizePointsPalletted(width, height int, pts []Vec2, bounds [4]Float) *image.Paletted {
	img := image.NewPaletted(
		image.Rect(0, 0, width, width),
		[]color.Color{
			color.NRGBA{0, 0, 0, 255},
			color.NRGBA{255, 255, 255, 255},
		})

	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{0, 0, 0, 255})
		}
	}

	for _, p := range pts {
		xi, yi := WindowTransformPoint(width, height, p, bounds)
		img.Set(xi, yi, color.NRGBA{255, 255, 255, 255})
	}

	return img
}
