package main

import (
	"os"
	"strings"
	"path/filepath"
	"image/png"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	g "github.com/joshua-wright/go-graphics/graphics"
	"image/color"
	"image"
	"runtime"
	"strconv"
)

const bufferSize = 10240

type pixelMsg struct {
	index int64
	c     color.Color
}

type downsampledPpm struct {
	outWidth, outHeight  int64
	pixelRequest         chan int64
	pixelRequestResponse chan color.Color
}

func (img *downsampledPpm) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *downsampledPpm) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(img.outWidth), int(img.outHeight))
}

func (img *downsampledPpm) At(x_, y_ int) color.Color {
	x := int64(x_)
	y := int64(y_)
	idx := y*img.outWidth + x
	img.pixelRequest <- idx
	return <-img.pixelRequestResponse
}

func DownsamplePixel(ppm *memory_mapped.PPMFile, index, factor int64, out chan pixelMsg) {
	w := ppm.W / factor
	x := (index % w) * factor
	y := (index / w) * factor

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

	maxCacheSize := 1024 * 10
	nextIndex := int64(0)
	maxIndex := img.outWidth * img.outHeight
	numRunningWorkers := 0
	maxNumRunningWorkers := runtime.GOMAXPROCS(-1)

	println("maxIndex ", maxIndex)

	// start initial workers
	for i := int64(0); i < nextIndex; i++ {
		go DownsamplePixel(ppm, nextIndex, factor, unsortedPixels)
		numRunningWorkers++
		nextIndex++
	}

	cache := make(map[int64]color.Color)
	waitingFor := int64(-1)
	for {
		select {

		case reqIdx := <-img.pixelRequest:
			if px, ok := cache[reqIdx]; ok {
				delete(cache, reqIdx)
				img.pixelRequestResponse <- px
			} else {
				waitingFor = reqIdx
				// start a new runner for this, possibly creating too many runners and possibly doing extra work
				go DownsamplePixel(ppm, reqIdx, factor, unsortedPixels)
				numRunningWorkers++
				nextIndex = reqIdx + 1
			}

		case pxMsg := <-unsortedPixels:
			if pxMsg.index == waitingFor {
				// if we're waiting for it, send it directly
				img.pixelRequestResponse <- pxMsg.c
				waitingFor = -1 // probably overkill
			} else {
				cache[pxMsg.index] = pxMsg.c
			}
			numRunningWorkers--

		}

		// start more workers if necessary
		if len(cache) < maxCacheSize && numRunningWorkers < maxNumRunningWorkers && nextIndex < maxIndex {
			for i := 0; i < (maxNumRunningWorkers - numRunningWorkers); i++ {
				go DownsamplePixel(ppm, nextIndex, factor, unsortedPixels)
				numRunningWorkers++
				nextIndex++
			}
		}

	}
}

func MakeDownsampler(ppm *memory_mapped.PPMFile, factor int64) *downsampledPpm {
	out := downsampledPpm{
		outWidth:             ppm.W / factor,
		outHeight:            ppm.H / factor,
		pixelRequest:         make(chan int64),
		pixelRequestResponse: make(chan color.Color),
	}

	go out.worker(ppm, factor)

	return &out
}

func main() {
	filename := os.Args[1]
	factor, err := strconv.Atoi(os.Args[2])
	g.Die(err)
	//filename := "/home/j0sh/Documents/code/go-graphics/src/github.com/joshua-wright/go-graphics/util/mandelbrot.ppm"
	//factor := 10
	//runtime.GOMAXPROCS(1)

	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".png"

	ppmImage, err := memory_mapped.OpenPPM(filename)
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
