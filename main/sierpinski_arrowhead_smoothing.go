package main

import (
	m ".."
	"os"
	"fmt"
	"sync"
	"image"
	"image/gif"
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
	frames := make([]*image.Paletted, n_frames)

	var wg sync.WaitGroup
	for i := 0; i < n_frames; i++ {
		wg.Add(1)
		go func(smoothness int) {
			fmt.Println("smoothness:", smoothness)
			pts2 := m.BSplineAdaptive(pts, smoothness, 2/m.Float(width))

			img := m.RasterizePointsPalletted(width, pts2, bounds)
			frames[smoothness] = img

			//img := m.RasterizePoints0(width, pts2, bounds)
			//file, err := os.Create(fmt.Sprintf("%s/%s_%d.png", folder_path, folder_path, smoothness))
			//m.Die(err)
			//m.Die(png.Encode(file, img))
			//m.Die(file.Close())

			wg.Done()
		}(i)
	}
	wg.Wait()

	outGif := &gif.GIF{}
	outGif.Image = frames
	outGif.Delay = make([]int, n_frames)
	for i := 0; i < n_frames; i++ {
		outGif.Delay[i] = 0
	}

	file, err := os.OpenFile(m.ExecutableNameWithExtension("gif"), os.O_WRONLY|os.O_CREATE, 0600)
	m.Die(err)
	defer file.Close()
	gif.EncodeAll(file, outGif)

}
