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
)

type PPMFile struct {
	W, H       int64
	Filename   string
	File       *os.File
	mappedFile mmap.MMap
	headerSize int64
}

func (p *PPMFile) Close() error {
	p.mappedFile.Unmap()
	err := p.mappedFile.Unmap()
	if err != nil {
		return err
	}
	// reset all fields
	*p = PPMFile{}
	return nil
	return p.File.Close()
}

func (p *PPMFile) getOffset(x, y int64) int64 {
	if x < 0 || y < 0 || x >= p.W || y >= p.H {
		panic(fmt.Sprintf("(%v,%v) out of image bounds (%v,%v)", x, y, p.W, p.H))
	}
	return p.headerSize + 3*(p.W*y+x)
}

func (p *PPMFile) Set(x, y int, c color.Color) {
	offset := p.getOffset(int64(x), int64(y))

	rgbColor := color.RGBAModel.Convert(c).(color.RGBA)
	p.mappedFile[offset+0] = rgbColor.R
	p.mappedFile[offset+1] = rgbColor.G
	p.mappedFile[offset+2] = rgbColor.B
}

func (p *PPMFile) Set64(x, y int64, c color.Color) {
	offset := p.getOffset(x, y)

	rgbColor := color.RGBAModel.Convert(c).(color.RGBA)
	p.mappedFile[offset+0] = rgbColor.R
	p.mappedFile[offset+1] = rgbColor.G
	p.mappedFile[offset+2] = rgbColor.B
}

func (p *PPMFile) Set64RGB(x, y int64, r, g, b uint8) {
	offset := p.getOffset(x, y)

	p.mappedFile[offset+0] = r
	p.mappedFile[offset+1] = g
	p.mappedFile[offset+2] = b
}

func (p *PPMFile) ColorModel() color.Model {
	return color.RGBAModel
}

// if we tell the PNG encoder that we are opaque, it will traverse the image exactly once
func (*PPMFile) Opaque() bool {
	return true
}

func (p *PPMFile) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(p.W), int(p.H))
}

func (p *PPMFile) At(x, y int) color.Color {
	offset := p.getOffset(int64(x), int64(y))
	return color.RGBA{
		R: p.mappedFile[offset+0],
		G: p.mappedFile[offset+1],
		B: p.mappedFile[offset+2],
		A: 255,
	}
}

func (p *PPMFile) At64(x, y int64) (r, g, b uint8) {
	offset := p.getOffset(x, y)
	return p.mappedFile[offset+0],
		p.mappedFile[offset+1],
		p.mappedFile[offset+2]
}

// RGB format, 8 bytes per color
const BYTES_PER_PIXEL = 3

func OpenOrCreatePPM(width, height int64, filename string) (*PPMFile, error) {
	img, err := OpenPPM(filename)
	if err != nil {
		// file doesn't exist, so create it
		img, err = CreatePPM(width, height, filename)
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

func CreatePPM(width, height int64, filename string) (*PPMFile, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	// write header
	var bytes1 int
	if bytes1, err = fmt.Fprintln(file, "P6"); err != nil {
		return nil, err
	}

	// x, y, depth
	bytes2, err := fmt.Fprintf(file, "%d %d\n255\n", width, height)
	if err != nil {
		return nil, err
	}

	headerSize := int64(bytes1 + bytes2)

	// allocate space in the file
	err = syscall.Fallocate(int(file.Fd()), unix.FALLOC_FL_ZERO_RANGE,
		headerSize, int64(width*height*BYTES_PER_PIXEL))
	if err != nil {
		return nil, err
	}

	mappedFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	pgm := PPMFile{
		W:          width,
		H:          height,
		Filename:   filename,
		File:       file,
		mappedFile: mappedFile,
		headerSize: headerSize,
	}
	return &pgm, nil
}

func OpenPPM(filename string) (*PPMFile, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// only bother to support the exact format that the create function writes
	var width, height int64
	n, err := fmt.Fscanf(file, "P6\n%d %d\n255\n", &width, &height)
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

	pgm := PPMFile{
		W:          int64(width),
		H:          int64(height),
		Filename:   filename,
		File:       file,
		mappedFile: mappedFile,
		headerSize: headerSize,
	}
	return &pgm, nil
}
