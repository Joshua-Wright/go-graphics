package mandelbrot_fixed_point

import (
	"github.com/ncw/gmp"
)

func MandelbrotKernel(cr, ci, threshold2 *gmp.Int, maxIter int64, base_power_2 uint) float64 {
	zr := &gmp.Int{}
	zi := &gmp.Int{}
	zr2 := &gmp.Int{}
	zi2 := &gmp.Int{}
	magnitude := &gmp.Int{}
	zri := &gmp.Int{}
	for i := int64(0); i < maxIter; i++ {
		zr2.Mul(zr, zr)
		zi2.Mul(zi, zi)
		magnitude.Add(zr2, zi2)
		if magnitude.CmpAbs(threshold2) > 0 {
			// TODO smooth coloring
			return float64(i)
		}
		// otherwise re-use those values

		// calculate 2ab/B + ci. Instead of multiplying by 2, just shift right by one less
		// (calculate out of place to not disturb zi for the next calculation)
		//zri.Mul(zr, zi).Rsh(zri, base_power_2-1).Add(zri, ci)
		zri.Mul(zr, zi).MulUint32(zri, 2).Rsh(zri, base_power_2).Add(zri, ci)

		// calculate (a^2 - b^2)/B + cr (in place this time)
		zr.Sub(zr2, zi2).Rsh(zr, base_power_2).Add(zr, cr)

		// swap zri and zi. zi is now the real value, and zri can maybe have its memory re-used next iteration
		// TODO: check if the memory is really re-used
		zri, zi = zi, zri
	}
	// explicit not in set sentinel value
	return -1.0
}

func MandelbrotCoordinate(x, y, width, height int64, centerR, centerI, zoom *gmp.Int,
	base_power_2 uint) (cr *gmp.Int, ci *gmp.Int) {
	r := gmp.NewInt(x - width/2)
	i := gmp.NewInt(height/2 - y)

	denom := gmp.NewInt(width)
	denom.Mul(denom, zoom)

	cr = gmp.NewInt(4)
	cr.Mul(cr, r).Lsh(cr, base_power_2).Div(cr, denom).Add(cr, centerR)

	ci = gmp.NewInt(4)
	ci.Mul(ci, i).Lsh(ci, base_power_2).Div(ci, denom).Add(ci, centerI)

	return cr, ci
}

// TODO: implement per_pixel_image.PixelFunction
