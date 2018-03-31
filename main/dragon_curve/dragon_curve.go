package dragon_curve

import (
	m "github.com/joshua-wright/go-graphics/graphics"
	"github.com/fogleman/gg"
	"image/color"
	"math"
)

func main() {
	width := 1280
	height := 720
	//height := 2560
	xmid := -0.25
	ymid := -0.33
	dw := 1.75
	depth := 10
	smoothness := 12
	line_thickness := 4.0

	//bounds := [4]m.Float{xmid - dw, xmid + dw, ymid - dw, ymid + dw}
	bounds := m.BoundsForResolution(xmid, ymid, dw*2, width, height)
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

	ctx := gg.NewContext(width, height)
	// set black background
	ctx.SetColor(color.NRGBA{0, 0, 0, 255})
	ctx.DrawRectangle(0, 0, float64(width), float64(height))
	ctx.Fill()

	dragon := func(transform m.Matrix3, color color.Color) {

		//pts := TransformPoints([]m.Vec2{m.Vec2Zero}, depth)
		pts := TransformPoints([]m.Vec2{a, b, c}, depth)
		for i, _ := range pts {
			p := transform.TransformPoint(&pts[i])
			x, y := m.WindowTransformPoint(width, height, p, bounds)
			pts[i] = m.Vec2{m.Float(x), m.Float(y)}
		}
		pts = m.BSplineAdaptive(pts, smoothness, 1.0)
		//pts = m.BSplineAdaptive(pts, 0, 10.0)
		//for i := 0; i < smoothness*5; i++ {
		//	pts = m.BSplineAverage(pts)
		//}

		//ctx.SetColor(color.NRGBA{255, 255, 255, 255})
		ctx.SetColor(color)
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
	}

	// do 4 rotations with different colors
	dragon(m.Matrix3Identity, color.NRGBA{255, 35, 24, 255})
	dragon(m.Rotate2D(-1, 0, math.Pi/2), color.NRGBA{255, 133, 24, 255})
	dragon(m.Rotate2D(-1, 0, math.Pi), color.NRGBA{23, 206, 224, 255})
	dragon(m.Rotate2D(-1, 0, math.Pi*3.0/2.0), color.NRGBA{22, 237, 48, 255})

	m.Die(ctx.SavePNG(filename))
}
