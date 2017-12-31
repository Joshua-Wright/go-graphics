package main

import (
	m ".."
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"image"
	"os"
	"image/png"
	"sync"
	"fmt"
)

func main() {

	// 16:9 monitor in center flanked by two 5:4 monitors (vertical)
	// equivalent to a 16:9 monitor joined horizontally with a 8:5 monitor
	// need to make 9 equal to 5, so multiply by 5/9
	// so final ratio is 5:(8*16*5/9) = 5:16.88888888
	// for 1440px height, width should be 16.9*1440=24320.0000000000000000016
	width := 24320 / 4
	height := 24320 / 4
	// also, center about (-1,0) to get the joining point in the middle of the middle moitor
	xmid := 0.0
	ymid := 0.0
	dw := 4.0
	depth := 10
	smoothness := 12
	line_thickness := 3.0
	neon_glow_radius := 16

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

	//pts := TransformPoints([]m.Vec2{m.Vec2Zero}, depth)
	pts := TransformPoints([]m.Vec2{a, b, c}, depth)
	println("transformed")
	pts = m.BSplineAdaptive(pts, smoothness, 1.0/400.0)
	println("bsplined")

	var wg sync.WaitGroup

	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{0, 0, 0, 255})
		}
	}

	dragon := func(transform m.Matrix3, linecolor color.NRGBA) {
		ctx := gg.NewContext(width, height)
		// set black background
		ctx.SetColor(color.NRGBA{0, 0, 0, 255})
		ctx.DrawRectangle(0, 0, float64(width), float64(height))
		ctx.Fill()

		pts2 := make([]m.Vec2, len(pts))
		for i := 0; i < len(pts); i++ {
			p := transform.TransformPoint(&pts[i])
			x, y := m.WindowTransformPoint(width, height, p, bounds)
			pts2[i] = m.Vec2{m.Float(x), m.Float(y)}
		}
		//ctx.SetColor(linecolor)

		// neon glow effect
		for i := 0; i <= neon_glow_radius; i++ {
			if i != neon_glow_radius {
				c2 := linecolor
				// limit max alpha for glow
				c2.A = uint8(float64(i) / float64(neon_glow_radius) * 32.0)
				lw := line_thickness + float64(neon_glow_radius-i)
				ctx.SetColor(c2)
				ctx.SetLineWidth(lw)
				fmt.Println("i:", i, "linewidth:", lw, "alpha:", c2.A)
			} else {
				ctx.SetColor(linecolor)
				ctx.SetLineWidth(line_thickness)
			}
			for i, p := range pts2 {
				if i+1 == len(pts2) {
					break
				}
				x1 := p.X
				y1 := p.Y
				x2 := pts2[i+1].X
				y2 := pts2[i+1].Y
				ctx.DrawLine(x1, y1, x2, y2)
			}
			ctx.Stroke()
		}
		println("drew lines")

		im := ctx.Image()
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				r, g, b, _ := im.At(x, y).RGBA()
				if r != 0 || g != 0 || b != 0 {
					img.Set(x, y, im.At(x, y))
				}
			}
		}
		println("merged")
		wg.Done()
	}

	// do 4 rotations with different colors
	red := color.NRGBA{255, 35, 24, 255}
	orange := color.NRGBA{255, 133, 24, 255}
	green := color.NRGBA{22, 237, 48, 255}
	blue := color.NRGBA{23, 206, 224, 255}
	go dragon(m.Matrix3Identity, red)
	go dragon(m.Rotate2D(-1, 0, math.Pi/2), orange)
	go dragon(m.Rotate2D(-1, 0, math.Pi), green)
	go dragon(m.Rotate2D(-1, 0, math.Pi*3.0/2.0), blue)
	wg.Add(4)

	go dragon(m.Rotate2D(1, 0, math.Pi*3.0/2.0), blue)
	go dragon(m.Rotate2D(1, 0, math.Pi), green)
	go dragon(m.Rotate2D(1, 0, math.Pi/2), orange)
	wg.Add(3)

	wg.Wait()

	println("saving")
	//m.Die(ctx.SavePNG(filename))
	file, err := os.Create(filename)
	m.Die(err)
	m.Die(png.Encode(file, img))
	m.Die(file.Close())
}
