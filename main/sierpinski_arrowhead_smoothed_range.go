package main

import (
	m "github.com/joshua-wright/image_stuff/go"
	"os"
	"fmt"
	"sync"
	"image"
	"image/color"
	"image/png"
)

func main() {
	//filename := m.ExecutableNamePng()
	width := 1000
	xmid := 0.0
	ymid := 0.0
	dw := 1.125
	depth := 5
	bounds := [4]m.Float{xmid - dw, xmid + dw, ymid - dw, ymid + dw}
	folder_path := m.ExecutableName()

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

	os.Mkdir(folder_path, 0777)

	n_frames := 150
	pointsAtSmoothness := make([][]m.Vec2, n_frames)

	var wg sync.WaitGroup
	for i := 0; i < n_frames; i++ {
		wg.Add(1)
		go func(smoothness int) {
			fmt.Println("smoothness:", smoothness)
			pts2 := m.BSplineAdaptive(pts, smoothness, 2/m.Float(width))

			for i := 0; i < len(pts2); i++ {
				x, y := m.TransformPoint(width, pts2[i], bounds)
				pts2[i] = m.Vec2{m.Float(x), m.Float(y)}
			}

			pointsAtSmoothness[smoothness] = pts2

			wg.Done()
		}(i)
	}
	wg.Wait()

	img := image.NewNRGBA(image.Rect(0, 0, width, width))
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{0, 0, 0, 255})
		}
	}

	for i := n_frames - 1; i >= 0; i-- {
		gray := uint8(255 - 255*m.Float(i)/m.Float(n_frames))
		smooth_pts := pointsAtSmoothness[i]
		for _, p := range smooth_pts {
			xi := int(p.X)
			yi := int(p.Y)
			img.Set(xi, yi, color.NRGBA{gray, gray, gray, 255})
		}
	}

	file, err := os.Create(m.ExecutableNamePng())
	m.Die(err)
	m.Die(png.Encode(file, img))
	m.Die(file.Close())
}
