package main

import (
	"github.com/fogleman/gg"
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"fmt"
)

func main() {
	width := 2048
	height := width
	lineWidth := 6.0 * 2

	ctx := gg.NewContext(width, height)
	ctx.SetRGB(1, 1, 1)
	ctx.DrawRectangle(0, 0, float64(width), float64(height))
	ctx.Fill()
	ctx.Stroke()

	ctx.SetRGB(0, 0, 0)
	ctx.SetLineWidth(lineWidth)

	rfact := 0.9
	arc := func(a1, a2 float64) {
		ctx.DrawArc(float64(width/2), float64(height/2), float64(width/2)*rfact,
			a1, a2)
		ctx.Stroke()
		fmt.Println(rfact, a1, a2)
		rfact -= 0.1
	}

	arc(-math.Pi, math.Pi*0.8)
	arc(MinimalAngleMapping(-math.Pi, math.Pi*0.8))
	arc(0.8*math.Pi, 0.6*math.Pi)
	arc(MinimalAngleMapping(0.8*math.Pi, 0.6*math.Pi))

	img := ctx.Image()
	g.SaveAsPNG(img, g.ExecutableNamePng())

}

func MinimalAngleMapping(angle1, angle2 float64) (float64, float64) {
	type p struct {
		a1, a2 float64
	}
	angles := []p{
		{angle1, angle2},
		{angle1 + 2*math.Pi, angle2},
		{angle1 - 2*math.Pi, angle2},
		{angle1, angle2 + 2*math.Pi},
		{angle1, angle2 - 2*math.Pi},
	}
	minDiff := 99999.0
	a1 := 0.0
	a2 := 0.0
	for i := 0; i < len(angles); i++ {
		diff := math.Abs(angles[i].a2 - angles[i].a1)
		fmt.Println(i, angles[i], diff, minDiff)
		if diff < minDiff {
			minDiff = diff
			a1 = angles[i].a1
			a2 = angles[i].a2
		}
	}
	return a1, a2
}
