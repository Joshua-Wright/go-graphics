package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/parallel"
	"github.com/joshua-wright/go-graphics/Slice2D"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"gopkg.in/cheggaaa/pb.v1"
	"image"
	"math"
	"runtime"
	"math/rand"
)

func chua_circuit(a, b, c, d float64, v g.Vec3) g.Vec3 {
	fx := c*v.X + 0.5*(d-c)*(math.Abs(v.X+1)-math.Abs(v.X-1))
	return g.Vec3{
		X: a * (v.Y - v.X - fx),
		Y: v.X - v.Y + v.Z,
		Z: -b * v.Y,
	}
}

func Uint16AddSaturate(a, b uint16) uint16 {
	if a < (math.MaxUint16 - b) {
		return a + b
	} else {
		return math.MaxUint16
	}
}

func main() {
	//height := 1440*2
	//width := int(height*16.0/9.0) + 2*int(height*4.0/5.0)
	width := 1920
	height := 1080
	darkPerPoint := uint16(8)
	nPts := 32
	nPreIter := 200
	nIter := 5000000
	deltaTime := 0.0001

	scale := g.Vec2{
		X: 3.0,
		Y: 1.5,
	}

	a := 15.446
	b := 28.0
	c := -0.714
	d := -1.143
	//a := 15.6
	//b := 28.0
	//c := -0.714
	//d := -1.143

	nProcs := runtime.GOMAXPROCS(-1)
	nPtsPerProc := nPts / nProcs
	buffers := make([]Slice2D.Uint16Slice2D, nProcs)
	for i := 0; i < len(buffers); i++ {
		buffers[i] = Slice2D.NewUint16Slice2D(width, height)
	}

	bound_width := 10.0
	bounds := [4]g.Float{
		// a little extra space around the rose
		-bound_width, bound_width,
		-bound_width, bound_width,
		//-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	println("iterate points")
	iterateBar := pb.StartNew(nPts)
	parallel.ParallelFor(0, nProcs, func(workerIndex int) {
		counts := buffers[workerIndex]

		addPixel := func(i, j int, ammount uint16) {
			if i >= width || j >= height || i < 0 || j < 0 {
				return
			}
			v := Uint16AddSaturate(ammount, counts.Get(i, j))
			counts.Set(i, j, v)
		}

		for i := 0; i < nPtsPerProc; i++ {
			pt := g.Vec3{
				//X: 0.7,
				//Y: 0,
				//Z: 0,
				X: rand.Float64(),
				Y: rand.Float64(),
				Z: rand.Float64(),
			}

			//println(pt.String())
			for i := 0; i < nPreIter; i++ {
				dt := chua_circuit(a, b, c, d, pt)
				pt = pt.AddV(dt.MulS(deltaTime))
			}
			//println(pt.String())

			// iterate the point the rest of the time
			for j := 0; j < nIter; j++ {
				dt := chua_circuit(a, b, c, d, pt)
				pt = pt.AddV(dt.MulS(deltaTime))
				//pt2 := g.Vec2{pt.X, pt.Y}
				pt2 := g.Vec2{pt.X, pt.Z}.MulV(scale)
				x, y := g.WindowTransformPoint(width, height, pt2, bounds)

				addPixel(x, height-1-y, darkPerPoint)
				addPixel(x, height-1-y+1, darkPerPoint/2)
				addPixel(x, height-1-y-1, darkPerPoint/2)
				addPixel(x+1, height-1-y, darkPerPoint/2)
				addPixel(x-1, height-1-y, darkPerPoint/2)

			}
			iterateBar.Increment()
		}
	})
	iterateBar.Finish()

	println("reduce")
	reduceBar := pb.StartNew(width)
	//img := image.NewGray(image.Rect(0, 0, width, height))
	img := image.NewPaletted(image.Rect(0, 0, width, height), colormap.InfernoColorMap())
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
