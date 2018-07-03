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
	cr,ci := mandelbrot.MandelbrotCoordinate(
		parseInt64(os.Args[7]),
		parseInt64(os.Args[8]),
		parseInt64(os.Args[1]),
		parseInt64(os.Args[2]),
		mandelbrot.ParseFixnum(os.Args[4], basePower2),
		mandelbrot.ParseFixnum(os.Args[5], basePower2),
		zoom,
		basePower2,
	)
	fmt.Println("cr =", cr.String(), "/2^", basePower2)
	fmt.Println("ci =", ci.String(), "/2^", basePower2)
}


func parseInt64(str string) int64 {
	v, err := strconv.ParseInt(str, 10, 64)
	g.Die(err)
	return v
}