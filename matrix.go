package golang_raytracer

type Matrix4 [4][4]Float

func (m *Matrix4) TransformPoint(p *Vec3) Vec3 {
	return Vec3{
		m[0][0]*p.X + m[0][1]*p.Y + m[0][2]*p.Z + m[0][3],
		m[1][0]*p.X + m[1][1]*p.Y + m[1][2]*p.Z + m[1][3],
		m[2][0]*p.X + m[2][1]*p.Y + m[2][2]*p.Z + m[2][3],
	}
}

func (m *Matrix4) TransformVec(v *Vec3) Vec3 {
	return Vec3{
		m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z,
		m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z,
		m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z,
	}
}

func (m *Matrix4) TransformRay(t *Ray) Ray {
	return Ray{
		m.TransformPoint(&t.Origin),
		m.TransformVec(&t.Direction),
	}
}

type Matrix3 [3][3]Float

func (m *Matrix3) TransformPoint(p *Vec2) Vec2 {
	return Vec2{
		m[0][0]*p.X + m[0][1]*p.Y + m[0][2],
		m[1][0]*p.X + m[1][1]*p.Y + m[1][2],
	}
}

func (m *Matrix3) TransformVec(v *Vec2) Vec2 {
	return Vec2{
		m[0][0]*v.X + m[0][1]*v.Y,
		m[1][0]*v.X + m[1][1]*v.Y,
	}
}

func (m1 *Matrix3) Compose(m2 *Matrix3) (rvalue Matrix3) {
	// matrix multiplication
	for i := 0; i < 3; i += 1 {
		for j := 0; j < 3; j += 1 {
			sum := Float(0.0)
			for k := 0; k < 3; k += 1 {
				sum += m1[i][k] * m2[k][j]
			}
			rvalue[i][j] = sum;
		}
	}
	return rvalue;
}

func Scale2D(x, y, a Float) Matrix3 {
	return Matrix3{
		{a, 0, (1 - a) * x},
		{0, a, (1 - a) * y},
		{0, 0, 1},
	}
}

func Scale2Dv(v Vec2, a Float) Matrix3 {
	return Matrix3{
		{a, 0, (1 - a) * v.X},
		{0, a, (1 - a) * v.Y},
		{0, 0, 1},
	}
}

func Rotate2D(x, y, theta Float) Matrix3 {
	return Matrix3{
		{Cos(theta), -Sin(theta), x + y*Sin(theta) - x*Cos(theta)},
		{Sin(theta), Cos(theta), y - y*Cos(theta) - x*Sin(theta)},
		{0, 0, 1},
	}
}

func Rotate2Dv(v Vec2, theta Float) Matrix3 {
	return Matrix3{
		{Cos(theta), -Sin(theta), v.X + v.Y*Sin(theta) - v.X*Cos(theta)},
		{Sin(theta), Cos(theta), v.Y - v.Y*Cos(theta) - v.X*Sin(theta)},
		{0, 0, 1},
	}
}

func Translate2D(x, y Float) Matrix3 {
	return Matrix3{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	}
}

func Translate2Dv(v Vec2, ) Matrix3 {
	return Matrix3{
		{1, 0, v.X},
		{0, 1, v.Y},
		{0, 0, 1},
	}
}
