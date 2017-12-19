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
	for i := 0; i < len(pts)-1; i++ {
		pts2 = append(pts2, pts[i].AddV(pts[i+1]).DivS(2))
	}
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
