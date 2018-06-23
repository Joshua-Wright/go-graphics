package per_pixel_image

import (
	"os"
	"sync"
	"runtime"
	"gopkg.in/cheggaaa/pb.v1"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
)

// assumes pixel function writes its own pixels
type PixelFunction interface {
	GetPixel(i, j int64)
	Bounds() (w int64, h int64)
}

type PixelJobSizeOverride interface {
	GetJobSize() int64
}

const jobSize int64 = 64 * 1024

func pixelRowWorker(
	doneMask *memory_mapped.AtomicBitset,
	pixelFunc PixelFunction,
	jobs chan int64,
	wg *sync.WaitGroup,
) {
	var localJobSize = jobSize
	if override, ok := pixelFunc.(PixelJobSizeOverride); ok {
		localJobSize = override.GetJobSize()
	}
	w, h := pixelFunc.Bounds()
	size := w * h
	for start := range jobs {
		end := start + localJobSize
		if end >= size {
			end = size
		}
		for i := start; i < end; i++ {
			if !doneMask.Test(i) {
				x := i % w
				y := i / w
				pixelFunc.GetPixel(x, y)
				doneMask.Set(i)
			}
		}
		wg.Done()
	}
}

// TODO better name
func PerPixelImage(pixelFunc PixelFunction, doneMaskFilename string) error {
	width, height := pixelFunc.Bounds()
	numPixels := width * height

	var localJobSize = jobSize
	if override, ok := pixelFunc.(PixelJobSizeOverride); ok {
		localJobSize = override.GetJobSize()
	}

	var err error

	// open bitset
	doneMask, err := memory_mapped.OpenOrCreateAtomicBitset(numPixels, doneMaskFilename)
	if err != nil {
		return err
	}

	jobs := make(chan int64)
	var wg sync.WaitGroup

	// start workers
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go pixelRowWorker(doneMask, pixelFunc, jobs, &wg)
	}

	numTasks := 1 + ((numPixels - 1) / localJobSize)

	// start progress bar
	bar := pb.New64(numTasks)
	bar.Start()

	// queue
	wg.Add(int(numTasks))
	for i := int64(0); i < numTasks; i++ {
		jobs <- i * localJobSize
		bar.Increment()
	}
	close(jobs)
	wg.Wait()

	os.Remove(doneMaskFilename)
	return nil
}
