package naive_fixnum

//go:generate bash generate_fixnums.sh

import (
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_2"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_3"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_4"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_5"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_6"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_7"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_8"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_10"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_12"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_14"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_16"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_18"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_22"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_26"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_30"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"github.com/joshua-wright/go-graphics/graphics/per_pixel_image"
)

type MandelbrotPerPixelConstructor func(width, height, maxIter int64,
	centerR, centerI, zoom, threshold string,
	Wrap, MaxVal float64, cmap colormap.ColorMap,
	OutImage *memory_mapped.PPMFile,
	OutIter *memory_mapped.Array2dFloat64) per_pixel_image.PixelFunction

var FixnumConstructorWords = []uint{
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	10,
	12,
	14,
	16,
	18,
	22,
	26,
	30,
}

var FixnumConstructors = []MandelbrotPerPixelConstructor{
	naive_fixnum_2.NewMandelbrotPerPixel,
	naive_fixnum_3.NewMandelbrotPerPixel,
	naive_fixnum_4.NewMandelbrotPerPixel,
	naive_fixnum_5.NewMandelbrotPerPixel,
	naive_fixnum_6.NewMandelbrotPerPixel,
	naive_fixnum_7.NewMandelbrotPerPixel,
	naive_fixnum_8.NewMandelbrotPerPixel,
	naive_fixnum_10.NewMandelbrotPerPixel,
	naive_fixnum_12.NewMandelbrotPerPixel,
	naive_fixnum_14.NewMandelbrotPerPixel,
	naive_fixnum_16.NewMandelbrotPerPixel,
	naive_fixnum_18.NewMandelbrotPerPixel,
	naive_fixnum_22.NewMandelbrotPerPixel,
	naive_fixnum_26.NewMandelbrotPerPixel,
	naive_fixnum_30.NewMandelbrotPerPixel,
}
