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
	nPts := 256
	overstep := 32
	a := 4.0
	b := 3.0
	linewidth := 2.0
	blur_radius := 8 * linewidth
	maxTheta := (2 * math.Pi)
	bounds := [4]g.Float{
		// a little extra space around the rose
		-1.1, 1.1,
		-1.1, 1.1,
	}

	pts := make([]g.Vec2, nPts)
	for i := 0; i < nPts; i++ {
		theta := float64(i) / float64(nPts) * maxTheta
		x, y := g.WindowTransformPoint(
			width, height,
			g.Vec2{
				math.Sin(a * theta),
				math.Cos(b * theta)},
			bounds)
		pts[i] = g.Vec2{float64(x), float64(y)}
	}

	ctx := gg.NewContext(width, height)
	ctx.SetLineWidth(linewidth)
	for i := 0; i < nPts; i++ {
		t := float64(i) / float64(nPts) * 360.0
		c := colorful.Hsv(t, 1.0, 1.0)
		ctx.SetColor(c)
		p0 := pts[i]
		p1 := pts[(i+nPts/overstep)%nPts]
		p2 := pts[(i+1)%nPts]
		// inner line
		ctx.DrawLine(p0.X, p0.Y, p1.X, p1.Y)
		// outer line
		ctx.DrawLine(p0.X, p0.Y, p2.X, p2.Y)
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
