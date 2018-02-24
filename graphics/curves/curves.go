package curves

import (
	"math"
	g "github.com/joshua-wright/go-graphics/graphics"
)

func Hypotrochoid(R, r, d, t float64) g.Vec2 {
	// https://en.wikipedia.org/wiki/Hypotrochoid
	return g.Vec2{
		X: (R-r)*math.Cos(t) + d*math.Cos(t*(R-r)/r),
		Y: (R-r)*math.Sin(t) - d*math.Sin(t*(R-r)/r),
	}
}

func Clifford(a, b, c, d float64, v g.Vec2) g.Vec2 {
	return g.Vec2{
		X: math.Sin(a*v.Y) + c*math.Cos(a*v.X),
		Y: math.Sin(b*v.X) + d*math.Cos(b*v.Y),
	}
}
