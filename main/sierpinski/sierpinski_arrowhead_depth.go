package main

import (
	m "github.com/joshua-wright/go-graphics/graphics"
	"fmt"
	"sync"
	"image"
	"image/color"
	"github.com/fogleman/gg"
	"math"
)

var colors = []string{
	"#13a2dd",
	"#2a8c2a",
	"#c73232",
	"#80267f",
	"#2e8c74",
	"#66361f",
	"#d9a641",
	"#c73232",
	"#195555",
	"#3dcc3d",
	"#4545e6",
	"#b037b0",
}

func main() {
	width := 2000
	height := 2000
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
		go func(depth int) {
			fmt.Println("depth:", depth)
			pts := m.TransformPointsSerial(start_points, mats, depth)

			for i := 0; i < len(pts); i++ {
				x, y := m.WindowTransformPoint(width, height, pts[i], bounds)
				pts[i] = m.Vec2{m.Float(x), m.Float(y)}
			}

			pointsAtDepth[depth] = pts

			wg.Done()
		}(i)
	}

	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{0, 0, 0, 255})
		}
	}

	ctx := gg.NewContext(width, height)
	// set black background
	ctx.SetColor(color.NRGBA{0, 0, 0, 255})
	ctx.DrawRectangle(0, 0, float64(width), float64(height))
	ctx.Fill()

	// don't wait for points until we absolutely need to
	wg.Wait()

	for i := n_depths - 1; i >= 0; i-- {
		smooth_pts := m.BSplineAdaptive(pointsAtDepth[i], 0, 1)
		c := colors[i%len(colors)]

		ctx.SetHexColor(c)
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
