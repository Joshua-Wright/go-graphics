package naive_fixnum_32

import (
	"math"
	g "github.com/joshua-wright/go-graphics/graphics"
	"testing"
	"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
)

const BasePower2 = (FpWords - 1) * 32

const centerRstr = "-0.7490868103567926827901785714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714286"
const centerIstr = "0.000354852939810431279017857142857142857142857142857142857142857142857142857142857142857142857142857142857142857142857142857142857"

//const centerRstr = "-1.23048000000000000000000000000000000000000"
//const centerIstr = "-0.02471000000000000000000000000000000000000"
const Thresholdstr = "32.0"
const MaxIter = 128

func BenchmarkFixnum(b *testing.B) {
	cr, err := FromString(centerRstr)
	g.Die(err)
	ci, err := FromString(centerIstr)
	g.Die(err)
	threshold2, err := FromString(Thresholdstr)
	threshold2.Mul(threshold2, threshold2)
	g.Die(err)
	for n := 0; n < b.N; n++ {
		_, v, _ := MandelbrotKernelFixnum(cr, ci, threshold2, MaxIter, BasePower2)
		if v != -1 {
			b.Fail()
		}
	}
}

func BenchmarkGmpFixedPoint(b *testing.B) {
	cr, err := mandelbrot_fixed_point.ParseFixnumSafe(centerRstr, BasePower2)
	g.Die(err)
	ci, err := mandelbrot_fixed_point.ParseFixnumSafe(centerIstr, BasePower2)
	g.Die(err)
	threshold2, err := mandelbrot_fixed_point.ParseFixnumSafe(Thresholdstr, BasePower2)
	threshold2.Mul(threshold2, threshold2)
	g.Die(err)
	//cr, ci = mandelbrot_fixed_point.MandelbrotCoordinate(500,500,1000,1000,cr,ci,gmp.NewInt(1 << 60), BasePower2)
	for n := 0; n < b.N; n++ {
		_, v, _ := mandelbrot_fixed_point.MandelbrotKernel(cr, ci, threshold2, MaxIter, BasePower2)
		if v != -1 {
			b.Fail()
		}
	}
}

// TODO make a whole thing out of this
func MandelbrotKernelFixnum(cr, ci, threshold2 *Fixnum, maxIter int64, basePower2 uint) (i int64, v float64, mag2 *Fixnum) {
	one_half, err := FromString("0.5")
	g.Die(err)
	zr := new(Fixnum)
	zi := new(Fixnum)
	zr2 := new(Fixnum)
	zi2 := new(Fixnum)
	magnitude2 := new(Fixnum)
	zri := new(Fixnum)
	for i := int64(0); i < maxIter; i++ {
		zr2.Mul(zr, zr)
		zi2.Mul(zi, zi)
		magnitude2.Add(zr2, zi2)
		if magnitude2.cmpWords(threshold2) > 0 {
			z2 := magnitude2.Float64()
			v := float64(i-1) - math.Log2(math.Log2(z2)-2.0) + 1
			if math.IsNaN(v) {
				v = float64(i - 1)
			}
			return i - 1, v, magnitude2
		}
		// otherwise re-use those values

		// calculate 2ab/B + ci. Instead of multiplying by 2, just shift right by one less
		// (calculate out of place to not disturb zi for the next calculation)
		zri.Mul(zr, zi)
		zri.Mul(zri, one_half).Add(zri, ci)

		// calculate (a^2 - b^2)/B + cr (in place this time)
		zr.Sub(zr2, zi2).Add(zr, cr)

		// swap zri and zi. zi is now the real value, and zri can maybe have its memory re-used next iteration
		zri, zi = zi, zri
	}
	// explicit not in set sentinel value
	return 0, -1.0, nil
}
