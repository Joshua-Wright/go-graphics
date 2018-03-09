package interpolation

import "math"

type CubicEquidistant struct {
	Cubic
}

func MakeCubicEquidistantInterpolator(ymin, ymax float64, nPts int, f func(x float64) float64) *CubicEquidistant {
	xs := make([]float64, nPts)
	ys := make([]float64, nPts)
	for i := 0; i < nPts; i++ {
		x := float64(i)/float64(nPts)*(ymax-ymin) + ymin
		xs[i] = x
		ys[i] = x
	}

	var c CubicEquidistant
	c2 := MakeCubicInterpolator(ymin, ymax, xs, ys)
	c.ymin = c2.ymin
	c.ymax = c2.ymax
	c.ys = c2.ys
	c.xs = c2.xs
	c.ms = c2.ms

	return &c
}

func (c *CubicEquidistant) Get(x float64) float64 {
	// FIXME: this probably doesn't work when the range doesn't start at 0
	i := int(math.Mod(x, c.ymax -c.ymin) / (c.ymax - c.ymin) * float64(len(c.xs)));
	return c.clamp(hermite(x,
		c.xs[i-1], c.xs[i],
		c.ys[i-1], c.ys[i],
		c.ms[i-1], c.ms[i]));
}