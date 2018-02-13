package main

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"github.com/joshua-wright/go-graphics/graphics/distance_transform"
	"image"
	"image/draw"
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
	"fmt"
	"os"
	"github.com/joshua-wright/go-graphics/parallel"
)

func KochCurve(x1, y1, x2, y2, theta float64) []g.Matrix3 {
	denom := 2 + 2*g.Cos(theta)
	return []g.Matrix3{
		g.Scale2D(x1, y1, 1/denom),
		g.Translate2D((x2-x1)/denom, (y2-y1)/denom).
			Compose(g.Rotate2D(x1, y1, theta)).
			Compose(g.Scale2D(x1, y1, 1/denom)),
		g.Translate2D((x1-x2)/denom, (y1-y2)/denom).
			Compose(g.Rotate2D(x2, y2, -theta)).
			Compose(g.Scale2D(x2, y2, 1/denom)),
		g.Scale2D(x2, y2, 1/denom),
	}
}

// these are in format [x1,y1,x2,y2]
var SidesPointingIn = [4][4]float64{
	// inner diagonals pointing in
	{0, -1, 1, 0},
	{1, 0, 0, 1},
	{0, 1, -1, 0},
	{-1, 0, 0, -1},
}
var SidesPointingOut = [4][4]float64{
	{1, 0, 0, -1},
	{0, 1, 1, 0},
	{-1, 0, 0, 1},
	{0, -1, -1, 0},
}

func main() {
	width := 800
	height := 800
	depth := 11

	foldername := g.ExecutableName()
	err := os.Mkdir(foldername, 0777)
	fmt.Println(err)

	bound_width := 1.1
	bounds := [4]g.Float{
		-bound_width, bound_width,
		-bound_width, bound_width,
	}

	frameFunc := func(theta float64) *image.NRGBA {

		distancePts := []g.Vec2i{}

		//g.ParallelFor(0, len(SidesPointingOut), func(i int) {
		for i := 0; i < len(SidesPointingOut); i++ {
			mats := KochCurve(
				SidesPointingOut[i][0],
				SidesPointingOut[i][1],
				SidesPointingOut[i][2],
				SidesPointingOut[i][3],
				theta)
			pts := g.TransformPoints([]g.Vec2{g.Vec2Zero}, mats, depth)
			for _, p := range pts {
				x, y := g.WindowTransformPoint(width, height, p, bounds)
				if x < 0 {
					x = 0
				}
				if y < 0 {
					y = 0
				}
				if x >= width {
					x = width - 1
				}
				if y >= height {
					y = height - 1
				}
				distancePts = append(distancePts, g.Vec2i{x, y})
			}
		}
		//})
		distanceGrid := distance_transform.DistanceTransform(width, height, distancePts)

		max_dist := 0.0
		// find largest distance to normalize distanceGrid
		for i := 0; i < len(distanceGrid); i++ {
			for j := 0; j < len(distanceGrid[i]); j++ {
				if max_dist < distanceGrid[i][j] {
					max_dist = distanceGrid[i][j]
				}
			}
		}

		out := image.NewNRGBA(image.Rect(0, 0, width, height))
		draw.Draw(out, out.Bounds(), image.Black, image.ZP, draw.Over)

		parallel.ParallelFor(0, len(distanceGrid), func(i int) {
			//for i := 0; i < len(distanceGrid); i++ {
			for j := 0; j < len(distanceGrid[i]); j++ {
				theta_deg := distanceGrid[i][j] / max_dist * 360
				c := colorful.Hsv(theta_deg, 1.0, 1.0)
				out.Set(i, j, c)
			}
			//}
		})

		for _, p := range distancePts {
			out.Set(p.X, p.Y, color.Black)
		}
		//g.SaveAsPNG(out, g.ExecutableNamePng())
		return out
	}
	N := 300
	//g.ParallelFor(0, N, func(i int) {
	for i := 0; i < N; i++ {
		theta := math.Sin(float64(i)/float64(N)*math.Pi-math.Pi/2) * math.Pi / 2
		fmt.Println(i)
		img := frameFunc(theta)
		g.SaveAsPNG(img, g.ExecutableFolderFileName(fmt.Sprint(N+i))+".png")
		g.SaveAsPNG(img, g.ExecutableFolderFileName(fmt.Sprint(N-i))+".png")
	}
	//})

}
