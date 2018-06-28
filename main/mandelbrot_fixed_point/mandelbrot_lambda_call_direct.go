package main

import (
	"image"
	"github.com/joshua-wright/go-graphics/parallel"
	m "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	g "github.com/joshua-wright/go-graphics/graphics"
	l "github.com/joshua-wright/go-graphics/lambda"
	"math"
	"image/color"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"fmt"
	"bytes"
	"gopkg.in/cheggaaa/pb.v1"
	"runtime"
	"github.com/ncw/gmp"
	"encoding/gob"
	"github.com/aws/aws-lambda-go/events"
)

func main() {
	width := int64(64)
	height := int64(64)
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
	runtime.GOMAXPROCS(int(2 * width))
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
		err := gob.NewEncoder(cfgBuf).Encode(cfg)
		g.Die(err)

		resp, err := l.MandelbrotPixel(events.APIGatewayProxyRequest{
			Body: cfgBuf.String(),
		})

		if resp.StatusCode != 200 {
			fmt.Println(resp.Body)
			fmt.Println(j, "bad response")
			return
		}

		responseData := m.MandelbrotPixelRangeResponse{}
		gob.NewDecoder(bytes.NewBuffer([]byte(resp.Body))).Decode(&responseData)
		if err != nil {
			fmt.Println("failed to decode gob", j, err)
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
