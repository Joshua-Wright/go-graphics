package main

import (
	"image"
	"github.com/joshua-wright/go-graphics/parallel"
	m "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"image/color"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"os"
	"fmt"
	"net/http"
	"bytes"
	"io"
	"gopkg.in/cheggaaa/pb.v1"
	"github.com/ncw/gmp"
	"encoding/gob"
	"encoding/base64"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	endpoint := os.Args[1]
	api_key := os.Args[2]
	width := int64(512)
	height := int64(512)
	maxIter := int64(512)
	wrap := 8.0
	basePower2 := uint(64)
	//basePower2 := uint(20)

	cmap := colormap.NewXyzInterpColormap(colormap.InfernoColorMap())

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	threshold2 := gmp.NewInt(32)
	threshold2.Lsh(threshold2, basePower2)
	threshold2.Mul(threshold2, threshold2)

	zoom := gmp.NewInt(1)
	centerR := gmp.NewInt(0)
	centerI := gmp.NewInt(0)

	bar := pb.StartNew(int(height))
	parallel.ParallelFor(0, int(height), func(j_ int) {
		j := int64(j_)
		defer bar.Increment()

		//if j != 0 {
		//	return
		//}

		cfg := m.MandelbrotPixelRangeConfig{
			Imin:            0,
			Imax:            width,
			Jmin:            j,
			Jmax:            j + 1,
			Width:           width,
			Height:          height,
			MaxIter:         maxIter,
			CenterR:         centerR,
			CenterI:         centerI,
			Zoom:            zoom,
			Threshold2:      threshold2,
			BasePower2:      basePower2,
			ReturnIteration: false,
			ReturnVal:       true,
			ReturnMag2:      false,
		}
		cfgBuf := new(bytes.Buffer)
		//err := gob.NewEncoder(cfgBuf).Encode(cfg)
		//err := gob.NewEncoder(base64.NewEncoder(base64.StdEncoding, cfgBuf)).Encode(cfg)
		err := gob.NewEncoder(cfgBuf).Encode(cfg)
		g.Die(err)
		//fmt.Println(cfgBuf.String())

		request, err := http.NewRequest("POST", endpoint, cfgBuf)
		g.Die(err)
		request.Header.Set("Content-Type", "application/octet-stream")
		request.Header.Set("x-api-key", api_key)
		client := &http.Client{}
		response, err := client.Do(request)
		g.Die(err)

		if response.StatusCode != 200 {
			io.Copy(os.Stdout, response.Body)
			fmt.Println()
			fmt.Println(j, "bad response")
			return
		}

		responseData := m.MandelbrotPixelRangeResponse{}
		gob.NewDecoder(base64.NewDecoder(base64.StdEncoding, response.Body)).Decode(&responseData)
		if err != nil {
			fmt.Println("failed to decode gob", j, err)
			return
		}
		if len(responseData.Val) != int(width) {
			fmt.Println("bad response, invalid length")
			spew.Dump(responseData)
			spew.Dump(response)
			io.Copy(os.Stdout, response.Body)
			panic(1)
			return
		}

		for i := int64(0); i < width; i++ {

			val := responseData.Val[i]

			var col color.Color
			if val >= 0 {
				val = math.Log2(val+1) / math.Log2(float64(maxIter)+1) * wrap
				val = math.Sin(val*2*math.Pi)/2.0 + 0.5
				col = cmap.GetColor(val)
			} else {
				col = color.Black
			}
			img.Set(int(i), int(j), col)
		}
	})
	bar.Finish()

	g.SaveAsPNG(img, g.ExecutableNamePng())
}
