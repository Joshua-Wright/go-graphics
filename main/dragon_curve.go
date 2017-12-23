package main

import (
	m ".."
	"github.com/fogleman/gg"
	"image/color"
)

func main() {
	width := 2000
	xmid := -0.25
	ymid := -0.25
	dw := 1.75
	depth := 11
	smoothness := 1
	line_thickness := 2.0

	bounds := [4]m.Float{xmid - dw, xmid + dw, ymid - dw, ymid + dw}
	filename := m.ExecutableNamePng()

	a := m.Vec2{-1, 0}
	b := m.Vec2{0, -1}
	c := m.Vec2{1, 0}
	d := m.Vec2{-1, -1}
	e := m.Vec2{0, 0}
	mats := []m.Matrix3{
		m.MatrixImage2D(
			a, b, c,
			a, d, b,
		),
		m.MatrixImage2D(
			a, b, c,
			c, e, b,
		),
	}
	t1 := mats[0]
	t2 := mats[1]

	// specialize for dragon curve: reverse points from second transform
	TransformPoints := func(pts []m.Vec2, max_depth int) []m.Vec2 {
		if max_depth == 0 {
			return pts
		}

		for {
			if max_depth == 0 {
				return pts
			}

			newpts := make([]m.Vec2, 0, len(mats)*len(pts))

			for _, p := range pts {
				newpts = append(newpts, t1.TransformPoint(&p))
			}
			for i := len(pts) - 1; i >= 0; i-- {
				p := pts[i]
				newpts = append(newpts, t2.TransformPoint(&p))
			}

			if max_depth == 0 {
				return newpts
			} else {
				max_depth -= 1
				pts = newpts
			}
		}
	}

	ctx := gg.NewContext(width, width)
	// set black background
	ctx.SetColor(color.NRGBA{0, 0, 0, 255})
	ctx.DrawRectangle(0, 0, float64(width), float64(width))
	ctx.Fill()

	pts := TransformPoints([]m.Vec2{m.Vec2Zero}, depth)
	for i, _ := range pts {
		x, y := m.WindowTransformPoint(width, pts[i], bounds)
		pts[i] = m.Vec2{m.Float(x), m.Float(y)}
	}
	pts = m.BSplineAdaptive(pts, smoothness, 1.0)

	ctx.SetColor(color.NRGBA{255, 255, 255, 255})
	ctx.SetLineWidth(line_thickness)
	for i, p := range pts {
		if i+1 == len(pts) {
			break
		}
		x1 := p.X
		y1 := p.Y
		x2 := pts[i+1].X
		y2 := pts[i+1].Y
		ctx.DrawLine(x1, y1, x2, y2)
	}
	ctx.Stroke()

	m.Die(ctx.SavePNG(filename))
}
