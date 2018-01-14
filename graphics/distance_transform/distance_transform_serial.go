package distance_transform

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
)

//go:generate go run ../../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE pixelSerial
type pixelSerial struct {
	minPos   g.Vec2i
	minDist2 int // integer distance only because on grid and distance squared
}

type pixelUpdateSerial struct {
	zeroPoint g.Vec2i
	from      g.Vec2i
	to        g.Vec2i
}

func DistanceTransform(width, height int, zero_points []g.Vec2i) [][]float64 {
	//mesh := Make2DSlicePixelSerial(width, height, pixelSerial{g.Vec2i{-1, -1}, math.MaxInt64})
	mesh := NewPixelSerialSlice2D(width, height)
	for i := 0; i < len(mesh.Data); i++ {
		mesh.Data[i] = pixelSerial{g.Vec2i{-1, -1}, math.MaxInt64}
	}

	output_mesh := Make2DSliceFloat64(width, height, 0.0)

	msgBuf := []pixelUpdateSerial{}
	// send initial control point messages
	for _, p := range zero_points {
		msgBuf = append(msgBuf, pixelUpdateSerial{p, p, p})
	}

	for {
		newMsgBuf := []pixelUpdateSerial{}
		for _, u := range msgBuf {
			// compute new distance
			dx := u.to.X - u.zeroPoint.X
			dy := u.to.Y - u.zeroPoint.Y
			new_dist2 := dx*dx + dy*dy

			if new_dist2 < mesh.At(u.to.X, u.to.Y).minDist2 {
				// update closest pixel
				mesh.At(u.to.X, u.to.Y).minPos.X = u.zeroPoint.X
				mesh.At(u.to.X, u.to.Y).minPos.Y = u.zeroPoint.Y
				mesh.At(u.to.X, u.to.Y).minDist2 = new_dist2

				// update neighbors
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						if i == j {
							// don't send to ourself
							continue
						}
						x := u.to.X + i
						y := u.to.Y + j
						newMessage := pixelUpdateSerial{
							zeroPoint: u.zeroPoint,
							from:      u.to,
							to:        g.Vec2i{x, y},
						}
						if x < 0 || x >= width || y < 0 || y >= height {
							// don't send out of bounds
							continue
						}
						if newMessage.to == u.to {
							// don't send message back
							continue
						}
						newMsgBuf = append(newMsgBuf, newMessage)
					}
				}
			}
		}
		if len(newMsgBuf) == 0 {
			break
		} else {
			msgBuf = newMsgBuf
		}
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			//output_mesh[x][y] = math.Sqrt(float64(mesh[x][y].minDist2))
			output_mesh[x][y] = math.Sqrt(float64(mesh.At(x, y).minDist2))
		}
	}

	return output_mesh
}
