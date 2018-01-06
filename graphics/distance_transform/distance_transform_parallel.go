package distance_transform

import (
	"math"
	g "github.com/joshua-wright/go-graphics/graphics"
	"sync"
	"fmt"
	"time"
)

type pixelUpdateParallel struct {
	x, y         int
	fromX, fromY int
}

type pixelParallel struct {
	x, y       int
	minX, minY int
	minDist2   int // integer distance only because on grid and distance squared
	queue      chan pixelUpdateParallel
}

func (p *pixelParallel) worker(quit <-chan bool, width, height int, mesh [][]pixelParallel, wg *sync.WaitGroup) {
	//var u pixelUpdateParallel
	for {
		select {
		case _ = <-quit:
			return

		case u := <-p.queue:
			wg.Done()
			fmt.Println(p, "received", u)
			// compute new distance
			dx := p.x - u.x
			dy := p.y - u.y
			new_dist2 := dx*dx + dy*dy

			if new_dist2 < p.minDist2 {
				fmt.Println(p, "updating dist")
				// update closest pixelParallel
				p.minX = u.x
				p.minY = u.y
				p.minDist2 = new_dist2

				// update neighbors
				newMessage := pixelUpdateParallel{u.x, u.y, p.x, p.y}
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						if i == j {
							continue
						}
						x := p.x + i
						y := p.y + j
						if x < 0 || x >= width || y < 0 || y >= width {
							continue
						}
						if x == u.fromX && y == u.fromY {
							continue
						}
						fmt.Println(p, "sending to", x, y)
						wg.Add(1)
						mesh[x][y].queue <- newMessage
					}
				}
			}
			fmt.Println(p, "done")
		}
	}
}

// TODO: this 100%  does not work
func distanceTransformParallel(width, height int, zero_points []g.Vec2) [][]float64 {
	//width := img.Bounds().Dx()
	//height := img.Bounds().Dy()
	mesh := Make2DSlicePixelParallel(width, height, pixelParallel{})
	output_mesh := Make2DSliceFloat64(width, height, 0.0)
	quit := make(chan bool)
	wg := &sync.WaitGroup{}

	// start workers
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			mesh[x][y] = pixelParallel{
				x:        x,
				y:        y,
				minX:     -1,
				minY:     -1,
				minDist2: math.MaxInt64,
				queue:    make(chan pixelUpdateParallel, 5),
			}
			go mesh[x][y].worker(quit, width, height, mesh, wg)
		}
	}

	// send initial control point messages
	//wg.Add(len(zero_points))
	for _, p := range zero_points {
		x := int(p.X)
		y := int(p.Y)
		wg.Add(1)
		mesh[x][y].queue <- pixelUpdateParallel{x, y, x, y}
	}

	// wait for everything to finish
	//wg.Wait()

	fmt.Println("waited")
	time.Sleep(1 * time.Second)

	g.Parallel(
		func() {
			// close all pixelParallel goroutines
			for i := 0; i < width*height; i++ {
				quit <- true
			}
		},
		func() {
			// copy output into buffer
			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					output_mesh[x][y] = math.Sqrt(float64(mesh[x][y].minDist2))
				}
			}
		},
	)

	return output_mesh
}
