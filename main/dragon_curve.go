package main

import (
	m ".."
)

func main() {
	// these don't work, not sure why

	//t1 := m.Scale2D(-1, 0, math.Sqrt2).Compose(m.Rotate2D(-1, 0, -math.Pi/4.0))
	//t2 := t1.Compose(m.Rotate2D(0,-1,math.Pi/2))

	//t1 := m.MatrixImage2D(
	//	m.Vec2{-1, 0}, m.Vec2{0, -1}, m.Vec2{1, 0},
	//	m.Vec2{-1, 0}, m.Vec2{-1, 1}, m.Vec2{0, -1},
	//)
	//t2 := m.MatrixImage2D(
	//	m.Vec2{-1, 0}, m.Vec2{0, -1}, m.Vec2{1, 0},
	//	m.Vec2{0, -1}, m.Vec2{0, 0}, m.Vec2{1, 0},
	//)

	//r := 1.0 / math.Sqrt2
	//theta1 := math.Pi / 4
	//theta2 := 5 * math.Pi / 4
	//t1 := m.Matrix3{
	//	{r * math.Cos(theta1), -r * math.Sin(theta1), 0},
	//	{r * math.Sin(theta1), r * math.Cos(theta1), 0},
	//	{0, 0, 1},
	//}
	//t2 := m.Matrix3{
	//	{r * math.Cos(theta2), -r * math.Sin(theta2), 1},
	//	{r * math.Sin(theta2), r * math.Cos(theta2), 0},
	//	{0, 0, 1},
	//}
	//
	//mats := []m.Matrix3{t1, t2}

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

	width := 1000
	xmid := 0.0
	ymid := 0.0
	dw := 1.5

	//m.RenderFractal(mats, "dragon_curve.png", 15)
	m.RenderFractal0(mats, m.ExecutableNamePng(), 17, width, [4]m.Float{xmid - dw, xmid + dw, ymid - dw, ymid + dw})
}
