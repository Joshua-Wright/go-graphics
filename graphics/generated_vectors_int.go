package graphics

// TODO generate this file instead of modifying generated_vectors.go

import "fmt"


type Vec2i struct {
	X int
	Y int
}

func (v Vec2i) Clone() *Vec2i {
	return &Vec2i {
		X: v.X,
		Y: v.Y,
	}
}

func (v Vec2i) AddV(u Vec2i) Vec2i {
	return Vec2i { 
		X: v.X + u.X,
		Y: v.Y + u.Y,
	}
}
//func (v *Vec2i) AddVi(u *Vec2i) *Vec2i {
//	
//	v.X += u.X
//
//	v.Y += u.Y
//
//	return v
//}

func (v Vec2i) SubV(u Vec2i) Vec2i {
	return Vec2i { 
		X: v.X - u.X,
		Y: v.Y - u.Y,
	}
}
//func (v *Vec2i) SubVi(u *Vec2i) *Vec2i {
//	
//	v.X -= u.X
//
//	v.Y -= u.Y
//
//	return v
//}

func (v Vec2i) MulV(u Vec2i) Vec2i {
	return Vec2i { 
		X: v.X * u.X,
		Y: v.Y * u.Y,
	}
}
//func (v *Vec2i) MulVi(u *Vec2i) *Vec2i {
//	
//	v.X *= u.X
//
//	v.Y *= u.Y
//
//	return v
//}

func (v Vec2i) DivV(u Vec2i) Vec2i {
	return Vec2i { 
		X: v.X / u.X,
		Y: v.Y / u.Y,
	}
}
//func (v *Vec2i) DivVi(u *Vec2i) *Vec2i {
//	
//	v.X /= u.X
//
//	v.Y /= u.Y
//
//	return v
//}

func (v Vec2i) AddS(s int) Vec2i {
	return Vec2i { 
		X: v.X + s,
		Y: v.Y + s,
	}
}
//func (v *Vec2i) AddSi(s int) *Vec2i {
//	
//	v.X += s
//
//	v.Y += s
//
//	return v
//}

func (v Vec2i) SubS(s int) Vec2i {
	return Vec2i { 
		X: v.X - s,
		Y: v.Y - s,
	}
}
//func (v *Vec2i) SubSi(s int) *Vec2i {
//	
//	v.X -= s
//
//	v.Y -= s
//
//	return v
//}

func (v Vec2i) MulS(s int) Vec2i {
	return Vec2i { 
		X: v.X * s,
		Y: v.Y * s,
	}
}
//func (v *Vec2i) MulSi(s int) *Vec2i {
//	
//	v.X *= s
//
//	v.Y *= s
//
//	return v
//}

func (v Vec2i) DivS(s int) Vec2i {
	return Vec2i { 
		X: v.X / s,
		Y: v.Y / s,
	}
}
//func (v *Vec2i) DivSi(s int) *Vec2i {
//	
//	v.X /= s
//
//	v.Y /= s
//
//	return v
//}

func (v Vec2i) Dot(u Vec2i) int {
	return  v.X*v.X + v.Y*v.Y
}

var Vec2iZero = Vec2i {
		X: 0,
		Y: 0,
}
var Vec2iOne = Vec2i {
		X: 1,
		Y: 1,
}



var Vec2iX = Vec2i {
	X:  1, 
	Y:  0, 
}

var Vec2iY = Vec2i {
	X:  0, 
	Y:  1, 
}


func (v Vec2i) String() string {
	return fmt.Sprint("[", v.X, ",", v.Y, "]")
}

type Vec3i struct {
	X int
	Y int
	Z int
}

