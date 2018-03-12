package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/parallel"
	"os"
	"github.com/joshua-wright/go-graphics/Slice2D"
	"gopkg.in/src-d/go-git.v4/utils/binary"
	"gopkg.in/cheggaaa/pb.v1"
	"image"
	"image/draw"
)

func main() {
	width := 5120
	height := int(float64(width) * 9.0 / 16.0)
	darkPerPoint := uint8(8)
	nPtsPerX := 8000
	//nIter := 5000

	img := image.NewGray(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Over)

	println("allocate")
	valuesR := Slice2D.NewFloat64Slice2D(width, nPtsPerX)
	valuesX := Slice2D.NewFloat64Slice2D(width, nPtsPerX)

	println("open files")
	fileX, err := os.Open("logistic_map_points._x.dat")
	g.Die(err)
	defer fileX.Close()
	fileY, err := os.Open("logistic_map_points._y.dat")
	g.Die(err)
	defer fileY.Close()

	println("read files")
	binary.Read(fileX, valuesR.Data)
	binary.Read(fileY, valuesX.Data)

	print("render")
	bar := pb.StartNew(width)
	parallel.ParallelFor(0, width, func(i int) {
		for j := 0; j < nPtsPerX; j++ {
			//r := valuesR.Get(i, j)
			x := valuesX.Get(i, j)
			if x < 0 || x > 1 {
				panic("bad x")
			}
			yval := int(float64(x) * float64(height))
			c := img.GrayAt(i, height-1-yval)
			if c.Y > darkPerPoint {
				c.Y -= darkPerPoint
			}
			img.SetGray(i, height-1-yval, c)
		}
		bar.Increment()
	})
	bar.Finish()

	println("write output")
	g.SaveAsPNG(img, g.ExecutableNamePng())

	//out := image.NewRGBA(img.Bounds())
	//draw.Draw(out, out.Bounds(), image.White, image.ZP, draw.Over)
	//draw.Draw(out, out.Bounds(), img, image.ZP, draw.Over)
	//g.SaveAsPNG(out, g.ExecutableNamePng())

}
