package graphics

import "fmt"

//go:generate go run ../codegen/generate_vectors.go
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


func (v Vec2) String() string {
	return fmt.Sprint("[", v.X, ",", v.Y, "]")
}

//go:generate go run ../codegen/generate_vectors.go
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


func (v Vec3) String() string {
	return fmt.Sprint("[", v.X, ",", v.Y, ",", v.Z, "]")
}

//go:generate go run ../codegen/generate_vectors.go
type Vec4 struct {
	X Float
	Y Float
	Z Float
	W Float
}

func (v Vec4) Clone() *Vec4 {
	return &Vec4 {
		X: v.X,
		Y: v.Y,
		Z: v.Z,
		W: v.W,
	}
}

func (v Vec4) AddV(u Vec4) Vec4 {
	return Vec4 { 
		X: v.X + u.X,
		Y: v.Y + u.Y,
		Z: v.Z + u.Z,
		W: v.W + u.W,
	}
}
//func (v *Vec4) AddVi(u *Vec4) *Vec4 {
//	
//	v.X += u.X
//
//	v.Y += u.Y
//
//	v.Z += u.Z
//
//	v.W += u.W
//
//	return v
//}

func (v Vec4) SubV(u Vec4) Vec4 {
	return Vec4 { 
		X: v.X - u.X,
		Y: v.Y - u.Y,
		Z: v.Z - u.Z,
		W: v.W - u.W,
	}
}
//func (v *Vec4) SubVi(u *Vec4) *Vec4 {
//	
//	v.X -= u.X
//
//	v.Y -= u.Y
//
//	v.Z -= u.Z
//
//	v.W -= u.W
//
//	return v
//}

func (v Vec4) MulV(u Vec4) Vec4 {
	return Vec4 { 
		X: v.X * u.X,
		Y: v.Y * u.Y,
		Z: v.Z * u.Z,
		W: v.W * u.W,
	}
}
//func (v *Vec4) MulVi(u *Vec4) *Vec4 {
//	
//	v.X *= u.X
//
//	v.Y *= u.Y
//
//	v.Z *= u.Z
//
//	v.W *= u.W
//
//	return v
//}

func (v Vec4) DivV(u Vec4) Vec4 {
	return Vec4 { 
		X: v.X / u.X,
		Y: v.Y / u.Y,
		Z: v.Z / u.Z,
		W: v.W / u.W,
	}
}
//func (v *Vec4) DivVi(u *Vec4) *Vec4 {
//	
//	v.X /= u.X
//
//	v.Y /= u.Y
//
//	v.Z /= u.Z
//
//	v.W /= u.W
//
//	return v
//}

func (v Vec4) AddS(s Float) Vec4 {
	return Vec4 { 
		X: v.X + s,
		Y: v.Y + s,
		Z: v.Z + s,
		W: v.W + s,
	}
}
//func (v *Vec4) AddSi(s Float) *Vec4 {
//	
//	v.X += s
//
//	v.Y += s
//
//	v.Z += s
//
//	v.W += s
//
//	return v
//}

func (v Vec4) SubS(s Float) Vec4 {
	return Vec4 { 
		X: v.X - s,
		Y: v.Y - s,
		Z: v.Z - s,
		W: v.W - s,
	}
}
//func (v *Vec4) SubSi(s Float) *Vec4 {
//	
//	v.X -= s
//
//	v.Y -= s
//
//	v.Z -= s
//
//	v.W -= s
//
//	return v
//}

func (v Vec4) MulS(s Float) Vec4 {
	return Vec4 { 
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
		W: v.W * s,
	}
}
//func (v *Vec4) MulSi(s Float) *Vec4 {
//	
//	v.X *= s
//
//	v.Y *= s
//
//	v.Z *= s
//
//	v.W *= s
//
//	return v
//}

func (v Vec4) DivS(s Float) Vec4 {
	return Vec4 { 
		X: v.X / s,
		Y: v.Y / s,
		Z: v.Z / s,
		W: v.W / s,
	}
}
//func (v *Vec4) DivSi(s Float) *Vec4 {
//	
//	v.X /= s
//
//	v.Y /= s
//
//	v.Z /= s
//
//	v.W /= s
//
//	return v
//}

func (v Vec4) Mag2() Float {
	return  v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W
}
func (v Vec4) Mag() Float {
	return Sqrt( v.X*v.X+ v.Y*v.Y+ v.Z*v.Z+ v.W*v.W)
}
func (v Vec4) UnitV() Vec4 {
	l := Float(1.0)/Sqrt( v.X*v.X+ v.Y*v.Y+ v.Z*v.Z+ v.W*v.W)
	return Vec4 {
		X: v.X*l,
		Y: v.Y*l,
		Z: v.Z*l,
		W: v.W*l,
	}
}
//func (v *Vec4) UnitVi() *Vec4 {
//	l := Float(1.0)/Sqrt( v.X*v.X+ v.Y*v.Y+ v.Z*v.Z+ v.W*v.W)
//
//	v.X *= l
//
//	v.Y *= l
//
//	v.Z *= l
//
//	v.W *= l
//
//	return v
//}

func (v Vec4) Dot(u Vec4) Float {
	return  v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W
}

var Vec4Zero = Vec4 {
		X: 0,
		Y: 0,
		Z: 0,
		W: 0,
}
var Vec4One = Vec4 {
		X: 1,
		Y: 1,
		Z: 1,
		W: 1,
}



var Vec4X = Vec4 {
	X:  1, 
	Y:  0, 
	Z:  0, 
	W:  0, 
}

var Vec4Y = Vec4 {
	X:  0, 
	Y:  1, 
	Z:  0, 
	W:  0, 
}

var Vec4Z = Vec4 {
	X:  0, 
	Y:  0, 
	Z:  1, 
	W:  0, 
}

var Vec4W = Vec4 {
	X:  0, 
	Y:  0, 
	Z:  0, 
	W:  1, 
}


func (v Vec4) String() string {
	return fmt.Sprint("[", v.X, ",", v.Y, ",", v.Z, ",", v.W, "]")
}
