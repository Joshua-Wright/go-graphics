package per_pixel_image

import (
	"os"
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/joshua-wright/go-graphics/graphics/file_backed_image"
	"github.com/pkg/errors"
	"github.com/willf/bitset"
	"image/color"
	"sync"
	"runtime"
	"gopkg.in/cheggaaa/pb.v1"
	"time"
)

type PixelFunction interface {
	GetPixel(i, j int) color.RGBA
	Bounds() (w int, h int)
}

type donePixel struct {
	X, Y int
	Pix  color.RGBA
}

func pixelRowWorker(
	doneMask *bitset.BitSet,
	pixelFunc PixelFunction,
	jobs chan int, done chan donePixel,
) {
	w, h := pixelFunc.Bounds()
	for x := range jobs {
		for y := 0; y < h; y++ {
			if !doneMask.Test(uint(w*y + x)) {
				pixel := pixelFunc.GetPixel(x, y)
				done <- donePixel{x, y, pixel}
			}
		}
	}
}

// TODO better name
func PerPixelImage(pixelFunc PixelFunction, foldername string) error {
	width, height := pixelFunc.Bounds()
	// OK if folder already exists
	os.Mkdir(foldername, 0777)
	err := os.Chdir(foldername)
	if err != nil {
		return err
	}

	ppmFilename := foldername + ".ppm"
	bitmapFilename := foldername + ".bitmap"
	// TODO save/load config
	//configFilename := foldername + ".json"

	// open image file and bitset
	var doneMask = bitset.New(uint(width * height))
	var doneMaskFile *os.File
	var img *file_backed_image.PPMFile
	if g.FileExists(ppmFilename) {
		img, err = file_backed_image.OpenPPM(ppmFilename)
		if err != nil {
			return err
		}
		// also load bitmap
		doneMaskFile, err = os.Open(bitmapFilename)
		if err != nil {
			return err
		}
		defer doneMaskFile.Close()
		doneMask.ReadFrom(doneMaskFile)
	} else {
		img, err = file_backed_image.CreatePPM(width, height, ppmFilename)
		if err != nil {
			return err
		}
		doneMaskFile, err = os.Create(bitmapFilename)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	if int(img.W) != width || int(img.H) != height {
		return errors.New("bad previous data")
	}

	jobs := make(chan int, width)
	donePixels := make(chan donePixel, width)
	var wg sync.WaitGroup

	// start workers
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go pixelRowWorker(doneMask, pixelFunc, jobs, donePixels)
	}

	// start reducer
	bar := pb.StartNew(width*height - int(doneMask.Count()))
	bar.Set(int(doneMask.Count()))
	reducerQuit := make(chan struct{})
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for {
			select {
			case <-ticker.C:
				bar.Prefix("writing checkpoint")
				// occassionally write doneMask to file
				doneMask.WriteTo(doneMaskFile)
				bar.Prefix("")

			case pix := <-donePixels:
				img.Set(pix.X, pix.Y, pix.Pix)
				doneMask.Set(uint(width*pix.Y + pix.X))
				wg.Done()
				bar.Increment()
			case <-reducerQuit:
				break
			}
		}
		bar.Finish()
	}()

	// queue
	wg.Add(width*height - int(doneMask.Count()))
	go func() {
		for i := 0; i < width; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	wg.Wait()
	reducerQuit <- struct{}{}

	return nil
}
