package main

import (
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot"
)

func main() {
	factor := int64(1)
	width := int64(2560)
	height := int64(1440)
	//maxIter := int64(25600)
	//factor := int64(2)
	//width := int64(192)
	//height := int64(108)
	//maxIter := int64(2560)

	topLeft, dr, di := mandelbrot.MandelbrotBounds(width*factor, height*factor, complex(-0.74364085, 0.13182733), 25497*1.1)

	println(topLeft, dr, di)
}
