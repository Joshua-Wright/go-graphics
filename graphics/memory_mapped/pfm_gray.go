package memory_mapped

import (
	"os"
	"syscall"
	"golang.org/x/sys/unix"
	"image"
	"image/color"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"github.com/edsrzf/mmap-go"
	"math"
)

type PFMGrayFile struct {
	W, H       int64
	Filename   string
	File       *os.File
	mappedFile mmap.MMap
	headerSize int64
}

// one 32-bit float per pixel
const PFM_GRAY_BYTES_PER_PIXEL = 4

func (p *PFMGrayFile) Close() error {
	p.mappedFile.Unmap()
	err := p.mappedFile.Unmap()
	if err != nil {
		return err
	}
	// reset all fields
	*p = PFMGrayFile{}
	return nil
	return p.File.Close()
}

func (p *PFMGrayFile) getOffset(x, y int64) int64 {
	if x < 0 || y < 0 || x >= p.W || y >= p.H {
		panic(fmt.Sprintf("(%v,%v) out of image bounds (%v,%v)", x, y, p.W, p.H))
	}
	return p.headerSize + PFM_GRAY_BYTES_PER_PIXEL*(p.W*y+x)
}

func (p *PFMGrayFile) Set(x, y int, c color.Color) {
	grayColor := color.Gray16Model.Convert(c).(color.Gray16)
	p.SetFloat(int64(x), int64(y), float32(grayColor.Y))
}

func (p *PFMGrayFile) Set64(x, y int64, c color.Color) {
	grayColor := color.Gray16Model.Convert(c).(color.Gray16)
	p.SetFloat(x, y, float32(grayColor.Y))
}

func (p *PFMGrayFile) SetFloat(x, y int64, f float32) {
	offset := p.getOffset(x, y)
	bits := math.Float32bits(f)
	p.mappedFile[offset+0] = byte(bits >> 24)
	p.mappedFile[offset+1] = byte(bits >> 16)
	p.mappedFile[offset+2] = byte(bits >> 8)
	p.mappedFile[offset+3] = byte(bits)
}

func (p *PFMGrayFile) GetFloat(x, y int64) float32 {
	offset := p.getOffset(x, y)
	bits := uint32(p.mappedFile[offset+0])<<24 |
		uint32(p.mappedFile[offset+1])<<16 |
		uint32(p.mappedFile[offset+2])<<8 |
		uint32(p.mappedFile[offset+3])
	return math.Float32frombits(bits)
}

func (p *PFMGrayFile) ColorModel() color.Model {
	return color.Gray16Model
}

// if we tell the PNG encoder that we are opaque, it will traverse the image exactly once
func (*PFMGrayFile) Opaque() bool {
	return true
}

func (p *PFMGrayFile) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(p.W), int(p.H))
}

func (p *PFMGrayFile) At(x, y int) color.Color {
	f := p.GetFloat(int64(x), int64(y))
	return color.Gray16{uint16(f)}
}

func OpenOrCreatePFMGray(width, height int64, filename string) (*PFMGrayFile, error) {
	img, err := OpenPFMGray(filename)
	if err != nil {
		// file doesn't exist, so create it
		img, err = CreatePFMGray(width, height, filename)
		if err != nil {
			return nil, err
		} else {
			return img, nil
		}
	} else {
		// file does exist, so verify it's still valid
		if img.W != width || img.H != height {
			return nil, errors.New(
				fmt.Sprintf("bad width or height: file (%v, %v), requested: (%v, %v)",
					img.W, img.H, width, height))
		} else {
			return img, nil
		}
	}
}

func CreatePFMGray(width, height int64, filename string) (*PFMGrayFile, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	// write header
	var bytes1 int
	if bytes1, err = fmt.Fprintln(file, "Pf"); err != nil {
		return nil, err
	}

	// x, y, depth
	bytes2, err := fmt.Fprintf(file, "%d %d\n1.0\n", width, height)
	if err != nil {
		return nil, err
	}

	headerSize := int64(bytes1 + bytes2)

	// allocate space in the file
	err = syscall.Fallocate(int(file.Fd()), unix.FALLOC_FL_ZERO_RANGE,
		headerSize, int64(width*height*PFM_GRAY_BYTES_PER_PIXEL))
	if err != nil {
		return nil, err
	}

	mappedFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	pgm := PFMGrayFile{
		W:          width,
		H:          height,
		Filename:   filename,
		File:       file,
		mappedFile: mappedFile,
		headerSize: headerSize,
	}
	return &pgm, nil
}

func OpenPFMGray(filename string) (*PFMGrayFile, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// only bother to support the exact format that the create function writes
	var width, height int64
	n, err := fmt.Fscanf(file, "Pf\n%d %d\n1.0\n", &width, &height)
	if err != nil {
		return nil, err
	}
	if n != 2 {
		return nil, errors.New("failed to parse header")
	}

	headerSize, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	mappedFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	pgm := PFMGrayFile{
		W:          int64(width),
		H:          int64(height),
		Filename:   filename,
		File:       file,
		mappedFile: mappedFile,
		headerSize: headerSize,
	}
	return &pgm, nil
}
