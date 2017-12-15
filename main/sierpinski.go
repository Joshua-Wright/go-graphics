package main

import m "github.com/joshua-wright/image_stuff/go"

func main() {
	sierpinski := []m.Matrix3{
		m.Scale2D(-1, -1, 0.5),
		m.Scale2D(0, 1, 0.5),
		m.Scale2D(1, -1, 0.5),
	}
	m.RenderFractal(sierpinski, "sierpinski.png", 12)
}