func (v Vec3i) Clone() *Vec3i {
	return &Vec3i {
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
}

func (v Vec3i) AddV(u Vec3i) Vec3i {
	return Vec3i { 
		X: v.X + u.X,
		Y: v.Y + u.Y,
		Z: v.Z + u.Z,
	}
}
//func (v *Vec3i) AddVi(u *Vec3i) *Vec3i {
//	
//	v.X += u.X
//
//	v.Y += u.Y
//
//	v.Z += u.Z
//
//	return v
//}

func (v Vec3i) SubV(u Vec3i) Vec3i {
	return Vec3i { 
		X: v.X - u.X,
		Y: v.Y - u.Y,
		Z: v.Z - u.Z,
	}
}
//func (v *Vec3i) SubVi(u *Vec3i) *Vec3i {
//	
//	v.X -= u.X
//
//	v.Y -= u.Y
//
//	v.Z -= u.Z
//
//	return v
//}

func (v Vec3i) MulV(u Vec3i) Vec3i {
	return Vec3i { 
		X: v.X * u.X,
		Y: v.Y * u.Y,
		Z: v.Z * u.Z,
	}
}
//func (v *Vec3i) MulVi(u *Vec3i) *Vec3i {
//	
//	v.X *= u.X
//
//	v.Y *= u.Y
//
//	v.Z *= u.Z
//
//	return v
//}

func (v Vec3i) DivV(u Vec3i) Vec3i {
	return Vec3i { 
		X: v.X / u.X,
		Y: v.Y / u.Y,
		Z: v.Z / u.Z,
	}
}
//func (v *Vec3i) DivVi(u *Vec3i) *Vec3i {
//	
//	v.X /= u.X
//
//	v.Y /= u.Y
//
//	v.Z /= u.Z
//
//	return v
//}

func (v Vec3i) AddS(s int) Vec3i {
	return Vec3i { 
		X: v.X + s,
		Y: v.Y + s,
		Z: v.Z + s,
	}
}
//func (v *Vec3i) AddSi(s int) *Vec3i {
//	
//	v.X += s
//
//	v.Y += s
//
//	v.Z += s
//
//	return v
//}

func (v Vec3i) SubS(s int) Vec3i {
	return Vec3i { 
		X: v.X - s,
		Y: v.Y - s,
		Z: v.Z - s,
	}
}
//func (v *Vec3i) SubSi(s int) *Vec3i {
//	
//	v.X -= s
//
//	v.Y -= s
//
//	v.Z -= s
//
//	return v
//}

func (v Vec3i) MulS(s int) Vec3i {
	return Vec3i { 
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}
//func (v *Vec3i) MulSi(s int) *Vec3i {
//	
//	v.X *= s
//
//	v.Y *= s
//
//	v.Z *= s
//
//	return v
//}

func (v Vec3i) DivS(s int) Vec3i {
	return Vec3i { 
		X: v.X / s,
		Y: v.Y / s,
		Z: v.Z / s,
	}
}
//func (v *Vec3i) DivSi(s int) *Vec3i {
//	
//	v.X /= s
//
//	v.Y /= s
//
//	v.Z /= s
//
//	return v
//}

func (v Vec3i) Mag2() int {
	return  v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3i) Dot(u Vec3i) int {
	return  v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

var Vec3iZero = Vec3i {
		X: 0,
		Y: 0,
		Z: 0,
}
var Vec3iOne = Vec3i {
		X: 1,
		Y: 1,
		Z: 1,
}



var Vec3iX = Vec3i {
	X:  1, 
	Y:  0, 
	Z:  0, 
}

var Vec3iY = Vec3i {
	X:  0, 
	Y:  1, 
	Z:  0, 
}

var Vec3iZ = Vec3i {
	X:  0, 
	Y:  0, 
	Z:  1, 
}


func (v Vec3i) String() string {
	return fmt.Sprint("[", v.X, ",", v.Y, ",", v.Z, "]")
}


type Vec4i struct {
	X int
	Y int
	Z int
	W int
}

func (v Vec4i) Clone() *Vec4i {
	return &Vec4i {
		X: v.X,
		Y: v.Y,
		Z: v.Z,
		W: v.W,
	}
}

func (v Vec4i) AddV(u Vec4i) Vec4i {
	return Vec4i { 
		X: v.X + u.X,
		Y: v.Y + u.Y,
		Z: v.Z + u.Z,
		W: v.W + u.W,
	}
}
//func (v *Vec4i) AddVi(u *Vec4i) *Vec4i {
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

func (v Vec4i) SubV(u Vec4i) Vec4i {
	return Vec4i { 
		X: v.X - u.X,
		Y: v.Y - u.Y,
		Z: v.Z - u.Z,
		W: v.W - u.W,
	}
}
//func (v *Vec4i) SubVi(u *Vec4i) *Vec4i {
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

func (v Vec4i) MulV(u Vec4i) Vec4i {
	return Vec4i { 
		X: v.X * u.X,
		Y: v.Y * u.Y,
		Z: v.Z * u.Z,
		W: v.W * u.W,
	}
}
//func (v *Vec4i) MulVi(u *Vec4i) *Vec4i {
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

func (v Vec4i) DivV(u Vec4i) Vec4i {
	return Vec4i { 
		X: v.X / u.X,
		Y: v.Y / u.Y,
		Z: v.Z / u.Z,
		W: v.W / u.W,
	}
}
//func (v *Vec4i) DivVi(u *Vec4i) *Vec4i {
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

func (v Vec4i) AddS(s int) Vec4i {
	return Vec4i { 
		X: v.X + s,
		Y: v.Y + s,
		Z: v.Z + s,
		W: v.W + s,
	}
}
//func (v *Vec4i) AddSi(s int) *Vec4i {
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

func (v Vec4i) SubS(s int) Vec4i {
	return Vec4i { 
		X: v.X - s,
		Y: v.Y - s,
		Z: v.Z - s,
		W: v.W - s,
	}
}
//func (v *Vec4i) SubSi(s int) *Vec4i {
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

func (v Vec4i) MulS(s int) Vec4i {
	return Vec4i { 
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
		W: v.W * s,
	}
}
//func (v *Vec4i) MulSi(s int) *Vec4i {
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

func (v Vec4i) DivS(s int) Vec4i {
	return Vec4i { 
		X: v.X / s,
		Y: v.Y / s,
		Z: v.Z / s,
		W: v.W / s,
	}
}
//func (v *Vec4i) DivSi(s int) *Vec4i {
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

func (v Vec4i) Mag2() int {
	return  v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W
}

func (v Vec4i) Dot(u Vec4i) int {
	return  v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W
}

var Vec4iZero = Vec4i {
		X: 0,
		Y: 0,
		Z: 0,
		W: 0,
}
var Vec4iOne = Vec4i {
		X: 1,
		Y: 1,
		Z: 1,
		W: 1,
}



var Vec4iX = Vec4i {
	X:  1, 
	Y:  0, 
	Z:  0, 
	W:  0, 
}

var Vec4iY = Vec4i {
	X:  0, 
	Y:  1, 
	Z:  0, 
	W:  0, 
}

var Vec4iZ = Vec4i {
	X:  0, 
	Y:  0, 
	Z:  1, 
	W:  0, 
}

var Vec4iW = Vec4i {
	X:  0, 
	Y:  0, 
	Z:  0, 
	W:  1, 
}


func (v Vec4i) String() string {
	return fmt.Sprint("[", v.X, ",", v.Y, ",", v.Z, ",", v.W, "]")
}
