package mandelbrot

import (
	"github.com/joshua-wright/go-graphics/Slice2D"
	"math"
	"github.com/joshua-wright/go-graphics/parallel"
	"gopkg.in/cheggaaa/pb.v1"
)

func MandelbrotPolynomial(z, c complex128) (z2 complex128) {
	z2 = z*z + c
	return z2
}

func IterateMandelbrot(z, c complex128, threshold float64, maxIter int) float64 {
	threshold2 := threshold*threshold
	for i := 0; i < maxIter; i++ {
		z = MandelbrotPolynomial(z, c)
		//if cmplx.Abs(z) > threshold {
		if real(z)*real(z)+imag(z)*imag(z) >= threshold2 {
			// smooth code from wikipedia
			// sqrt of inner term removed using log simplification rules.
			log_zn := math.Log(real(z)*real(z)+imag(z)*imag(z)) / 2
			nu := math.Log(log_zn/math.Log(2)) / math.Log(2)
			// Rearranging the potential function.
			// Dividing log_zn by log(2) instead of log(N = 1<<8)
			// because we want the entire palette to range from the
			// center to radius 2, NOT our bailout radius.
			iteration := float64(i) + 1 - nu
			return iteration
		}
	}
	return 0.0
}

func Mandelbrot(bounds [4]float64, width, height, maxIter int) Slice2D.Float64Slice2D {
	topLeft := complex(bounds[0], bounds[2])
	dr := (bounds[1] - bounds[0]) / float64(width)
	di := (bounds[3] - bounds[2]) / float64(height)

	out := Slice2D.NewFloat64Slice2D(width, height)

	println("iterate points")
	bar := pb.StartNew(width)
	parallel.ParallelFor(0, width, func(i int) {
		for j := 0; j < height; j++ {
			// translate from bounds to index
			z := complex(0, 0)
			c := topLeft + complex(dr*float64(i), di*float64(j))
			out.Set(i, j, IterateMandelbrot(z, c, 4.0, maxIter))
		}
		bar.Increment()
	})
	bar.Finish()

	return out
}
