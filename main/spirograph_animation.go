package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"github.com/fogleman/gg"
	"image"
	"image/draw"
	"github.com/lucasb-eyer/go-colorful"
	"os"
	"fmt"
	"github.com/joshua-wright/go-graphics/parallel"
)

func Hypotrochoid(R, r, d, t float64) g.Vec2 {
	// https://en.wikipedia.org/wiki/Hypotrochoid
	return g.Vec2{
		X: (R-r)*math.Cos(t) + d*math.Cos(t*(R-r)/r),
		Y: (R-r)*math.Sin(t) - d*math.Sin(t*(R-r)/r),
	}
}

func main() {
	width := 600
	height := 600
	//height := 1080
	nPts := 1 << 12
	linewidth := 6.0
	//blur_radius := 8 * linewidth
	maxTheta := (3 * 2 * math.Pi)

	nFrames := 600

	R := 80.0
	r := 30.0
	d := 60.0
	//dtheta := (1.0 / 8.0 * 7.0 / 8.0) * maxTheta
	pointRadius := 10.0
	nCircles := 8

	bound_width := 140.0
	bounds := [4]g.Float{
		// a little extra space around the rose
		-bound_width, bound_width,
		-bound_width, bound_width,
		//-bound_width * 9.0 / 16.0, bound_width * 9.0 / 16.0,
	}

	foldername := g.ExecutableName()
	os.Mkdir(foldername, 0777)

	//transform := g.Rotate2D(0, 0, math.Pi*0.09)

	ctx := gg.NewContext(width, height)
	ctx.SetLineWidth(linewidth)

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

	background := ctx.Image()

	for i := 0; i < nPts-1; i++ {
		t := float64(i) / float64(nPts) * 360
		ctx.SetColor(colorful.Hsv(t, 1.0, 1.0))
		p0 := pts[i]
		p2 := pts[(i+1)%nPts]
		ctx.DrawLine(p0.X, p0.Y, p2.X, p2.Y)
		ctx.Stroke()
	}

	frameFunc := func(frame int) {

		ctx := gg.NewContext(width, height)
		ctx.SetLineWidth(linewidth)
		//ctx.SetColor(color.White)

		for i := 0; i < nCircles; i++ {
			ctx.SetColor(colorful.Hsv(float64(i)/float64(nCircles)*360, 1.0, 1.0))
			//theta := dtheta*float64(i) + float64(frame)/float64(nFrames)*maxTheta
			//theta := (float64(i)+float64(i)/float64(nCircles-1))/float64(nCircles)*maxTheta + float64(frame)/float64(nFrames)*maxTheta
			theta := (float64(i)*1.1)/float64(nCircles)*maxTheta + float64(frame)/float64(nFrames)*maxTheta
			v := Hypotrochoid(R, r, d, theta)
			x, y :=
				g.WindowTransformPoint(
					width, height,
					v,
					bounds)
			//println(frame, i, x, y)
			ctx.DrawCircle(float64(x), float64(y), pointRadius)
			ctx.Stroke()
		}

		foreground := ctx.Image()
		out := image.NewRGBA(background.Bounds())
		draw.Draw(out, out.Bounds(), image.Black, image.ZP, draw.Over)
		draw.Draw(out, out.Bounds(), background, image.ZP, draw.Over)
		//// make blurred background image
		//blurred := imaging.Blur(img, blur_radius)
		//draw.Draw(out, out.Bounds(), blurred, image.ZP, draw.Over)
		draw.Draw(out, out.Bounds(), foreground, image.ZP, draw.Over)
		g.SaveAsPNG(out, g.ExecutableFolderFileName(fmt.Sprint(frame)+".png"))
		fmt.Println(frame)
	}

	parallel.ParallelFor(0, nFrames, frameFunc)
}
