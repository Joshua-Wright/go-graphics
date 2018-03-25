package memory_mapped

import (
	"github.com/edsrzf/mmap-go"
	"unsafe"
	"sync/atomic"
	"fmt"
	"os"
	"syscall"
	"golang.org/x/sys/unix"
	"errors"
)

type AtomicBitset struct {
	length     int64
	filename   string
	mappedFile mmap.MMap
	// no header
}

func (b *AtomicBitset) Close() error {
	err := b.mappedFile.Flush()
	if err != nil {
		return err
	}
	err = b.mappedFile.Unmap()
	if err != nil {
		return err
	}
	// reset all fields
	*b = AtomicBitset{}
	return nil
}

const log2WordSize = 6
const bitsPerUInt64 = 64
const bytesPerUInt64 = 8

func (b *AtomicBitset) Len() int64       { return b.length }
func (b *AtomicBitset) Filename() string { return b.filename }

func (b *AtomicBitset) getPointerAtIndex(i int64) *uint64 {
	offset := uintptr((i >> log2WordSize) << 3)
	return (*uint64)(unsafe.Pointer(&b.mappedFile[offset]))
}

func (b *AtomicBitset) Test(i_ int64) bool {
	if i_ >= b.length {
		panic("out of bounds")
	}
	addr := b.getPointerAtIndex(i_)
	i := uint64(i_)
	word := atomic.LoadUint64(addr)
	return (word & (1 << (i & (bitsPerUInt64 - 1))) ) != 0
}

func (b *AtomicBitset) Set(i_ int64) {
	addr := b.getPointerAtIndex(i_)
	i := uint64(i_)
	for {
		oldWord := atomic.LoadUint64(addr)
		newWord := oldWord | (1 << (i & (bitsPerUInt64 - 1)))
		if atomic.CompareAndSwapUint64(addr, oldWord, newWord) {
			return
		}
	}
}

func (b *AtomicBitset) Clear(i_ int64) {
	addr := b.getPointerAtIndex(i_)
	i := uint64(i_)
	for {
		oldWord := atomic.LoadUint64(addr)
		newWord := oldWord &^ 1 << (i & (bitsPerUInt64 - 1))
		if atomic.CompareAndSwapUint64(addr, oldWord, newWord) {
			return
		}
	}
}

func OpenOrCreateAtomicBitset(length int64, filename string) (*AtomicBitset, error) {
	arr, err := OpenAtomicBitset(filename)
	if err != nil {
		// file doesn't exist, so create it
		arr, err = CreateAtomicBitset(length, filename)
		if err != nil {
			return nil, err
		} else {
			return arr, nil
		}
	} else {
		// file does exist, so verify it's still valid
		if arr.length != length {
			return nil, errors.New(fmt.Sprintf("bad length: file (%v), requested: (%v)", arr.length, length))
		} else {
			return arr, nil
		}
	}
}

func CreateAtomicBitset(length int64, filename string) (*AtomicBitset, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// allocate space in the file
	err = syscall.Fallocate(int(file.Fd()), unix.FALLOC_FL_ZERO_RANGE, 0, int64(length/8))
	if err != nil {
		return nil, err
	}

	mappedFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	pgm := AtomicBitset{
		length:     length,
		filename:   filename,
		mappedFile: mappedFile,
	}
	return &pgm, nil
}

func OpenAtomicBitset(filename string) (*AtomicBitset, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	length := stat.Size() * 8

	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mappedFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	pgm := AtomicBitset{
		length:     length,
		filename:   filename,
		mappedFile: mappedFile,
	}
	return &pgm, nil
}
