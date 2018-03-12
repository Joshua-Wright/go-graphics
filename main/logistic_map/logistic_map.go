package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/parallel"
	"github.com/joshua-wright/go-graphics/Slice2D"
	"gopkg.in/cheggaaa/pb.v1"
	"image"
	"image/draw"
	"math"
	"math/rand"
	"runtime"
	"image/color"
)

func logistic_map(r, x float64) float64 {
	return r * x * (1.0 - x)
}

func logistic_map_iter(r, x_init float64, nIter int) (x float64) {
	x = x_init
	for i := 0; i < nIter; i++ {
		x = logistic_map(r, x)
	}
	return x
}

func Uint16AddSaturate(a, b uint16) uint16 {
	if a < (math.MaxUint16 - b) {
		return a + b
	} else {
		return math.MaxUint16
	}
}

func main() {
	//height := 1440
	//width := int(height*16.0/9.0) + 2*int(height*4.0/5.0)
	width := 1920
	height := 1080
	darkPerPoint := uint16(64)
	nPtsPerX := 64000
	nIter := 5000
	exponent := 1 / 5.0

	//rMin := 3.0
	//rMin := 3.54409
	// do the full range, and then squish it sideways later
	rMin := 1.0
	rMax := 4.0

	nProcs := runtime.GOMAXPROCS(-1)
	nPtsPerXPerProc := nPtsPerX / nProcs
	buffers := make([]Slice2D.Uint16Slice2D, nProcs)
	for i := 0; i < len(buffers); i++ {
		buffers[i] = Slice2D.NewUint16Slice2D(width, height)
	}

	println("iterate points")
	iterateBar := pb.StartNew(nProcs*width)
	parallel.ParallelFor(0, nProcs, func(workerIndex int) {
		counts := buffers[workerIndex]
		addPixel := func(i, j int, ammount uint16) {
			if i >= width || j >= height || i < 0 || j < 0 {
				return
			}
			v := Uint16AddSaturate(ammount, counts.Get(i, j))
			counts.Set(i, j, v)
		}

		for i := 0; i < width; i++ {

			t := float64(i)/float64(width) + (rand.Float64()-0.5)/float64(width)
			r := g.Lerp(rMin, rMax, math.Pow(t, exponent))
			for j := 0; j < nPtsPerXPerProc; j++ {
				// locate point
				x_init := float64(j)/float64(nPtsPerXPerProc) + (rand.Float64()-0.5)/float64(nPtsPerXPerProc)
				// iterate the point
				x_final := logistic_map_iter(r, x_init, nIter)

				yval := int(float64(x_final) * float64(height))
				addPixel(i, height-1-yval, darkPerPoint)
				addPixel(i, height-1-yval+1, darkPerPoint/2)
				addPixel(i, height-1-yval-1, darkPerPoint/2)
				addPixel(i+1, height-1-yval, darkPerPoint/2)
				addPixel(i-1, height-1-yval, darkPerPoint/2)
			}
			iterateBar.Increment()
		}
	})
	iterateBar.Finish()

	println("reduce")
	reduceBar := pb.StartNew(width)
	img := image.NewGray(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Over)
	parallel.ParallelFor(0, width, func(i int) {
		for j := 0; j < height; j++ {
			val := uint16(0)
			for _, buf := range buffers {
				val = Uint16AddSaturate(val, buf.Get(i, j))
			}
			img.SetGray(i, j, color.Gray{255 - uint8(val/256)})
		}
		reduceBar.Increment()
	})
	reduceBar.Finish()

	println("write output")
	g.SaveAsPNG(img, g.ExecutableNamePng())

}
