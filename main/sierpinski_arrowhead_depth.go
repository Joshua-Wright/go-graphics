package main

import (
	m ".."
	"fmt"
	"sync"
	"image"
	"image/color"
	"github.com/fogleman/gg"
	"math"
)

var colors = []color.NRGBA{
	{63, 81, 181, 255},
	{0, 150, 136, 255},
	{0, 188, 212, 255},
	{103, 58, 183, 255},
	{139, 195, 74, 255},
	{156, 39, 176, 255},
	{205, 220, 57, 255},
	{233, 30, 99, 255},
	{244, 67, 54, 255},
	{255, 152, 0, 255},
	{255, 193, 7, 255},
	{255, 235, 59, 255},
	{255, 87, 34, 255},
	{3, 169, 244, 255},
	{33, 150, 243, 255},
	{76, 175, 80, 255},
}

func main() {
	width := 2000
	thickest_line := 20.0
	line_step := 0.7
	xmid := 0.0
	ymid := 0.0
	dw := 1.125
	n_depths := 7
	bounds := [4]m.Float{xmid - dw, xmid + dw, ymid - dw, ymid + dw}

	a := m.Vec2{-1, -1}
	b := m.Vec2{0, 1}
	c := m.Vec2{1, -1}
	d := m.Vec2{-0.5, 0}
	e := m.Vec2{0.5, 0}
	f := m.Vec2{0, -1}

	start_points := []m.Vec2{a, d, e, c}

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

	pointsAtDepth := make([][]m.Vec2, n_depths)

	var wg sync.WaitGroup
	for i := 0; i < n_depths; i++ {
		wg.Add(1)
		//go
		func(depth int) {
			fmt.Println("depth:", depth)
			pts := m.TransformPointsSerial(start_points, mats, depth)

			for i := 0; i < len(pts); i++ {
				x, y := m.TransformPoint(width, pts[i], bounds)
				pts[i] = m.Vec2{m.Float(x), m.Float(y)}
			}

			pointsAtDepth[depth] = pts

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

	ctx := gg.NewContext(width, width)
	// set black background
	ctx.SetColor(color.NRGBA{0, 0, 0, 255})
	ctx.DrawRectangle(0, 0, float64(width), float64(width))
	ctx.Fill()

	for i := n_depths - 1; i >= 0; i-- {
		smooth_pts := m.BSplineAdaptive(pointsAtDepth[i], 0, 1)
		c := colors[i%len(colors)]

		ctx.SetColor(c)
		ctx.SetLineWidth(thickest_line * math.Pow(line_step, float64(i)))
		for i, p := range smooth_pts {
			if (i+1 == len(smooth_pts)) {
				break
			}
			x1 := p.X
			y1 := p.Y
			x2 := smooth_pts[i+1].X
			y2 := smooth_pts[i+1].Y
			ctx.DrawLine(x1, y1, x2, y2)
		}
		ctx.Stroke()
	}

	m.Die(ctx.SavePNG(m.ExecutableNamePng()))
}
