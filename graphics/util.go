package graphics

import (
	"math"
	"log"
	"path/filepath"
	"os"
	"fmt"
	"image"
	"image/png"
)

//type Float = float32
type Float = float64

func Sqrt(f Float) Float { return Float(math.Sqrt(float64(f))) }
func Sin(f Float) Float  { return Float(math.Sin(float64(f))) }
func Cos(f Float) Float  { return Float(math.Cos(float64(f))) }

func Lerp(v0, v1, t float64) float64 {
	// Precise method, which guarantees v = v1 when t = 1 (from wikipedia)
	return (1-t)*v0 + t*v1
}

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
func ExecutableFolderFileName(filename string) string {
	os.Mkdir(ExecutableName()+"_frames", 0777)
	return filepath.Join(ExecutableName()+"_frames", filename)
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
