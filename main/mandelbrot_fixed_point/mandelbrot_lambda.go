package main

import (
	"image"
	"github.com/joshua-wright/go-graphics/parallel"
	m "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	g "github.com/joshua-wright/go-graphics/graphics"
	"math"
	"image/color"
	"github.com/joshua-wright/go-graphics/graphics/colormap"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"github.com/ncw/gmp"
	"fmt"
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
	threshold2str := "1024.0"

	//// need to use a large threshold for the smooth coloring to work
	//threshold2 := gmp.NewInt(32)
	//threshold2.Lsh(threshold2, basePower2)
	//threshold2.Mul(threshold2, threshold2)
	//zoom := gmp.NewInt(1)
	//centerR := gmp.NewInt(0)
	//centerI := gmp.NewInt(0)
	//zoom := gmp.NewInt(28047)
	//centerR := m.ParseFixnum("-0.74364085", basePower2)
	//centerI := m.ParseFixnum("0.13182733", basePower2)

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
			Threshold2:      threshold2str,
			BasePower2:      basePower2,
			ReturnIteration: false,
			ReturnVal:       true,
			ReturnMag2:      false,
		}
		jsonCfg, err := json.Marshal(cfg)
		g.Die(err)

		//request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonCfg))
		//g.Die(err)
		//request.Header.Set("x-api-key", api_key)
		//client := &http.Client{}
		//response, err := client.Do(request)
		//g.Die(err)
		//
		//if response.StatusCode != 200 {
		//	io.Copy(os.Stdout, response.Body)
		//	fmt.Println()
		//	fmt.Println(j, "bad response")
		//	return
		//	//panic("bad server response")
		//}
		//
		//responseData := m.MandelbrotPixelRangeResponse{}
		//responseBuf := new(bytes.Buffer)
		//_, err = responseBuf.ReadFrom(response.Body)
		//if err != nil {
		//	fmt.Println(j, err)
		//	return
		//}
		//json.Unmarshal(responseBuf.Bytes(), &responseData)

		_ = endpoint
		_ = api_key
		req := events.APIGatewayProxyRequest{
			Body: string(jsonCfg),
		}
		resp, err := MandelbrotPixel(req)
		if err != nil {
			fmt.Println(j, err)
			return
		}

		responseData := m.MandelbrotPixelRangeResponse{}
		json.Unmarshal([]byte(resp.Body), &responseData)

		for i := int64(0); i < width; i++ {

			//cr, ci := m.MandelbrotCoordinate(i, j, width, height, centerR, centerI, zoom, basePower2)
			//_, val, _ := m.MandelbrotKernel(cr, ci, threshold2, maxIter, basePower2)
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
			//}
		}
		bar.Increment()
	})
	bar.Finish()

	g.SaveAsPNG(img, g.ExecutableNamePng())
}

func MandelbrotPixel(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var cfg m.MandelbrotPixelRangeConfig
	err := json.Unmarshal([]byte(request.Body), &cfg)
	//decoder := gob.NewDecoder(bytes.NewBuffer([]byte(request.Body)))
	//err := decoder.Decode(&cfg)

	if err != nil {
		log.Println("bad json: ", request.Body)
		return events.APIGatewayProxyResponse{
			Body:       "failed: " + err.Error(),
			StatusCode: 500,
		}, err
	}

	centerR, err := m.ParseFixnumSafe(cfg.CenterR, cfg.BasePower2)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "failed: " + err.Error(),
			StatusCode: 500,
		}, err
	}
	centerI, err := m.ParseFixnumSafe(cfg.CenterI, cfg.BasePower2)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "failed: " + err.Error(),
			StatusCode: 500,
		}, err
	}
	zoom := new(gmp.Int)
	zoom, success := zoom.SetString(cfg.Zoom, 10)
	if !success {
		return events.APIGatewayProxyResponse{
			Body:       "failed to parse zoom",
			StatusCode: 500,
		}, err
	}
	threshold2, err := m.ParseFixnumSafe(cfg.Threshold2, cfg.BasePower2)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "failed: " + err.Error(),
			StatusCode: 500,
		}, err
	}

	iRange := cfg.Imax - cfg.Imin
	jRange := cfg.Jmax - cfg.Jmin
	size := iRange * jRange

	resp := m.MandelbrotPixelRangeResponse{}
	if cfg.ReturnIteration {
		resp.Iteration = make([]int64, size)
	}
	if cfg.ReturnVal {
		resp.Val = make([]float64, size)
	}
	if cfg.ReturnMag2 {
		resp.Mag2 = make([]string, size)
	}

	for i0 := int64(0); i0 < iRange; i0++ {
		for j0 := int64(0); j0 < jRange; j0++ {
			i := i0 + cfg.Imin
			j := j0 + cfg.Jmin

			cr, ci := m.MandelbrotCoordinate(i, j, cfg.Width, cfg.Height, centerR, centerI, zoom, cfg.BasePower2)
			iteration, val, mag2 := m.MandelbrotKernel(cr, ci, threshold2, cfg.MaxIter, cfg.BasePower2)

			idx := iRange*j0 + i0
			if cfg.ReturnIteration {
				resp.Iteration[idx] = iteration
			}
			if cfg.ReturnVal {
				resp.Val[idx] = val
			}
			if cfg.ReturnMag2 {
				resp.Mag2[idx] = mag2.String()
			}
		}
	}

	respStr, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "failed: " + err.Error(),
			StatusCode: 500,
		}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       string(respStr),
		StatusCode: 200,
	}, err
}
