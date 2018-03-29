package main

import (
	"os"
	"strings"
	"path/filepath"
	"image/png"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	"strconv"
	g "github.com/joshua-wright/go-graphics/graphics"
	"image/color"
	"image"
	"runtime"
)

const bufferSize = 10240

type pixelMsg struct {
	index int64
	c     color.Color
}

type downsampledPpm struct {
	width, height int64
	pixels        chan pixelMsg
}

func (img *downsampledPpm) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *downsampledPpm) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(img.width), int(img.height))
}

func (img *downsampledPpm) At(x_, y_ int) color.Color {
	x := int64(x_)
	y := int64(y_)
	px := <-img.pixels
	if px.index != y*img.width+x {
		panic("bad pixel order access")
	}
	//if x == 0 {
	//	println(x, y, px.index)
	//}
	return px.c
}

func DownsamplePixel(ppm *memory_mapped.PPMFile, index, factor int64, out chan pixelMsg) {
	x := (index % ppm.W) * factor
	y := (index / ppm.W) * factor

	var rs float64
	var gs float64
	var bs float64

	for j := y; j < y+factor; j++ {
		for i := x; i < x+factor; i++ {
			r, g, b := ppm.At64(i, j)
			rs += float64(r)
			gs += float64(g)
			bs += float64(b)
		}
	}

	d := float64(factor * factor)

	out <- pixelMsg{
		index: index,
		c: color.RGBA{
			R: uint8(rs / d),
			G: uint8(gs / d),
			B: uint8(bs / d),
			A: 255,
		},
	}
}

func (img *downsampledPpm) worker(ppm *memory_mapped.PPMFile, factor int64) {
	unsortedPixels := make(chan pixelMsg, bufferSize)

	lastIndex := int64(runtime.GOMAXPROCS(-1))
	nextSortedIndex := int64(0)

	// start initial workers
	for i := int64(0); i < lastIndex; i++ {
		go DownsamplePixel(ppm, i, factor, unsortedPixels)
	}

	cache := make(map[int64]color.Color)

	for px := range unsortedPixels {
		// if we're ready for this, send it
		if px.index == nextSortedIndex {
			img.pixels <- px
			nextSortedIndex++
			// start next pixel
			if lastIndex < ppm.W * ppm.H {
				go DownsamplePixel(ppm, lastIndex, factor, unsortedPixels)
				lastIndex++
			}

			// check what else we can send also
			for {
				if px2, ok := cache[nextSortedIndex]; ok {
					delete(cache, nextSortedIndex)
					img.pixels <- pixelMsg{nextSortedIndex, px2}
					nextSortedIndex++
					if lastIndex < ppm.W * ppm.H {
						go DownsamplePixel(ppm, lastIndex, factor, unsortedPixels)
						lastIndex++
					}
				} else {
					break
				}
			}
		} else {
			// otherwise put it in the cache
			cache[px.index] = px.c
		}
	}
}

func MakeDownsampler(ppm *memory_mapped.PPMFile, factor int64) *downsampledPpm {
	pixels := make(chan pixelMsg, bufferSize)

	out := downsampledPpm{
		width:  ppm.W,
		height: ppm.H,
		pixels: pixels,
	}

	go out.worker(ppm, factor)

	return &out
}

func main() {
	filename := os.Args[1]
	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".png"

	ppmImage, err := memory_mapped.OpenPPM(filename)
	g.Die(err)

	factor, err := strconv.Atoi(os.Args[2])
	g.Die(err)

	if ppmImage.W%int64(factor) != 0 {
		panic("bad width")
	}
	if ppmImage.H%int64(factor) != 0 {
		panic("bad height")
	}

	sampled := MakeDownsampler(ppmImage, int64(factor))

	file, err := os.Create(newFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, sampled)
	if err != nil {
		panic(err)
	}
}
