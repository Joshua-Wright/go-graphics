package main

import (
	"github.com/fogleman/gg"
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
)

// 1/phi
const inv_phi = 1.0/1.618033988749894848204586834365638117720309179805762862135

func main() {
	width := 2048  *2
	height := int(float64(width) * inv_phi)
	scale := float64(width) * 1
	linewidth := 6.0 * 2
	depth := 6

	ctx := gg.NewContext(width, height)
	ctx.SetHexColor("#ffffff")
	ctx.DrawRectangle(0, 0, float64(width), float64(height))
	ctx.Fill()
	ctx.Stroke()

	// initial
	halfKites := []HalfKite{
		{
			A: g.Vec2Zero,
			B: g.Vec2{1, 0},
			C: g.Vec2{0.81, 0.59},
		},
	}
	halfDarts := []HalfDart{}

	// deflate
	for i := 0; i < depth; i++ {
		halfDarts2 := []HalfDart{}
		halfKites2 := []HalfKite{}

		for _, v := range halfDarts {
			halfKites2, halfDarts2 = v.Split(halfKites2, halfDarts2)
		}
		for _, v := range halfKites {
			halfKites2, halfDarts2 = v.Split(halfKites2, halfDarts2)
		}

		halfKites = halfKites2
		halfDarts = halfDarts2
	}

	// scale points
	scalePoint := func(v *g.Vec2) {
		v.X *= scale
		v.Y = float64(height) - v.Y*scale
	}
	for i, _ := range halfKites {
		scalePoint(&halfKites[i].A)
		scalePoint(&halfKites[i].B)
		scalePoint(&halfKites[i].C)
	}
	for i, _ := range halfDarts {
		scalePoint(&halfDarts[i].A)
		scalePoint(&halfDarts[i].B)
		scalePoint(&halfDarts[i].C)
	}

	// draw
	ctx.SetLineCap(gg.LineCapRound)
	ctx.SetLineWidth(linewidth)

	for _, v := range halfKites {
		ctx.SetRGB(1,0,0)
		ctx.MoveTo(v.A.X, v.A.Y)
		ctx.LineTo(v.C.X, v.C.Y)
		ctx.LineTo(v.B.X, v.B.Y)
		ctx.Fill()
		ctx.Stroke()

		ctx.SetRGB(0,0,0)
		ctx.DrawLine(v.A.X, v.A.Y, v.C.X, v.C.Y)
		ctx.DrawLine(v.C.X, v.C.Y, v.B.X, v.B.Y)
		ctx.Stroke()
	}
	for _, v := range halfDarts {
		ctx.SetRGB(0,0,1)
		ctx.MoveTo(v.A.X, v.A.Y)
		ctx.LineTo(v.C.X, v.C.Y)
		ctx.LineTo(v.B.X, v.B.Y)
		ctx.Fill()
		ctx.Stroke()

		ctx.SetRGB(0,0,0)
		ctx.DrawLine(v.A.X, v.A.Y, v.C.X, v.C.Y)
		ctx.DrawLine(v.C.X, v.C.Y, v.B.X, v.B.Y)
		ctx.Stroke()
	}

	r1 := math.Sqrt(halfKites[0].A.SubV(halfKites[0].B).Mag2()) / 8
	ctx.SetHexColor("#ff0000")
	ctx.SetRGB(1, 1, 0)
	for _, v := range halfKites {
		abMid := g.Vec2Lerp(v.B, v.A, inv_phi)
		//abMid := g.Vec2Midpoint(v.B, v.A)
		ctx.DrawCircle(abMid.X, abMid.Y, r1)
		ctx.Fill()
		ctx.Stroke()
	}

	r2 := math.Sqrt(halfDarts[0].A.SubV(halfDarts[0].B).Mag2()) / 8
	ctx.SetRGB(0, 1, 0)
	for _, v := range halfDarts {
		//abMid := g.Vec2Lerp(v.B, v.A, inv_phi)
		abMid := g.Vec2Midpoint(v.B, v.A)
		ctx.DrawCircle(abMid.X, abMid.Y, r2)
		ctx.Fill()
		ctx.Stroke()
	}

	img := ctx.Image()
	g.SaveAsPNG(img, g.ExecutableNamePng())

}

type HalfKite struct {
	A, B, C g.Vec2
}

func (k *HalfKite) Split(halfKites []HalfKite, halfDarts []HalfDart) ([]HalfKite, []HalfDart) {
	abMid := g.Vec2Lerp(k.A, k.B, inv_phi)
	acMid := g.Vec2Lerp(k.C, k.A, inv_phi)
	k1 := HalfKite{
		A: k.C,
		B: abMid,
		C: acMid,
	}
	k2 := HalfKite{
		A: k.C,
		B: abMid,
		C: k.B,
	}
	d1 := HalfDart{
		A: acMid,
		B: k.A,
		C: abMid,
	}
	return append(halfKites, k1, k2), append(halfDarts, d1)
}

type HalfDart struct {
	A, B, C g.Vec2
}

func (d *HalfDart) Split(halfKites []HalfKite, halfDarts []HalfDart) ([]HalfKite, []HalfDart) {
	bcMid := g.Vec2Lerp(d.B, d.C, inv_phi)
	k1 := HalfKite{
		A: d.B,
		B: d.A,
		C: bcMid,
	}
	d1 := HalfDart{
		A: bcMid,
		B: d.C,
		C: d.A,
	}
	return append(halfKites, k1), append(halfDarts, d1)
}
