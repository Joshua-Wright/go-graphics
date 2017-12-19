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

func (m1 Matrix3) Compose(m2 Matrix3) (rvalue Matrix3) {
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

func MatrixImage2D(p, q, r, phat, qhat, rhat Vec2) Matrix3 {
	rvalue := Matrix3{}

	// efficient matrix inversion code (math from wikipedia)
	a := p.X
	b := q.X
	c := r.X
	d := p.Y
	e := q.Y
	f := r.Y
	g := 1.0
	h := 1.0
	i := 1.0
	A := e*i - f*h
	D := -(b*i - c*h)
	G := b*f - c*e
	B := -(d*i - f*g)
	E := a*i - c*g
	H := -(a*f - c*d)
	C := d*h - e*g
	F := -(a*h - b*g)
	I := a*e - b*d
	det := a*A + b*B + c*C
	det_inv := 1.0 / det

	A_inv := Matrix3{
		{A, D, G},
		{B, E, H},
		{C, F, I},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			A_inv[i][j] *= det_inv;
		}
	}

	// matrix of output points
	out_pts := Matrix3{
		{phat.X, qhat.X, rhat.X},
		{phat.Y, qhat.Y, rhat.Y},
		{1, 1, 1},
	};

	// now multiply output point matrix by inverse of input matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum := 0.0;
			for k := 0; k < 3; k++ {
				sum += out_pts[i][k] * A_inv[k][j];
			}
			rvalue[i][j] = sum;
		}
	}

	return rvalue;
}
