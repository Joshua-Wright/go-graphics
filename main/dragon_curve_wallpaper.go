package main

import (
	m "github.com/joshua-wright/image_stuff/go"
	"github.com/joshua-wright/image_stuff/go/dragon_turtle"
	"github.com/fogleman/gg"
	"github.com/anthonynsimon/bild/blend"
	"image/color"
	"image"
	"math"
	"fmt"
	"os"
)

func main() {

	// 16:9 monitor in center flanked by two 5:4 monitors (vertical)
	// equivalent to a 16:9 monitor joined horizontally with a 8:5 monitor
	// need to make 9 equal to 5, so multiply by 5/9
	// so final ratio is 5:(8*16*5/9) = 5:16.88888888
	// for 1440px height, width should be 16.9/5*1440=4867.2
	width := 4867
	height := 1440
	depth := 16
	line_thickness := 3.0
	//neon_glow_radius := 15
	neon_glow_radius := 1
	gridsize := 20.0
	middle := m.Vec2{float64(width / 2), float64(height / 2)}

	foldername := m.ExecutableName()
	os.Mkdir(foldername, 0777)

	syms := dragon_turtle.DragonCurveSymbols(depth)

	dragon := func(dir dragon_turtle.Direction, linecolor color.NRGBA, layername string) image.Image {
		ctx := gg.NewContext(width, height)
		pts2 := dragon_turtle.DragonCurve(syms, middle, dir, gridsize)

		arcRadius := (pts2[0].SubV(pts2[1]).Mag()) / 2.0

		// neon glow effect
		for i := 1; i <= neon_glow_radius; i++ {
			if i != neon_glow_radius {
				c2 := linecolor
				// limit max alpha for glow
				c2.A = uint8(float64(i) / float64(neon_glow_radius) * 24.0)
				lw := line_thickness + float64(neon_glow_radius-i)
				ctx.SetColor(c2)
				ctx.SetLineWidth(lw)
				fmt.Println("i:", i, "linewidth:", lw, "alpha:", c2.A)
			} else {
				ctx.SetColor(linecolor)
				ctx.SetLineWidth(line_thickness)
			}

			for i := 1; i < len(pts2)-1; i++ {
				p0 := pts2[i-1]
				p1 := pts2[i]
				p2 := pts2[i+1]
				center, t0, t1 := DragonCurveCenterAndAngles(p0, p1, p2)
				ctx.DrawArc(center.X, center.Y, arcRadius, t0, t1)
				ctx.Stroke()
			}
		}
		println("drew lines")

		err := ctx.SavePNG(m.ExecutableFolderFileName(layername + ".png"))
		m.Die(err)
		println("saved")
		return ctx.Image()
	}

	// do 4 rotations with different colors
	red := color.NRGBA{255, 35, 24, 255}
	orange := color.NRGBA{255, 133, 24, 255}
	green := color.NRGBA{22, 237, 48, 255}
	blue := color.NRGBA{23, 206, 224, 255}
	layers := make([]image.Image, 5)
	m.Parallel(
		func() { // black background
			ctx := gg.NewContext(width, height)
			ctx.SetColor(color.NRGBA{0, 0, 0, 255})
			ctx.DrawRectangle(0, 0, float64(width), float64(height))
			ctx.Fill()
			layers[0] = ctx.Image()
			err := ctx.SavePNG(m.ExecutableFolderFileName("background.png"))
			m.Die(err)
		},
		func() { layers[1] = dragon(dragon_turtle.NORTH, red, "north red") },
		func() { layers[2] = dragon(dragon_turtle.SOUTH, orange, "south orange") },
		func() { layers[3] = dragon(dragon_turtle.EAST, green, "east green") },
		func() { layers[4] = dragon(dragon_turtle.WEST, blue, "west blue") },
	)
	out := blend.Normal(layers[0], layers[1])
	for i := 2; i < len(layers); i++ {
		out = blend.Normal(out, layers[i])
	}
	m.SaveAsPNG(out, m.ExecutableNamePng())
}

func DragonCurveCenterAndAngles(start, mid, end m.Vec2) (arcCenter m.Vec2, theta1 float64, theta2 float64) {
	arcCenter = m.Vec2{
		X: (start.X + end.X) / 2.0,
		Y: (start.Y + end.Y) / 2.0,
	}
	start = start.SubV(arcCenter)
	end = end.SubV(arcCenter)
	mid = mid.SubV(arcCenter)

	theta_mid := math.Atan2(mid.Y, mid.X)
	if mid.Cross(start) > 0.0 {
		theta1 = theta_mid - math.Pi/4
		theta2 = theta_mid + math.Pi/4
	} else {
		theta1 = theta_mid + math.Pi/4
		theta2 = theta_mid - math.Pi/4
	}
	return arcCenter, theta1, theta2
}
