package golang_raytracer

import (
	"math"
	"log"
	"path/filepath"
	"os"
	"fmt"
	"sync"
)

//type Float = float32
type Float = float64

func Sqrt(f Float) Float { return Float(math.Sqrt(float64(f))) }
func Sin(f Float) Float  { return Float(math.Sin(float64(f))) }
func Cos(f Float) Float  { return Float(math.Cos(float64(f))) }

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r *Ray) PointAt(t Float) Vec3 {
	v := r.Direction.MulS(t)
	return v.AddV(r.Origin)
}

func Die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ExecutableName() string { return filepath.Base(os.Args[0]) }
func ExecutableNameWithExtension(s string) string {
	return fmt.Sprintf("%s.%s", ExecutableName(), s)
}
func ExecutableNamePng() string {
	return fmt.Sprintf("%s.png", ExecutableName())
}

func MaxAdjacentDistance(pts []Vec2) Float {
	dmax := pts[0].SubV(pts[1]).Mag2()
	for i := 1; i < len(pts)-1; i++ {
		d2 := pts[i].SubV(pts[i+1]).Mag2()
		if d2 > dmax {
			dmax = d2
		}
	}
	return Sqrt(dmax)
}

func ParallelFor(start, end int, f func(int)) {
	var wg sync.WaitGroup
	for i := start; i < end; i++ {
		wg.Add(1)
		go func(i int) {
			f(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Parallel(funcs ...func()) {
	ParallelFor(0, len(funcs), func(i int) {
		funcs[i]()
	})
}
