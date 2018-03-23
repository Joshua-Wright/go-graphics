package per_pixel_image

import (
	"os"
	g "github.com/joshua-wright/go-graphics/graphics"
	"github.com/willf/bitset"
	"image/color"
	"sync"
	"runtime"
	"gopkg.in/cheggaaa/pb.v1"
	"time"
	"io"
	"fmt"
)

// assumes pixel function writes its own pixels
type PixelFunction interface {
	GetPixel(i, j int64)
	Bounds() (w int64, h int64)
}

type donePixel struct {
	X, Y int
	Pix  color.RGBA
}

// must be multiple of 64 (word size) for bitmap to be (assumed) thread safe
const jobSize int64 = 64 * 256

func pixelRowWorker(
	doneMask *bitset.BitSet,
	pixelFunc PixelFunction,
	jobs chan int64,
	wg *sync.WaitGroup,
) {
	w, h := pixelFunc.Bounds()
	size := w * h
	for start := range jobs {
		end := start + jobSize
		if jobSize >= size {
			end = size
		}
		for i := start; i < end; i++ {
			if !doneMask.Test(uint(i)) {
				x := i % w
				y := i / w
				pixelFunc.GetPixel(x, y)
				doneMask.Set(uint(i))
			}
		}
		wg.Done()
	}
}

// TODO better name
func PerPixelImage(pixelFunc PixelFunction, doneMaskFilename string) error {
	width, height := pixelFunc.Bounds()
	var err error

	// open bitset
	var doneMask = bitset.New(uint(width * height))
	var doneMaskFile *os.File
	if g.FileExists(doneMaskFilename) {
		println("resuming")
		// load bitmap
		doneMaskFile, err = os.Open(doneMaskFilename)
		_, err = doneMask.ReadFrom(doneMaskFile)
	} else {
		println("starting fresh")
		doneMaskFile, err = os.Create(doneMaskFilename)
	}
	if err != nil {
		return err
	}

	jobs := make(chan int64)
	var wg sync.WaitGroup

	// start workers
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go pixelRowWorker(doneMask, pixelFunc, jobs, &wg)
	}

	numPixels := width * height
	numTasks := 1 + ((numPixels - 1) / jobSize)

	// start bitmap saver
	bar := pb.StartNew(int(numTasks))

	reducerQuit := make(chan struct{})
	go func() {
		//ticker := time.NewTicker(5 * time.Minute)
		ticker := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				bar.Prefix("writing checkpoint")
				// occassionally write doneMask to file
				_, err := doneMaskFile.Seek(0, io.SeekStart)
				if err != nil {
					fmt.Println(err)
				}
				doneMask.WriteTo(doneMaskFile)
				bar.Prefix("")
			case <-reducerQuit:
				break
			}
		}
		bar.Finish()
	}()

	// queue

	wg.Add(int(numTasks))
	for i := int64(0); i < numTasks; i++ {
		jobs <- i * jobSize
		bar.Increment()
	}
	close(jobs)
	wg.Wait()
	reducerQuit <- struct{}{}

	doneMaskFile.Close()
	os.Remove(doneMaskFilename)

	return nil
}
