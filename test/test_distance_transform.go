package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"image"
	"github.com/joshua-wright/go-graphics/graphics/distance_transform"
	"image/color"
	"math"
	"runtime"
)

func main() {
	width := 800
	height := 800

	zero_points := []g.Vec2i{
		{0, 0},
		{width / 2, height / 2},
		{0, height / 2},
	}

	mesh := distance_transform.DistanceTransform(width, height, zero_points)
	maxdist := 0.0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if mesh[x][y] > maxdist {
				maxdist = mesh[x][y]
			}
		}
	}
	//img := image.NewNRGBA(image.Rect(0, 0, width, height))
	img := image.NewGray16(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			//fmt.Println(x, y, mesh[x][y])
			c := color.Gray16{uint16(mesh[x][y] / maxdist * float64(math.MaxInt16))}
			img.SetGray16(x, y, c)
		}
	}

	g.SaveAsPNG(img, g.ExecutableNamePng())
}
