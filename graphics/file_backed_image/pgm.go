package file_backed_image

import (
	"os"
	"syscall"
	"golang.org/x/sys/unix"
	"image"
	"image/color"
	"fmt"
)

type PPMFile struct {
	W, H       int64
	Filename   string
	File       *os.File
	headerSize int64
}

func (p *PPMFile) Close() error {
	return p.File.Close()
}

func (p *PPMFile) getOffset(x, y int) int64 {
	return p.headerSize + 3*(p.W*int64(y)+int64(x))
}

func (p *PPMFile) Set(x, y int, c color.Color) {
	offset := p.getOffset(x, y)

	rgbColor := color.RGBAModel.Convert(c).(color.RGBA)
	var rgb = [3]byte{
		rgbColor.R,
		rgbColor.G,
		rgbColor.B,
	}
	_, err := p.File.WriteAt(rgb[:], offset)
	if err != nil {
		panic(err)
	}
}

func (p *PPMFile) ColorModel() color.Model {
	return color.RGBAModel
}

func (p *PPMFile) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(p.W), int(p.H))
}

func (p *PPMFile) At(x, y int) color.Color {
	offset := p.getOffset(x, y)

	var rgb [3]byte
	_, err := p.File.ReadAt(rgb[:], offset)
	if err != nil {
		panic(err)
	}

	return color.RGBA{
		R: rgb[0],
		G: rgb[1],
		B: rgb[2],
		A: 255,
	}
}

// RGB format, 8 bytes per color
const BYTES_PER_PIXEL = 3

func CreatePPM(width, height int, filename string) (*PPMFile, error) {
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

	header_size := int64(bytes1 + bytes2)

	// allocate space in the file
	err = syscall.Fallocate(int(file.Fd()), unix.FALLOC_FL_ZERO_RANGE,
		header_size, header_size+int64(width*height*BYTES_PER_PIXEL))
	if err != nil {
		return nil, err
	}

	pgm := PPMFile{
		W:          int64(width),
		H:          int64(height),
		Filename:   filename,
		File:       file,
		headerSize: header_size,
	}
	return &pgm, nil
}
