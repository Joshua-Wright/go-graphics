package graphics

type Octree struct {
	Surfaces  []Surface
	Infinites []Surface
	Root      *OctreeNode
}

type OctreeNode struct {
	BBox     BBox
	Depth    int32
	Children []OctreeNode
	LeafData []Surface
}

func (n *OctreeNode) Subdivide() {
	// TODO
	diag2 := n.BBox.Diag().MulS(1.0 / 2.0)
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {

				p1 := Vec3{Float(x), Float(y), Float(z),}
				p2 := Vec3{Float(1 - x), Float(1 - y), Float(1 - z),}

				child_bbox := BBox{
					Min: p1.AddV(n.BBox.Min).MulV(diag2),
					Max: p2.AddV(n.BBox.Max).MulV(diag2),
				}

				n.Children = append(n.Children, OctreeNode{
					BBox:     child_bbox,
					Depth:    n.Depth - 1,
					Children: nil,
				})

			}
		}
	}
}

func (n *OctreeNode) IsLeaf() bool { return len(n.Children) == 0 }

func (n *OctreeNode) Insert(surface Surface) {
	// max depth
	if n.Depth <= 0 {
		n.LeafData = append(n.LeafData, surface)
		return
	}

	// empty leaf node
	if n.IsLeaf() && len(n.LeafData) == 0 {
		n.LeafData = []Surface{surface}
		return
	}

	// Occupied leaf node and not max depth: subdivide node

	// Interior node (has children)

}
