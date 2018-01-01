package main

import (
	m ".."
	"os"
	"image/png"
	"image"
	"image/color"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	X       = iota
	Y
	FORWARD
	LEFT     // -
	RIGHT    // +
)

const (
	EAST  = iota
	NORTH
	WEST
	SOUTH
)

func IterateDragonCurve(instructions []int) []int {
	out := []int{}
	for _, move := range instructions {
		switch move {
		case X:
			out = append(out, X, RIGHT, Y, FORWARD, RIGHT)
		case Y:
			out = append(out, LEFT, FORWARD, X, LEFT, Y)
		case FORWARD, LEFT, RIGHT:
			out = append(out, move)
		}
	}
	return out
}

func main() {
	width := 1920
	height := 1080
	depth := 23

	moves := []int{FORWARD, X}
	for i := 0; i < depth; i++ {
		moves = IterateDragonCurve(moves)
	}

	x := width / 2
	y := height / 2

	dragon_curve := func(img *image.NRGBA, x, y, direction int, c colorful.Color) {
		for _, move := range moves {
			switch move {
			case X, Y:
				continue

			case LEFT:
				direction = (direction + 1) % 4
				continue
			case RIGHT:
				direction = (direction + 3) % 4
				continue

			case FORWARD:
				img.Set(x, y, c)
				switch direction {
				case EAST:
					x++
				case NORTH:
					y++
				case WEST:
					x--
				case SOUTH:
					y--
				}
			}

		}
	}

	//colors, err := colorful.WarmPalette(4)
	//m.Die(err)
	red := colorful.MakeColor(color.NRGBA{255, 35, 24, 255})
	orange := colorful.MakeColor(color.NRGBA{255, 133, 24, 255})
	green := colorful.MakeColor(color.NRGBA{22, 237, 48, 255})
	blue := colorful.MakeColor(color.NRGBA{23, 206, 224, 255})
	colors := []colorful.Color{
		red,
		orange,
		green,
		blue,
	}

	imgs := make([]*image.NRGBA, 4)
	m.ParallelFor(0, 4, func(i int) {
		img := image.NewNRGBA(image.Rect(0, 0, width, height))
		dragon_curve(img, x, y, i, colors[i])
		imgs[i] = img
	})

	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	m.ParallelFor(0, height, func(y int) {
		for x := 0; x < width; x++ {

			// get colors
			colors := []colorful.Color{}
			for i := 0; i < 4; i++ {
				c := imgs[i].NRGBAAt(x, y)
				if c.A != 0 {
					colors = append(colors, colorful.MakeColor(c))
				}
			}
			// blend them
			switch len(colors) {
			case 0:
				img.Set(x, y, color.Black)
			case 1:
				img.Set(x, y, colors[0])
			case 2:
				img.Set(x, y, colors[0].BlendRgb(colors[1], 0.5))
			case 3:
				img.Set(x, y, colors[0].BlendRgb(colors[1], 0.5).BlendRgb(colors[3], 0.5))
			case 4:
				img.Set(x, y, colors[0].BlendRgb(colors[1], 0.5).BlendRgb(colors[3], 0.5).BlendRgb(colors[3], 0.5))
			}

			//// no blending
			//for i := 0; i < 4; i++ {
			//	c := imgs[i].NRGBAAt(x, y)
			//	if c.A != 0 {
			//		img.Set(x, y, c)
			//		break
			//	}
			//}
		}
	})

	file, err := os.Create(m.ExecutableNamePng())
	m.Die(err)
	m.Die(png.Encode(file, img))
	m.Die(file.Close())

	//// save each dragon in a layer
	//os.Mkdir(m.ExecutableName(), 0777)
	//for i := 0; i < 4; i++ {
	//	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	//	dragon_curve(img, x, y, i, colors[i])
	//	file, err := os.Create(fmt.Sprintf("%s/%d.png", m.ExecutableName(), i))
	//	m.Die(err)
	//	m.Die(png.Encode(file, img))
	//	m.Die(file.Close())
	//}
}
