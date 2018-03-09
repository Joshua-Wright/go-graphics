package interpolation

import "math"

func h00(t float64) float64 { return 2.0*math.Pow(t, 3.0) - 3.0*math.Pow(t, 2.0) + 1; }
func h10(t float64) float64 { return math.Pow(t, 3.0) - 2.0*math.Pow(t, 2.0) + t; }
func h01(t float64) float64 { return -2.0*math.Pow(t, 3.0) + 3.0*math.Pow(t, 2.0); }
func h11(t float64) float64 { return math.Pow(t, 3.0) - math.Pow(t, 2.0); }

func hermite(x, x0, x1, p0, p1, m0, m1 float64) float64 {
	t := (x - x0) / (x1 - x0);
	return h00(t)*p0 +
		h10(t)*(x1-x0)*m0 +
		h01(t)*p1 +
		h11(t)*(x1-x0)*m1;
}

type Cubic struct {
	xs, ys, ms []float64
	ymin, ymax float64
}

func MakeCubicInterpolator(ymin, ymax float64, xs, ys []float64) *Cubic {
	var c Cubic
	c.ymin = ymin
	c.ymax = ymax
	c.xs = xs
	c.ys = ys
	// get slopes
	{ /*first element*/
		dx0 := xs[len(xs)-1] - xs[0];
		dy0 := ys[len(ys)-1] - ys[0];
		m0 := dy0 / dx0;
		dx1 := xs[0] - xs[1];
		dy1 := ys[0] - ys[1];
		m1 := dy1 / dx1;
		c.ms = append(c.ms, (m0+m1)/2);
	}
	for i := 1; i < len(xs)-1; i++ {
		dx0 := xs[i] - xs[i-1];
		dy0 := ys[i] - ys[i-1];
		m0 := dy0 / dx0;
		dx1 := xs[i+1] - xs[i];
		dy1 := ys[i+1] - ys[i];
		m1 := dy1 / dx1;
		c.ms = append(c.ms, (m0+m1)/2);
	}
	{ /*last element*/
		dx0 := xs[len(xs)-2] - xs[len(xs)-1];
		dy0 := ys[len(ys)-2] - ys[len(ys)-1];
		m0 := dy0 / dx0;
		dx1 := xs[0] - xs[len(xs)-1];
		dy1 := ys[0] - ys[len(ys)-1];
		m1 := dy1 / dx1;
		c.ms = append(c.ms, (m0+m1)/2);
	}

	// pre-clamp any obvious problems
	for i := 0; i < len(xs); i++ {
		if (ys[i] <= ymin || ys[i] >= ymax) {
			c.ms[i] = 0;
		}
	}
	return &c
}

func (c *Cubic) Get(x float64) float64 {

	i := 0;
	for c.xs[i] <= x {
		i++;
	}

	return c.clamp(hermite(x,
		c.xs[i-1], c.xs[i],
		c.ys[i-1], c.ys[i],
		c.ms[i-1], c.ms[i]));
}

func (c *Cubic) clamp(x float64) float64 {
	if (x < c.ymin) {
		return c.ymin;
	}
	if (x > c.ymax) {
		return c.ymax;
	}
	return x;
}
