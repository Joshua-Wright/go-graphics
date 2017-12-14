package golang_raytracer

type BBox struct {
	Min Vec3
	Max Vec3
}

func (b *BBox) Diag() Vec3 {
	return b.Max.SubV(b.Min)
}


func (b *BBox) union_point(v Vec3) {
	// TODO
	panic("TODO")
}
