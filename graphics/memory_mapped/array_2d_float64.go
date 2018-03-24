package memory_mapped

import (
	"os"
	"github.com/edsrzf/mmap-go"
	"fmt"
	"syscall"
	"golang.org/x/sys/unix"
	"io"
	"errors"
	"math"
)

type Array2dFloat64 struct {
	ArrayFloat64
	width, height int64
}

func (a *Array2dFloat64) Close() error {
	err := a.mappedFile.Flush()
	if err != nil {
		return err
	}
	err = a.mappedFile.Unmap()
	if err != nil {
		return err
	}
	// reset all fields
	*a = Array2dFloat64{}
	return nil
}

func (a *Array2dFloat64) get2dOffset(i, j int64) int64 {
	return a.getOffset(j*a.width + i)
}

func (a *Array2dFloat64) Get(i, j int64) float64 {
	offset := a.get2dOffset(i, j)
	return math.Float64frombits(uint64(a.mappedFile[offset+0])<<56 |
		uint64(a.mappedFile[offset+1])<<48 |
		uint64(a.mappedFile[offset+2])<<40 |
		uint64(a.mappedFile[offset+3])<<32 |
		uint64(a.mappedFile[offset+4])<<24 |
		uint64(a.mappedFile[offset+5])<<16 |
		uint64(a.mappedFile[offset+6])<<8 |
		uint64(a.mappedFile[offset+7])<<0)
}

func (a *Array2dFloat64) Set(i, j int64, v float64) {
	n := math.Float64bits(v)
	offset := a.get2dOffset(i, j)
	a.mappedFile[offset+0] = byte(n >> 56)
	a.mappedFile[offset+1] = byte(n >> 48)
	a.mappedFile[offset+2] = byte(n >> 40)
	a.mappedFile[offset+3] = byte(n >> 32)
	a.mappedFile[offset+4] = byte(n >> 24)
	a.mappedFile[offset+5] = byte(n >> 16)
	a.mappedFile[offset+6] = byte(n >> 8)
	a.mappedFile[offset+7] = byte(n)
}

func (a *Array2dFloat64) Len() int64       { return a.length }
func (a *Array2dFloat64) Width() int64     { return a.width }
func (a *Array2dFloat64) Height() int64    { return a.height }
func (a *Array2dFloat64) Filename() string { return a.filename }

func OpenOrCreateMmap2dArrayFloat64(width, height int64, filename string) (*Array2dFloat64, error) {
	arr, err := OpenMmap2dArrayFloat64(filename)
	if err != nil {
		// file doesn't exist, so create it
		arr, err = CreateMmap2dArrayFloat64(width, height, filename)
		if err != nil {
			return nil, err
		} else {
			return arr, nil
		}
	} else {
		// file does exist, so verify it's still valid
		if arr.width != width || arr.height != height {
			return nil, errors.New(
				fmt.Sprintf("bad width or height: file (%v, %v), requested: (%v, %v)",
					arr.width, arr.height, width, height))
		} else {
			return arr, nil
		}
	}
}

func CreateMmap2dArrayFloat64(width, height int64, filename string) (*Array2dFloat64, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// write header
	var bytes1 int
	if bytes1, err = fmt.Fprintf(file, "float64 2d\n%d %d\n", width, height); err != nil {
		return nil, err
	}

	headerSize := int64(bytes1)

	// allocate space in the file
	err = syscall.Fallocate(int(file.Fd()), unix.FALLOC_FL_ZERO_RANGE,
		headerSize, headerSize+int64(width*height*bytesPerFloat64))
	if err != nil {
		return nil, err
	}

	mappedFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	pgm := Array2dFloat64{
		ArrayFloat64: ArrayFloat64{
			length:     width * height,
			filename:   filename,
			mappedFile: mappedFile,
			headerSize: uintptr(headerSize),
		},
		width:  width,
		height: height,
	}
	return &pgm, nil
}

func OpenMmap2dArrayFloat64(filename string) (*Array2dFloat64, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// only bother to support the exact format that the create function writes
	var width, height int64
	n, err := fmt.Fscanf(file, "float64 2d\n%d %d\n", &width, &height)
	if err != nil {
		return nil, err
	}
	if n != 2 {
		return nil, errors.New("failed to parse header")
	}
	if width <= 0 || height <= 0 {
		return nil, errors.New("bad header")
	}

	headerSize, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	mappedFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	pgm := Array2dFloat64{
		ArrayFloat64: ArrayFloat64{
			length:     width * height,
			filename:   filename,
			mappedFile: mappedFile,
			headerSize: uintptr(headerSize),
		},
		width:  width,
		height: height,
	}
	return &pgm, nil
}
