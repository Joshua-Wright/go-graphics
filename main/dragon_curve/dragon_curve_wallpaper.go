package dragon_curve

import (
	m "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/dragon_turtle"
	"github.com/fogleman/gg"
	"image/color"
	"image"
	"math"
	"os"
	"github.com/anthonynsimon/bild/blur"
	"image/draw"
	"fmt"
	"github.com/joshua-wright/go-graphics/parallel"
)

func main() {

	// 16:9 monitor in center flanked by two 5:4 monitors (vertical)
	// equivalent to a 16:9 monitor joined horizontally with a 8:5 monitor
	// need to make 9 equal to 5, so multiply by 5/9
	// so final ratio is 5:(8*16*5/9) = 5:16.88888888
	// for 1440px height, width should be 16.9/5*1440=4867.2
	//width := 4867 * 2
	//height := 1440 * 2

	width := 160
	height := 90
	depth := 20
	line_thickness := 2.0
	neon_glow_radius := 10.0
	neon_glow_factor := 1
	gridsize := 10.0
	middle := m.Vec2{float64(width / 2), float64(height / 2)}

	foldername := m.ExecutableName()
	err := os.Mkdir(foldername, 0777)
	fmt.Println(err)

	syms := dragon_turtle.DragonCurveSymbols(depth)

	dragon := func(dir dragon_turtle.Direction, linecolor color.NRGBA, layername string) image.Image {
		ctx := gg.NewContext(width, height)
		pts2 := dragon_turtle.DragonCurve(syms, middle, dir, dragon_turtle.DiagonalAxes, gridsize)

		arcRadius := (pts2[0].SubV(pts2[1]).Mag()) / 2.0

		ctx.SetColor(linecolor)
		ctx.SetLineWidth(line_thickness)

		for i := 1; i < len(pts2)-1; i++ {
			p0 := pts2[i-1]
			p1 := pts2[i]
			p2 := pts2[i+1]
			center, t0, t1 := DragonCurveCenterAndAngles(p0, p1, p2)
			ctx.DrawArc(center.X, center.Y, arcRadius, t0, t1)
			ctx.Stroke()
		}
		//}
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
	//layers := make([]image.Image, 5)
	layers := make([]image.Image, 1+2*4)
	parallel.ParallelFuncs(
		func() { // black background
			ctx := gg.NewContext(width, height)
			ctx.SetColor(color.NRGBA{0, 0, 0, 255})
			ctx.DrawRectangle(0, 0, float64(width), float64(height))
			ctx.Fill()
			layers[0] = ctx.Image()
			err := ctx.SavePNG(m.ExecutableFolderFileName("background.png"))
			m.Die(err)
		},
		func() {
			layers[1] = dragon(dragon_turtle.NORTH, red, "north red")
		},
		func() {
			layers[2] = dragon(dragon_turtle.SOUTH, orange, "south orange")
		},
		func() {
			layers[3] = dragon(dragon_turtle.EAST, green, "east green")
		},
		func() {
			layers[4] = dragon(dragon_turtle.WEST, blue, "west blue")
		},
	)

	for i := 1; i <= 4; i++ { // run these in serial because they're already parallel
		layers[i+4] = blur.Gaussian(layers[i], neon_glow_radius)
		fmt.Println("blurred", i)
	}
	parallel.ParallelFuncs(
		func() { m.SaveAsPNG(layers[5], m.ExecutableFolderFileName("north red blur.png")) },
		func() { m.SaveAsPNG(layers[6], m.ExecutableFolderFileName("south orange blur.png")) },
		func() { m.SaveAsPNG(layers[7], m.ExecutableFolderFileName("east green blur.png")) },
		func() { m.SaveAsPNG(layers[8], m.ExecutableFolderFileName("west blue blur.png")) },
		func() {
			out := image.NewRGBA(image.Rect(0, 0, width, height))
			draw.Draw(out, layers[0].Bounds(), layers[0], image.ZP, draw.Over)
			for i := 0; i < neon_glow_factor; i++ {
				draw.Draw(out, layers[0].Bounds(), layers[5], image.ZP, draw.Over)
				draw.Draw(out, layers[0].Bounds(), layers[6], image.ZP, draw.Over)
				draw.Draw(out, layers[0].Bounds(), layers[7], image.ZP, draw.Over)
				draw.Draw(out, layers[0].Bounds(), layers[8], image.ZP, draw.Over)
			}
			draw.Draw(out, layers[0].Bounds(), layers[1], image.ZP, draw.Over)
			draw.Draw(out, layers[0].Bounds(), layers[2], image.ZP, draw.Over)
			draw.Draw(out, layers[0].Bounds(), layers[3], image.ZP, draw.Over)
			draw.Draw(out, layers[0].Bounds(), layers[4], image.ZP, draw.Over)
			m.SaveAsPNG(out, m.ExecutableNamePng())
		},
	)
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
