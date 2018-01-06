package graphics

type Surface interface {
	// may return null if no bounding box (e.g. planes)
	BoundingBox() *BBox
	Intersect(*Ray) *IntersectionResult
}

type IntersectionResult struct {
	TParameter Float
	Normal     Vec3
}
