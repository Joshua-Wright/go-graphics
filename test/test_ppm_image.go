package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/draw"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
)

func main() {
	width := 512
	height := 512
	nPts := 256
	a := 4.0
	b := 3.0
	linewidth := 2.0
	//blur_radius := 8 * linewidth
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
		p2 := pts[(i+1)%nPts]
		ctx.DrawLine(p0.X, p0.Y, p2.X, p2.Y)
		ctx.Stroke()
	}

	ctx.SavePNG(g.ExecutableNamePng())

	img := ctx.Image()
	// make blurred background image
	//blurred := imaging.Blur(img, blur_radius)

	// flatten layers
	//out := image.NewRGBA(img.Bounds())
	out, err := memory_mapped.CreatePPM(int64(width), int64(height), g.ExecutableNameWithExtension("ppm"))
	g.Die(err)
	draw.Draw(out, out.Bounds(), image.Black, image.ZP, draw.Over)
	//draw.Draw(out, out.Bounds(), blurred, image.ZP, draw.Over)
	draw.Draw(out, out.Bounds(), img, image.ZP, draw.Over)
	g.SaveAsPNG(out, g.ExecutableNamePng())
	out.Close()

	// open the image and make some changes
	out2, err := memory_mapped.OpenPPM(g.ExecutableNameWithExtension("ppm"))
	g.Die(err)
	defer out2.Close()
	draw.Draw(out2, image.Rect(10, 20, 100, 100), image.White, image.Pt(100, 100), draw.Over)
}
