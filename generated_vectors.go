package golang_raytracer

//go:generate go run codegen/generate_vectors.go
type Vec2 struct {
	X Float
	Y Float
}

func (v Vec2) Clone() *Vec2 {
	return &Vec2 {
		X: v.X,
		Y: v.Y,
	}
}

func (v Vec2) AddV(u Vec2) Vec2 {
	return Vec2 { 
		X: v.X + u.X,
		Y: v.Y + u.Y,
	}
}
//func (v *Vec2) AddVi(u *Vec2) *Vec2 {
//	
//	v.X += u.X
//
//	v.Y += u.Y
//
//	return v
//}

func (v Vec2) SubV(u Vec2) Vec2 {
	return Vec2 { 
		X: v.X - u.X,
		Y: v.Y - u.Y,
	}
}
//func (v *Vec2) SubVi(u *Vec2) *Vec2 {
//	
//	v.X -= u.X
//
//	v.Y -= u.Y
//
//	return v
//}

func (v Vec2) MulV(u Vec2) Vec2 {
	return Vec2 { 
		X: v.X * u.X,
		Y: v.Y * u.Y,
	}
}
//func (v *Vec2) MulVi(u *Vec2) *Vec2 {
//	
//	v.X *= u.X
//
//	v.Y *= u.Y
//
//	return v
//}

func (v Vec2) DivV(u Vec2) Vec2 {
	return Vec2 { 
		X: v.X / u.X,
		Y: v.Y / u.Y,
	}
}
//func (v *Vec2) DivVi(u *Vec2) *Vec2 {
//	
//	v.X /= u.X
//
//	v.Y /= u.Y
//
//	return v
//}

func (v Vec2) AddS(s Float) Vec2 {
	return Vec2 { 
		X: v.X + s,
		Y: v.Y + s,
	}
}
//func (v *Vec2) AddSi(s Float) *Vec2 {
//	
//	v.X += s
//
//	v.Y += s
//
//	return v
//}

func (v Vec2) SubS(s Float) Vec2 {
	return Vec2 { 
		X: v.X - s,
		Y: v.Y - s,
	}
}
//func (v *Vec2) SubSi(s Float) *Vec2 {
//	
//	v.X -= s
//
//	v.Y -= s
//
//	return v
//}

func (v Vec2) MulS(s Float) Vec2 {
	return Vec2 { 
		X: v.X * s,
		Y: v.Y * s,
	}
}
//func (v *Vec2) MulSi(s Float) *Vec2 {
//	
//	v.X *= s
//
//	v.Y *= s
//
//	return v
//}

func (v Vec2) DivS(s Float) Vec2 {
	return Vec2 { 
		X: v.X / s,
		Y: v.Y / s,
	}
}
//func (v *Vec2) DivSi(s Float) *Vec2 {
//	
//	v.X /= s
//
//	v.Y /= s
//
//	return v
//}

func (v Vec2) Mag2() Float {
	return  v.X*v.X + v.Y*v.Y
}
func (v Vec2) Mag() Float {
	return Sqrt( v.X*v.X+ v.Y*v.Y)
}
func (v Vec2) UnitV() Vec2 {
	l := Float(1.0)/Sqrt( v.X*v.X+ v.Y*v.Y)
	return Vec2 {
		X: v.X*l,
		Y: v.Y*l,
	}
}
//func (v *Vec2) UnitVi() *Vec2 {
//	l := Float(1.0)/Sqrt( v.X*v.X+ v.Y*v.Y)
//
//	v.X *= l
//
//	v.Y *= l
//
//	return v
//}

func (v Vec2) Dot(u Vec2) Float {
	return  v.X*v.X + v.Y*v.Y
}

var Vec2Zero = Vec2 {
		X: 0,
		Y: 0,
}
var Vec2One = Vec2 {
		X: 1,
		Y: 1,
}



var Vec2X = Vec2 {
	X:  1, 
	Y:  0, 
}

var Vec2Y = Vec2 {
	X:  0, 
	Y:  1, 
}


//go:generate go run codegen/generate_vectors.go
type Vec3 struct {
	X Float
	Y Float
	Z Float
}

