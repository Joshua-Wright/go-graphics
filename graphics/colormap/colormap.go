package colormap

import "image/color"

type ColorMap interface {
	GetColor(x float64) color.RGBA
}
