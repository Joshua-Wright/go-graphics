package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/draw"
	"github.com/disintegration/imaging"
)

func main() {
	width := 1920 * 2
	height := 1080 * 2
	nPts := 1 << 21
	//overstep := 32
	a := 1.0
	b := 1.04
	linewidth := 4.0
	blur_radius := 8 * linewidth
	maxTheta := (80 * math.Pi)
	d := 0.015

	bound_width := 0.7
	bounds := [4]g.Float{
		// a little extra space around the rose
		-bound_width, bound_width,
		-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	transform := g.Rotate2D(0, 0, math.Pi*0.09)

	pts := make([]g.Vec2, nPts)
	for i := 0; i < nPts; i++ {
		theta := float64(i)/float64(nPts)*maxTheta + 38
		v := g.Vec2{
			math.Sin(a * theta),
			math.Cos(b * theta),
		}
		v = v.MulS(math.Exp(-theta * d))
		v = transform.TransformPoint(&v)
		x, y :=
			g.WindowTransformPoint(
				width, height,
				v,
				bounds)
		pts[i] = g.Vec2{float64(x), float64(y)}
	}

	ctx := gg.NewContext(width, height)
	ctx.SetLineWidth(linewidth)
	for i := 0; i < nPts-1; i++ {
		t := float64(i) / float64(nPts) * 360.0
		c := colorful.Hsv(t, 1.0, 1.0)
		ctx.SetColor(c)
		p0 := pts[i]
		p2 := pts[(i+1)%nPts]
		ctx.DrawLine(p0.X, p0.Y, p2.X, p2.Y)
		ctx.Stroke()
	}

	//ctx.SavePNG(g.ExecutableNamePng())

	img := ctx.Image()
	// make blurred background image
	blurred := imaging.Blur(img, blur_radius)

	// flatten layers
	out := image.NewRGBA(img.Bounds())
	draw.Draw(out, out.Bounds(), image.Black, image.ZP, draw.Over)
	draw.Draw(out, out.Bounds(), blurred, image.ZP, draw.Over)
	draw.Draw(out, out.Bounds(), img, image.ZP, draw.Over)
	g.SaveAsPNG(out, g.ExecutableNamePng())
}
