package main

import (
	m "github.com/joshua-wright/image_stuff/go"
	"os"
	"image/png"
)

func main() {
	filename := m.ExecutableNamePng()
	width := 1000
	xmid := 0.0
	ymid := 0.0
	dw := 1.125
	depth := 5
	bounds := [4]m.Float{xmid - dw, xmid + dw, ymid - dw, ymid + dw}

	a := m.Vec2{-1, -1}
	b := m.Vec2{0, 1}
	c := m.Vec2{1, -1}
	d := m.Vec2{-0.5, 0}
	e := m.Vec2{0.5, 0}
	f := m.Vec2{0, -1}

	start_points := []m.Vec2{a, d, e, c}
	//start_points := []m.Vec2{a, d, b, e, c}
	//start_points := []m.Vec2{a, b, c}

	mats := []m.Matrix3{
		m.MatrixImage2D(
			a, b, c,
			a, f, d,
		),
		m.MatrixImage2D(
			a, b, c,
			d, b, e,
		),
		m.MatrixImage2D(
			a, b, c,
			e, f, c,
		),
	}

	pts := m.TransformPoints(start_points, mats, depth)

	pts = m.BSplineAdaptive(pts, 0, 2/m.Float(width))

	img := m.RasterizePoints0(width, pts, bounds)
	file, err := os.Create(filename)
	m.Die(err)
	m.Die(png.Encode(file, img))
	m.Die(file.Close())
}
