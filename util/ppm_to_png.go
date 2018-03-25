package main

import (
	"os"
	"strings"
	"path/filepath"
	"image/png"
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
)

func main() {
	filename := os.Args[1]
	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".png"

	img, err := memory_mapped.OpenPPM(filename)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(newFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
