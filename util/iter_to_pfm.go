package main

import (
	"strings"
	"path/filepath"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/parallel"
	"math"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
)

func main() {
	factor := int64(4)

	//filename := "mandelbrot.iter"
	filename := os.Args[1]
	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".pfm"

	floats, err := memory_mapped.OpenMmap2dArrayFloat64(filename)
	g.Die(err)
	defer floats.Close()
	width := floats.Width()
	height := floats.Height()

	outpfm, err := memory_mapped.CreatePFMGray(width, height, newFilename)
	g.Die(err)
	defer outpfm.Close()

	if floats.Width()%factor != 0 {
		panic("bad width")
	}
	if floats.Height()%factor != 0 {
		panic("bad height")
	}

	MaxVal := 2560.0
	Wrap := 3.0

	bar := pb.New64(height)
	bar.Start()
	parallel.ParallelFor(0, int(height), func(y_ int) {
		y := int64(y_)
		for x := int64(0); x < width; x++ {

			// RGB with gamma correction correction

			val := floats.Get(x, y)
			if val != 0.0 {
				val = math.Log2(val+1) / math.Log2(MaxVal+1) * Wrap
				//val = math.Log(100*val + 1)
				//val = math.Sqrt(val * Wrap)
				val = math.Sin(val*2*math.Pi)/2.0 + 0.5
			}
			outpfm.SetFloat(x, y, float32(val))
		}
		bar.Increment()
	})
	bar.Finish()
}
