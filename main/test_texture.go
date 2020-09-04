package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/texture"
	"image"
	"math"
)

func main() {
	//inputFilename := os.Args[1]
	tex, err := texture.TextureFromImage("768px-Hexagons.svg.pngq")
	g.Die(err)

	b := tex.Img.Bounds()
	img := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))

	center := g.Vec2{
		X: float64(b.Dx() / 2),
		Y: float64(b.Dy() / 2),
	}

	for i := 0; i < b.Dx(); i++ {
		for j := 0; j < b.Dy(); j++ {
			px := g.Vec2{float64(i), float64(j)}
			r := px.SubV(center).Mag() / float64(b.Dx()) * 2
			rhat := px.SubV(center).UnitV()
			theta := math.Atan2(rhat.Y, rhat.X) / (2 * math.Pi)
			if theta < 0 {
				theta += 1
			}
			pix := tex.GetPixel(g.Vec2{theta, r})
			img.Set(i, j, pix)
		}
	}

	g.SaveAsPNG(img, "test_texture.png")
}
