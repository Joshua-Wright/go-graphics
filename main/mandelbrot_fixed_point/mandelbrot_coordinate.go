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
	x := parseInt64(os.Args[7])
	y := parseInt64(os.Args[8])
	width := parseInt64(os.Args[1])
	height := parseInt64(os.Args[2])
	centerR := os.Args[4]
	centerI := os.Args[5]
	zoom := os.Args[6]

	decimalDigits := 128

	///
	centerRgmp := mandelbrot.ParseFixnum(centerR, basePower2)
	centerIgmp := mandelbrot.ParseFixnum(centerI, basePower2)
	zoomGmp, success := new(gmp.Int).SetString(zoom, 10)
	if !success {
		panic("bad zoomGmp")
	}
	cr, ci := mandelbrot.MandelbrotCoordinate(
		x, y, width, height,
		centerRgmp, centerIgmp,
		zoomGmp, basePower2)

	fmt.Println("cr =", cr.String(), "/2^", basePower2)
	fmt.Println("ci =", ci.String(), "/2^", basePower2)

	cr1, _ := mandelbrot.MandelbrotCoordinate(
		0, 0, width, height,
		centerRgmp, centerIgmp,
		zoomGmp, basePower2)
	cr2, _ := mandelbrot.MandelbrotCoordinate(
		1, 0, width, height,
		centerRgmp, centerIgmp,
		zoomGmp, basePower2)
	fmt.Println("pixel width = (", cr1.String(), " - ", cr2.String(), ")/2^", basePower2)

	// mathematica-flavored output
	fmt.Println()
	fmt.Println()
	fmt.Println("decimalDigits=", decimalDigits)
	fmt.Println("decimal[x_] := DecimalForm[N[x,decimalDigits],decimalDigits]")
	fmt.Println("cr := ", cr.String(), "/2^", basePower2)
	fmt.Println("ci := ", ci.String(), "/2^", basePower2)
	fmt.Println("dx := (", cr1.String(), " - ", cr2.String(), ")/2^", basePower2)
	fmt.Println("decimal[cr]")
	fmt.Println("decimal[ci]")
	fmt.Println("decimal[dx]")
	fmt.Println()
}

func parseInt64(str string) int64 {
	v, err := strconv.ParseInt(str, 10, 64)
	g.Die(err)
	return v
}
