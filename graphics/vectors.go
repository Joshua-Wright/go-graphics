package graphics

//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec2
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec2i
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec3
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec3i
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec4
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec4i

func (v Vec2) Cross(u Vec2) Float {
	return v.X*u.Y - v.Y*u.X
}