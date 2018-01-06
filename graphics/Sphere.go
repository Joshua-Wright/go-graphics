package graphics

type Sphere struct {
	Center Vec3
	Radius Float
}

func (s *Sphere) BoundingBox() *BBox {
	scaled_normal := Vec3{1, 1, 1}
	return &BBox{
		Min: s.Center.SubV(scaled_normal),
		Max: s.Center.AddV(scaled_normal),
	}
}

func (s *Sphere) Intersect(r *Ray) *IntersectionResult {
	i := r.Origin.SubV(s.Center)
	a := Float(1.0)
	b := 2.0 * r.Direction.Dot(i)
	c := i.Mag2() - s.Radius*s.Radius
	discriminant := b*b - 4.0*a*c

	if discriminant <= 0 {
		return nil
	} else {
		disc_sqrt := Sqrt(float64(discriminant))
		t1 := (-b + disc_sqrt) / 2.0 * a
		t2 := (-b - disc_sqrt) / 2.0 * a

		switch {
		case t1 >= 0 && t2 < 0:
			return &IntersectionResult{
				t1,
				r.PointAt(t1).SubV(s.Center).UnitV(),
			}
		case t2 >= 0 && t1 < 0:
			return &IntersectionResult{
				t2,
				r.PointAt(t2).SubV(s.Center).UnitV(),
			}
		case t1 <= t2:
			return &IntersectionResult{
				t1,
				r.PointAt(t1).SubV(s.Center).UnitV(),
			}
		case t2 <= t1:
			return &IntersectionResult{
				t2,
				r.PointAt(t2).SubV(s.Center).UnitV(),
			}
		}
		panic("unreachable")
	}
}
