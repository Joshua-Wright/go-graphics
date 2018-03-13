package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/parallel"
	"github.com/joshua-wright/go-graphics/Slice2D"
	"github.com/joshua-wright/go-graphics/graphics/colors"
	"gopkg.in/cheggaaa/pb.v1"
	"image"
	"math"
	"math/rand"
	"runtime"
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

func sigmoid_squish(x float64) float64 {
	x = x * 6
	return 2*math.Exp(x)/(math.Exp(x)+1) - 1
}

func main() {
	//height := 1440*2
	//width := int(height*16.0/9.0) + 2*int(height*4.0/5.0)
	width := 1920 * 2
	height := 1080 * 2
	darkPerPoint := uint16(8)
	nPtsPerX := 256
	nPreIter := 2000
	nIter := 5000
	exponent := 1 / 5.0

	//rMin := 3.0
	// do the full range, and then squish it sideways later
	rMin := 1.0
	// rMin := 3.54409
	rMax := 4.0

	_ = exponent

	nProcs := runtime.GOMAXPROCS(-1)
	nPtsPerXPerProc := nPtsPerX / nProcs
	buffers := make([]Slice2D.Uint16Slice2D, nProcs)
	for i := 0; i < len(buffers); i++ {
		buffers[i] = Slice2D.NewUint16Slice2D(width, height)
	}

	println("iterate points")
	iterateBar := pb.StartNew(nProcs * width)
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

			for j := 0; j < nPtsPerXPerProc; j++ {
				t := float64(i)/float64(width) + (rand.Float64()-0.5)/float64(width)
				// find r
				r := g.Lerp(rMin, rMax, math.Pow(t, exponent))
				// r := g.Lerp(rMin, rMax, sigmoid_squish(t))
				// r := g.Lerp(rMin, rMax, t)

				// locate point
				x_init := float64(j)/float64(nPtsPerXPerProc) + (rand.Float64()-0.5)/float64(nPtsPerXPerProc)

				// iterate the point
				x_final := logistic_map_iter(r, x_init, nPreIter)

				// iterate the point the rest of the time
				for k := 0; k < nIter; k++ {
					yval := int(float64(x_final) * float64(height))
					addPixel(i, height-1-yval, darkPerPoint)
					addPixel(i, height-1-yval+1, darkPerPoint/2)
					addPixel(i, height-1-yval-1, darkPerPoint/2)
					addPixel(i+1, height-1-yval, darkPerPoint/2)
					addPixel(i-1, height-1-yval, darkPerPoint/2)
					x_final = logistic_map(r, x_final)
				}
			}
			iterateBar.Increment()
		}
	})
	iterateBar.Finish()

	println("reduce")
	reduceBar := pb.StartNew(width)
	//img := image.NewGray(image.Rect(0, 0, width, height))
	img := image.NewPaletted(image.Rect(0, 0, width, height), colors.InfernoColorMap())
	//draw.Draw(img, img.Bounds(), image.Black, image.ZP, draw.Over)
	parallel.ParallelFor(0, width, func(i int) {
		for j := 0; j < height; j++ {
			val := uint16(0)
			for _, buf := range buffers {
				val = Uint16AddSaturate(val, buf.Get(i, j))
			}
			//img.SetGray(i, j, color.Gray{255 - uint8(val/256)})
			img.SetColorIndex(i, j, uint8(val/256))
		}
		reduceBar.Increment()
	})
	reduceBar.Finish()

	println("write output")
	g.SaveAsPNG(img, g.ExecutableNamePng())

}
