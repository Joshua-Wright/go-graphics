package per_pixel_image

import "image"

type PixelWorkFunc interface {
	GetPixel(i, j int) image.RGBA
}

