package graphics

func (v Vec2) Cross(u Vec2) Float {
	return v.X*u.Y - v.Y*u.X
}
