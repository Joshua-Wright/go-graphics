package naive_fixnum

import (
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
	"strconv"
	"math"
	"fmt"
)

type MandelbrotPerPixelConstructor func(width, height, maxIter int64,
	centerR, centerI, zoom, threshold string,
	Wrap, MaxVal float64, cmap colormap.ColorMap,
	OutImage *memory_mapped.PPMFile,
	OutIter *memory_mapped.Array2dFloat64) per_pixel_image.PixelFunction

func NewMandelbrotPerPixel(width, height, maxIter int64,
	centerR, centerI, zoom, threshold string,
	Wrap, MaxVal float64, cmap colormap.ColorMap,
	OutImage *memory_mapped.PPMFile,
	OutIter *memory_mapped.Array2dFloat64) per_pixel_image.PixelFunction {
	zoomf, err := strconv.ParseFloat(zoom, 64)
	if err != nil {
		panic(err)
	}
	bits := uint64(math.Log2(zoomf*math.Max(float64(width), float64(height)))) + 8
	words := bits/32 + 1
	c, ok := FixnumConstructorMap[words]
	if ok {
		fmt.Println("usng", words, "words (needed", bits, "bits)")
		return c(width, height, maxIter, centerR, centerI, zoom, threshold, Wrap, MaxVal, cmap, OutImage, OutIter)
	} else {
		fmt.Println("failed to find", bits, "bits, or", words, "words!")
		return nil
	}
}
