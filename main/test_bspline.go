package main

import (
	m "github.com/joshua-wright/image_stuff/go"
	"os"
	"image/png"
)

func main() {
	pts := []m.Vec2{
		{-1,-1},
		{-1,1},
		{1,1},
		{1,-1},
		{-1,-1},
	}

	pts = m.BSpline(pts, 9, 1)

	img := m.RasterizePoints0(800, pts, m.DefaultFractalBounds)
	file, err := os.Create(m.ExecutableNamePng())
	m.Die(err)
	m.Die(png.Encode(file, img))
	m.Die(file.Close())
}