func (v Vec3) Clone() *Vec3 {
	return &Vec3 {
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
}

func (v Vec3) AddV(u Vec3) Vec3 {
	return Vec3 { 
		X: v.X + u.X,
		Y: v.Y + u.Y,
		Z: v.Z + u.Z,
	}
}
//func (v *Vec3) AddVi(u *Vec3) *Vec3 {
//	
//	v.X += u.X
//
//	v.Y += u.Y
//
//	v.Z += u.Z
//
//	return v
//}

func (v Vec3) SubV(u Vec3) Vec3 {
	return Vec3 { 
		X: v.X - u.X,
		Y: v.Y - u.Y,
		Z: v.Z - u.Z,
	}
}
//func (v *Vec3) SubVi(u *Vec3) *Vec3 {
//	
//	v.X -= u.X
//
//	v.Y -= u.Y
//
//	v.Z -= u.Z
//
//	return v
//}

func (v Vec3) MulV(u Vec3) Vec3 {
	return Vec3 { 
		X: v.X * u.X,
		Y: v.Y * u.Y,
		Z: v.Z * u.Z,
	}
}
//func (v *Vec3) MulVi(u *Vec3) *Vec3 {
//	
//	v.X *= u.X
//
//	v.Y *= u.Y
//
//	v.Z *= u.Z
//
//	return v
//}

func (v Vec3) DivV(u Vec3) Vec3 {
	return Vec3 { 
		X: v.X / u.X,
		Y: v.Y / u.Y,
		Z: v.Z / u.Z,
	}
}
//func (v *Vec3) DivVi(u *Vec3) *Vec3 {
//	
//	v.X /= u.X
//
//	v.Y /= u.Y
//
//	v.Z /= u.Z
//
//	return v
//}

func (v Vec3) AddS(s Float) Vec3 {
	return Vec3 { 
		X: v.X + s,
		Y: v.Y + s,
		Z: v.Z + s,
	}
}
//func (v *Vec3) AddSi(s Float) *Vec3 {
//	
//	v.X += s
//
//	v.Y += s
//
//	v.Z += s
//
//	return v
//}

func (v Vec3) SubS(s Float) Vec3 {
	return Vec3 { 
		X: v.X - s,
		Y: v.Y - s,
		Z: v.Z - s,
	}
}
//func (v *Vec3) SubSi(s Float) *Vec3 {
//	
//	v.X -= s
//
//	v.Y -= s
//
//	v.Z -= s
//
//	return v
//}

func (v Vec3) MulS(s Float) Vec3 {
	return Vec3 { 
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}
//func (v *Vec3) MulSi(s Float) *Vec3 {
//	
//	v.X *= s
//
//	v.Y *= s
//
//	v.Z *= s
//
//	return v
//}

func (v Vec3) DivS(s Float) Vec3 {
	return Vec3 { 
		X: v.X / s,
		Y: v.Y / s,
		Z: v.Z / s,
	}
}
//func (v *Vec3) DivSi(s Float) *Vec3 {
//	
//	v.X /= s
//
//	v.Y /= s
//
//	v.Z /= s
//
//	return v
//}

func (v Vec3) Mag2() Float {
	return  v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
func (v Vec3) Mag() Float {
	return Sqrt( v.X*v.X+ v.Y*v.Y+ v.Z*v.Z)
}
func (v Vec3) UnitV() Vec3 {
	l := Float(1.0)/Sqrt( v.X*v.X+ v.Y*v.Y+ v.Z*v.Z)
	return Vec3 {
		X: v.X*l,
		Y: v.Y*l,
		Z: v.Z*l,
	}
}
//func (v *Vec3) UnitVi() *Vec3 {
//	l := Float(1.0)/Sqrt( v.X*v.X+ v.Y*v.Y+ v.Z*v.Z)
//
//	v.X *= l
//
//	v.Y *= l
//
//	v.Z *= l
//
//	return v
//}

func (v Vec3) Dot(u Vec3) Float {
	return  v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

var Vec3Zero = Vec3 {
		X: 0,
		Y: 0,
		Z: 0,
}
var Vec3One = Vec3 {
		X: 1,
		Y: 1,
		Z: 1,
}



var Vec3X = Vec3 {
	X:  1, 
	Y:  0, 
	Z:  0, 
}

var Vec3Y = Vec3 {
	X:  0, 
	Y:  1, 
	Z:  0, 
}

var Vec3Z = Vec3 {
	X:  0, 
	Y:  0, 
	Z:  1, 
}

