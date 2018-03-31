package main

import (
	m "github.com/joshua-wright/go-graphics/graphics"
)

func main() {
	sierpinski := []m.Matrix3{
		m.Scale2D(-1, -1, 0.5),
		m.Scale2D(0, 1, 0.5),
		//m.Scale2D(1, -1, 0.5),
		m.MatrixImage2D(
			m.Vec2{-1, -1}, m.Vec2{0, 1}, m.Vec2{1, -1},
			m.Vec2{0, -1}, m.Vec2{0.5, 0}, m.Vec2{1, -1},
		),
	}
	m.RenderFractal(sierpinski, "sierpinski.png", 15)
}
