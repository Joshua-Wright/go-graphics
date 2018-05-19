package main

import (
	"github.com/fogleman/gg"
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Hexagon(ctx *gg.Context, x, y, size float64) {
	height := size / (math.Sqrt(3) / 2)
	width := size
	ctx.MoveTo(x, y+height)
	ctx.LineTo(x+width, y+height/2)
	ctx.LineTo(x+width, y-height/2)
	ctx.LineTo(x, y-height)
	ctx.LineTo(x-width, y-height/2)
	ctx.LineTo(x-width, y+height/2)
	ctx.LineTo(x, y+height)
	ctx.Fill()
	ctx.Stroke()
}

func main() {
	width := 2048
	height := 2048
	n_rows := 100
	line_spacing := float64(width)/float64(n_rows-1)
	line_height := line_spacing * (math.Sqrt(3) / 2)
	prime_color := "#ff0000"
	non_prime_color := "#000000"

	ctx := gg.NewContext(width, height)
	ctx.SetHexColor("#ffffff")
	ctx.DrawRectangle(0, 0, float64(width), float64(height))
	ctx.Fill()
	ctx.Stroke()

	pack_has_primes := func(x, y, curNum int) bool {
		myNum := curNum + x
		for i := 0; i < n_rows; i++ {
			if isPrime(myNum) {
				return true
			}
			myNum += y
		}
		return false
	}

	curNum := 1
	for y := 1; y < n_rows; y++ {
		for x := 0; x < y; x++ {
			if pack_has_primes(x, y, curNum) {
				ctx.SetHexColor(prime_color)
			} else {
				ctx.SetHexColor(non_prime_color)
			}
			offset := float64(n_rows-y) / 2 * line_spacing
			Hexagon(ctx, offset+line_spacing*float64(x), line_height*float64(y), line_spacing/2+0.5)
		}
		curNum += n_rows * y
	}

	img := ctx.Image()
	g.SaveAsPNG(img, g.ExecutableNamePng())
}
