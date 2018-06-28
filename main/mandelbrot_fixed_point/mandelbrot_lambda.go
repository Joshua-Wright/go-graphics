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
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"io"
	"gopkg.in/cheggaaa/pb.v1"
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

	centerRstr := "0.0"
	centerIstr := "0.0"
	zoomstr := "1"
	thresholdstr := "32.0"

	bar := pb.StartNew(int(height))
	parallel.ParallelFor(0, int(height), func(j_ int) {
		j := int64(j_)

		cfg := m.MandelbrotPixelRangeConfig{
			Imin:            0,
			Imax:            width,
			Jmin:            j,
			Jmax:            j + 1,
			Width:           width,
			Height:          height,
			MaxIter:         maxIter,
			CenterR:         centerRstr,
			CenterI:         centerIstr,
			Zoom:            zoomstr,
			Threshold:       thresholdstr,
			BasePower2:      basePower2,
			ReturnIteration: false,
			ReturnVal:       true,
			ReturnMag2:      false,
		}
		jsonCfg, err := json.Marshal(cfg)
		g.Die(err)

		request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonCfg))
		g.Die(err)
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
		responseBuf := new(bytes.Buffer)
		_, err = responseBuf.ReadFrom(response.Body)
		if err != nil {
			fmt.Println(j, err)
			return
		}
		err = json.Unmarshal(responseBuf.Bytes(), &responseData)
		if err != nil {
			fmt.Println(j, err)
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
		bar.Increment()
	})
	bar.Finish()

	g.SaveAsPNG(img, g.ExecutableNamePng())
}
