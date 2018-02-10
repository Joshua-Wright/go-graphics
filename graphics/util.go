package graphics

import (
	"math"
	"log"
	"path/filepath"
	"os"
	"fmt"
	"sync"
	"image"
	"image/png"
	"runtime"
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

func Vec2Midpoint(a, b Vec2) Vec2 { return a.AddV(b).DivS(2.0) }
func Vec3Midpoint(a, b Vec3) Vec3 { return a.AddV(b).DivS(2.0) }
func Vec4Midpoint(a, b Vec4) Vec4 { return a.AddV(b).DivS(2.0) }

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
func ExecutableFolderFileName(filename string) string {
	os.Mkdir(ExecutableName(), 077)
	return filepath.Join(ExecutableName(), filename)
}

func SaveAsPNG(img image.Image, filename string) {
	file, err := os.Create(filename)
	Die(err)
	Die(png.Encode(file, img))
	Die(file.Close())
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

func parallelForWorker(wg *sync.WaitGroup, jobs chan int, f func(int)) {
	for i := range jobs {
		f(i)
		wg.Done()
	}
}

func ParallelFor(start, end int, f func(int)) {
	jobs := make(chan int)
	var wg sync.WaitGroup

	// start workers
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go parallelForWorker(&wg, jobs, f)
	}

	// queue
	wg.Add(end - start)
	for i := start; i < end; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func Parallel(funcs ...func()) {
	ParallelFor(0, len(funcs), func(i int) {
		funcs[i]()
	})
}
