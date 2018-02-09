package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"github.com/fogleman/gg"
	"image"
	"image/draw"
	"image/color"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/disintegration/imaging"
)

func Hypotrochoid(R, r, d, t float64) g.Vec2 {
	// https://en.wikipedia.org/wiki/Hypotrochoid
	return g.Vec2{
		X: (R-r)*math.Cos(t) + d*math.Cos(t*(R-r)/r),
		Y: (R-r)*math.Sin(t) - d*math.Sin(t*(R-r)/r),
	}
}

func main() {
	width := 1920
	height := 1920
	//height := 1080
	nPts := 1 << 15
	linewidth := 6.0
	blur_radius := 8 * linewidth
	maxTheta := (5 * 2 * math.Pi)

	bound_width := 100.0
	bounds := [4]g.Float{
		// a little extra space around the rose
		-bound_width, bound_width,
		-bound_width, bound_width,
		//-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	//transform := g.Rotate2D(0, 0, math.Pi*0.09)

	ctx := gg.NewContext(width, height)
	ctx.SetLineWidth(linewidth)

	params := func(R, r, d float64, c color.Color) {

		pts := make([]g.Vec2, nPts)
		for i := 0; i < nPts; i++ {
			theta := float64(i)/float64(nPts)*maxTheta + 38
			v := Hypotrochoid(R, r, d, theta)
			x, y :=
				g.WindowTransformPoint(
					width, height,
					v,
					bounds)
			pts[i] = g.Vec2{float64(x), float64(y)}
		}

		for i := 0; i < nPts-1; i++ {
			ctx.SetColor(c)
			p0 := pts[i]
			p2 := pts[(i+1)%nPts]
			ctx.DrawLine(p0.X, p0.Y, p2.X, p2.Y)
			ctx.Stroke()
		}
	}

	params(105, 30, 19, colorful.Hsv(0, 1.0, 1.0))
	params(105, 45, 20, colorful.Hsv(90, 1.0, 1.0))
	params(105, 60, 25, colorful.Hsv(180, 1.0, 1.0))
	params(105, 75, 30, colorful.Hsv(270, 1.0, 1.0))

	img := ctx.Image()

	// flatten layers
	out := image.NewRGBA(img.Bounds())
	draw.Draw(out, out.Bounds(), image.Black, image.ZP, draw.Over)
	// make blurred background image
	blurred := imaging.Blur(img, blur_radius)
	draw.Draw(out, out.Bounds(), blurred, image.ZP, draw.Over)
	draw.Draw(out, out.Bounds(), img, image.ZP, draw.Over)
	g.SaveAsPNG(out, g.ExecutableNamePng())
}
