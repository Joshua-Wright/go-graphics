package golang_raytracer

import "math"

//type Float = float32
type Float = float64

func Sqrt(f Float) Float {
	return Float(math.Sqrt(float64(f)))
}

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r *Ray) PointAt(t Float) Vec3 {
	v := r.Direction.MulS(t)
	return v.AddV(r.Origin)
}
