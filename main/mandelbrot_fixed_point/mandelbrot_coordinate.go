package main

import (
	mandelbrot "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	"os"
	"strconv"
	g "github.com/joshua-wright/go-graphics/graphics"
	"fmt"
	"github.com/ncw/gmp"
)

// usage: go run mandelbrot_coordinate.go width height basePower2 cr ci zoom x y
func main() {
	if len(os.Args) > 9 {
		panic("too many args")
	}
	basePower2 := uint(parseInt64(os.Args[3]))
	zoom, success := new(gmp.Int).SetString(os.Args[6], 10)
	if !success {
		panic("bad zoom")
	}
	x := parseInt64(os.Args[7])
	y := parseInt64(os.Args[8])
	w := parseInt64(os.Args[1])
	h := parseInt64(os.Args[2])
	centerR := mandelbrot.ParseFixnum(os.Args[4], basePower2)
	centerI := mandelbrot.ParseFixnum(os.Args[5], basePower2)
	cr, ci := mandelbrot.MandelbrotCoordinate(
		x, y, w, h,
		centerR, centerI,
		zoom, basePower2)

	fmt.Println("cr =", cr.String(), "/2^", basePower2)
	fmt.Println("ci =", ci.String(), "/2^", basePower2)

	cr1, _ := mandelbrot.MandelbrotCoordinate(
		0, 0, w, h,
		centerR, centerI,
		zoom, basePower2)
	cr2, _ := mandelbrot.MandelbrotCoordinate(
		1, 0, w, h,
		centerR, centerI,
		zoom, basePower2)
	fmt.Println("pixel width = (", cr1.String(), "-", cr2.String(), ")/2^", basePower2)
}

func parseInt64(str string) int64 {
	v, err := strconv.ParseInt(str, 10, 64)
	g.Die(err)
	return v
}
