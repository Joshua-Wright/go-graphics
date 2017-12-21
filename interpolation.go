package golang_raytracer

func BSplineSubdivide(pts []Vec2) []Vec2 {
	pts2 := make([]Vec2, 0, 2*len(pts))
	for i, p := range pts {
		if i == 0 {
			// no previous point for midpoint calculation
			pts2 = append(pts2, p)
		} else {
			// append midpoint and next point
			pts2 = append(pts2, p.AddV(pts[i-1]).DivS(2))
			pts2 = append(pts2, p)
		}
	}
	return pts2
}

func BSplineAverage(pts []Vec2) []Vec2 {
	pts2 := make([]Vec2, 0, 2*len(pts))
	pts2 = append(pts2, pts[0])
	for i := 0; i < len(pts)-1; i++ {
		pts2 = append(pts2, pts[i].AddV(pts[i+1]).DivS(2))
	}
	pts2 = append(pts2, pts[len(pts)-1])
	return pts2
}

// intentionally does not loop the input
func BSpline(pts []Vec2, iterations int, smoothness int) []Vec2 {
	pts2 := pts
	for i := 0; i < iterations; i++ {
		pts2 = BSplineSubdivide(pts2)
		for j := 0; j < smoothness; j++ {
			pts2 = BSplineAverage(pts2)
		}
	}
	return pts2
}

// interpolates until the maximum distance between two adjacent points is less than threshold
func BSplineAdaptive(pts []Vec2, smoothness int, threshold Float) []Vec2 {
	for {
		pts = BSpline(pts, 1, smoothness)
		if MaxAdjacentDistance(pts) < threshold {
			return pts
		}
	}
}

func Lagrange(pts []float64, i uint64, t float64) float64 {
	n := len(pts) - 2
	if n == 0 {
		// base case: linearly interpolate two points
		return pts[0]*(float64(i)+1-t) + pts[1]*(t-float64(i))
	}
	f := Lagrange(pts[:len(pts)-1], i, t)
	g := Lagrange(pts[1:], i+1, t)
	return ((float64(n)+1+float64(i)-t)*f + (t-float64(i))*g) / (float64(n) + 1)
}