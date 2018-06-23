package mandelbrot_fixed_point

import (
	"github.com/ncw/gmp"
	"strconv"
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
)

func MandelbrotKernel(cr, ci, threshold2 *gmp.Int, maxIter int64, basePower2 uint) (i int64, v float64, mag2 *gmp.Int) {
	zr := new(gmp.Int)
	zi := new(gmp.Int)
	zr2 := new(gmp.Int)
	zi2 := new(gmp.Int)
	magnitude2 := new(gmp.Int)
	zri := new(gmp.Int)
	for i := int64(0); i < maxIter; i++ {
		zr2.Mul(zr, zr)
		zi2.Mul(zi, zi)
		magnitude2.Add(zr2, zi2)
		if magnitude2.CmpAbs(threshold2) > 0 {
			// smooth (but less precise) coloring
			z2, err := strconv.ParseFloat(magnitude2.String(), 64)
			g.Die(err)
			v := float64(i-1) - math.Log2(math.Log2(z2)-float64(basePower2*2)) + 1
			return i, v, magnitude2
		}
		// otherwise re-use those values

		// calculate 2ab/B + ci. Instead of multiplying by 2, just shift right by one less
		// (calculate out of place to not disturb zi for the next calculation)
		zri.Mul(zr, zi).Rsh(zri, basePower2-1).Add(zri, ci)

		// calculate (a^2 - b^2)/B + cr (in place this time)
		zr.Sub(zr2, zi2).Rsh(zr, basePower2).Add(zr, cr)

		// swap zri and zi. zi is now the real value, and zri can maybe have its memory re-used next iteration
		// TODO: check if the memory is really re-used
		zri, zi = zi, zri
	}
	// explicit not in set sentinel value
	return 0, -1.0, nil
}

func MandelbrotCoordinate(x, y, width, height int64, centerR, centerI, zoom *gmp.Int,
	basePower2 uint) (cr *gmp.Int, ci *gmp.Int) {
	r := gmp.NewInt(x - width/2)
	i := gmp.NewInt(height/2 - y)

	denom := gmp.NewInt(width)
	denom.Mul(denom, zoom)

	cr = gmp.NewInt(4)
	cr.Mul(cr, r).Lsh(cr, basePower2).Div(cr, denom).Add(cr, centerR)

	ci = gmp.NewInt(4)
	ci.Mul(ci, i).Lsh(ci, basePower2).Div(ci, denom).Add(ci, centerI)

	return cr, ci
}

// TODO: implement per_pixel_image.PixelFunction
