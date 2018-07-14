package main

import (
	"github.com/fogleman/gg"
	g "github.com/joshua-wright/go-graphics/graphics"
)

// 1/phi
const inv_phi = 0.6180339887498948482

func main() {
	width := 2048 * 2
	height := int(float64(width) * inv_phi)
	scale := float64(width) * 1
	linewidth := 3.0
	depth := 7

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

	// draw
	ctx.SetHexColor("#000000")
	ctx.SetLineWidth(linewidth)
	for _, v := range halfKites {
		ctx.MoveTo(v.A.X*scale, v.A.Y*scale)
		ctx.LineTo(v.B.X*scale, v.B.Y*scale)
		ctx.LineTo(v.C.X*scale, v.C.Y*scale)
		ctx.LineTo(v.A.X*scale, v.A.Y*scale)
		ctx.Stroke()
	}
	for _, v := range halfDarts {
		ctx.MoveTo(v.A.X*scale, v.A.Y*scale)
		ctx.LineTo(v.B.X*scale, v.B.Y*scale)
		ctx.LineTo(v.C.X*scale, v.C.Y*scale)
		ctx.LineTo(v.A.X*scale, v.A.Y*scale)
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
		B: abMid,
		C: k.A,
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
