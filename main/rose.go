package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
)

func main() {
	width := 2048
	height := 2048
	nPts := 8192
	n := 2.0
	d := 3.0
	linewidth := 32.0
	blur_radius := 8 * linewidth
	maxTheta := (2 * math.Pi) * 3
	bounds := [4]g.Float{
		// a little extra space around the rose
		-1.1, 1.1,
		-1.1, 1.1,
	}

	k := float64(n) / float64(d)
	pts := make([]g.Vec2, nPts)
	for i := 0; i < nPts; i++ {
		theta := float64(i) / float64(nPts) * maxTheta
		x, y := g.WindowTransformPoint(
			width, height,
			g.Vec2{
				math.Cos(k*theta) * math.Cos(theta),
				math.Cos(k*theta) * math.Sin(theta)},
			bounds)
		pts[i] = g.Vec2{float64(x), float64(y)}
	}

	ctx := gg.NewContext(width, height)
	ctx.SetLineWidth(linewidth)
	for i := 1; i < nPts; i++ {
		t := float64(i) / float64(nPts) * 360.0
		c := colorful.Hsv(t, 1.0, 1.0)
		ctx.SetColor(c)
		ctx.DrawLine(
			pts[i-1].X, pts[i-1].Y,
			pts[i].X, pts[i].Y)
		ctx.Stroke()
	}

	ctx.SavePNG(g.ExecutableNamePng())

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
