package mandelbrot

import (
	"github.com/joshua-wright/go-graphics/Slice2D"
	"math"
	"github.com/joshua-wright/go-graphics/parallel"
	"gopkg.in/cheggaaa/pb.v1"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"image/color"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
)

func MandelbrotPolynomial(z, c complex128) (z2 complex128) {
	z2 = z*z + c
	return z2
}

func IterateMandelbrot(z, c complex128, threshold float64, maxIter int64) float64 {
	threshold2 := threshold * threshold
	for i := int64(0); i < maxIter; i++ {
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

func Mandelbrot(bounds [4]float64, width, height int, maxIter int64) Slice2D.Float64Slice2D {
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

func MandelbrotBounds(width, height int64, center complex128, zoom float64) (topLeft complex128, dr, di float64) {
	dx := 2.0 / zoom;
	dy := 2.0 / zoom;
	if width > height {
		/* widescreen image */
		dx = float64(width) / float64(height) * dy;
	} else if (height > width) {
		/* portrait */
		dy = float64(height) / float64(width) * dx;
	} // otherwise square
	bounds := [4]float64{
		real(center) - dx, real(center) + dx, imag(center) - dy, imag(center) + dy,
	};
	topLeft = complex(bounds[0], bounds[3])
	dr = (bounds[1] - bounds[0]) / float64(width)
	di = (bounds[3] - bounds[2]) / float64(height)
	return topLeft, dr, di
}

type MandelbrotPerPixel struct {
	TopLeft       complex128
	Dr, Di        float64
	Width, Height int64
	MaxIter       int64
	MaxVal        float64
	Wrap          float64
	Cmap          colormap.ColorMap
	OutImage      *memory_mapped.PPMFile
}

func (m *MandelbrotPerPixel) GetPixel(i, j int64) {
	z := complex(0, 0)
	c := m.TopLeft + complex(m.Dr*float64(i), -m.Di*float64(j))
	val := IterateMandelbrot(z, c, 1000, m.MaxIter)

	if val == 0.0 {
		m.OutImage.Set64(i, j, color.RGBA{0, 0, 0, 255})
		return
	} else {
		val = math.Log2(val+1) / math.Log2(m.MaxVal+1) * m.Wrap
		val = math.Sin(val*2*math.Pi)/2.0 + 0.5

		m.OutImage.Set64(i, j, m.Cmap.GetColor(val))
	}

}

func (m *MandelbrotPerPixel) Bounds() (w, h int64) { return m.Width, m.Height }
