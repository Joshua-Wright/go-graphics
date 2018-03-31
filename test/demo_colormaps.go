package main

import (
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"image"
	g "github.com/joshua-wright/go-graphics/graphics"
)

const HeightPerClormap = 128

func main() {
	cmaps := []colormap.ColorMap{
		colormap.NewInterpColormap(colormap.InfernoColorMap()),
		colormap.NewXyzInterpColormap(colormap.InfernoColorMap()),
		colormap.NewInterpColormap(colormap.JetColorMap()),
		colormap.NewXyzInterpColormap(colormap.JetColorMap()),
		colormap.NewInterpColormap(colormap.UltraFractalColors16),
		colormap.NewXyzInterpColormap(colormap.UltraFractalColors16),
		colormap.HotColormap,
		colormap.Hsv,
		colormap.Sinebow,
		colormap.DefaultCosine3D,
	}

	width := 2560
	height := HeightPerClormap * len(cmaps)

	out := image.NewRGBA(image.Rect(0, 0, width, height))

	for idx, cmap := range cmaps {
		for i := 0; i < width; i++ {
			x := float64(i) / float64(width)
			c := cmap.GetColor(x)
			for j := idx * HeightPerClormap; j < (idx+1)*HeightPerClormap; j++ {
				out.SetRGBA(i, j, c)
			}
		}
	}

	g.SaveAsPNG(out, g.ExecutableNamePng())

}
