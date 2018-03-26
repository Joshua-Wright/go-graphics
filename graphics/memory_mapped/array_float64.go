package memory_mapped

import (
	"os"
	"github.com/edsrzf/mmap-go"
	"unsafe"
	"fmt"
	"syscall"
	"golang.org/x/sys/unix"
	"io"
	"errors"
	"math"
)

type ArrayFloat64 struct {
	length     int64
	filename   string
	mappedFile mmap.MMap
	headerSize uintptr
}

func (a *ArrayFloat64) Close() error {
	// don't bother flushing, the kernel will take care of it eventually
	//err := a.mappedFile.Flush()
	//if err != nil {
	//	return err
	//}
	err := a.mappedFile.Unmap()
	if err != nil {
		return err
	}
	// reset all fields
	*a = ArrayFloat64{}
	return nil
}

const bytesPerFloat64 = 8

func (a *ArrayFloat64) getPointerAtIndex(i int64) *float64 {
	offset := a.headerSize + bytesPerFloat64*uintptr(i)
	return (*float64)(unsafe.Pointer(&a.mappedFile[offset]))
}

func (a *ArrayFloat64) getOffset(i int64) int64 {
	offset := a.headerSize + bytesPerFloat64*uintptr(i)
	return int64(offset)
}

func (a *ArrayFloat64) Get(i int64) float64 {
	offset := a.getOffset(i)
	return math.Float64frombits(uint64(a.mappedFile[offset+0])<<56 |
		uint64(a.mappedFile[offset+1])<<48 |
		uint64(a.mappedFile[offset+2])<<40 |
		uint64(a.mappedFile[offset+3])<<32 |
		uint64(a.mappedFile[offset+4])<<24 |
		uint64(a.mappedFile[offset+5])<<16 |
		uint64(a.mappedFile[offset+6])<<8 |
		uint64(a.mappedFile[offset+7])<<0)
}

func (a *ArrayFloat64) Set(i int64, v float64) {
	n := math.Float64bits(v)
	offset := a.getOffset(i)
	a.mappedFile[offset+0] = byte(n >> 56)
	a.mappedFile[offset+1] = byte(n >> 48)
	a.mappedFile[offset+2] = byte(n >> 40)
	a.mappedFile[offset+3] = byte(n >> 32)
	a.mappedFile[offset+4] = byte(n >> 24)
	a.mappedFile[offset+5] = byte(n >> 16)
	a.mappedFile[offset+6] = byte(n >> 8)
	a.mappedFile[offset+7] = byte(n)
}

func (a *ArrayFloat64) Len() int64       { return a.length }
func (a *ArrayFloat64) Filename() string { return a.filename }

func OpenOrCreateArrayFloat64(len int64, filename string) (*ArrayFloat64, error) {
	arr, err := OpenArrayFloat64(filename)
	if err != nil {
		// file doesn't exist, so create it
		arr, err = CreateArrayFloat64(len, filename)
		if err != nil {
			return nil, err
		} else {
			return arr, nil
		}
	} else {
		// file does exist, so verify it's still valid
		if arr.length != len {
			return nil, errors.New(fmt.Sprintf("bad length: file (%v), requested: (%v)", arr.length, len))
		} else {
			return arr, nil
		}
	}
}

func CreateArrayFloat64(length int64, filename string) (*ArrayFloat64, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// write header
	var bytes1 int
	if bytes1, err = fmt.Fprintf(file, "float64\n%d\n", length); err != nil {
		return nil, err
	}

	headerSize := int64(bytes1)

	// allocate space in the file
	err = syscall.Fallocate(int(file.Fd()), unix.FALLOC_FL_ZERO_RANGE,
		headerSize, headerSize+int64(length*bytesPerFloat64))
	if err != nil {
		return nil, err
	}

	mappedFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	pgm := ArrayFloat64{
		length:     length,
		filename:   filename,
		mappedFile: mappedFile,
		headerSize: uintptr(headerSize),
	}
	return &pgm, nil
}

func OpenArrayFloat64(filename string) (*ArrayFloat64, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// only bother to support the exact format that the create function writes
	var length int64
	n, err := fmt.Fscanf(file, "float64\n%d\n", &length)
	if err != nil {
		return nil, err
	}
	if n != 1 {
		return nil, errors.New("failed to parse header")
	}
	if length <= 0 {
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

	pgm := ArrayFloat64{
		length:     length,
		filename:   filename,
		mappedFile: mappedFile,
		headerSize: uintptr(headerSize),
	}
	return &pgm, nil
}
