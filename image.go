package golang_raytracer

import (
	"image"
	"image/color"
)

func RasterizePoints(width int, pts []Vec2) image.Image {
	xmin := -1.0
	ymin := -1.0
	xmax := 1.0
	ymax := 1.0
	img := image.NewNRGBA(image.Rect(0, 0, width, width))

	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{0, 0, 0, 255})
		}
	}

	for _, p := range pts {
		x := p.X
		y := p.Y
		xi := (x - xmin) / (xmax - xmin) * Float(width-1)
		yi := Float(width-1) - (y-ymin)/(ymax-ymin)*Float(width-1)
		img.Set(int(xi), int(yi), color.NRGBA{255, 255, 255, 255})
	}

	return img
}
