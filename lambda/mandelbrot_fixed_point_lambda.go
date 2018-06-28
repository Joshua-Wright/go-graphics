package lambda

import (
	m "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"
	"log"
	"github.com/ncw/gmp"
	"fmt"
)

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
	threshold2, err := m.ParseFixnumSafe(cfg.Threshold, cfg.BasePower2)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "failed: " + err.Error(),
			StatusCode: 500,
		}, err
	}
	threshold2.Mul(threshold2, threshold2)

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

	for j0 := int64(0); j0 < jRange; j0++ {
		for i0 := int64(0); i0 < iRange; i0++ {
			i := i0 + cfg.Imin
			j := j0 + cfg.Jmin

			cr, ci := m.MandelbrotCoordinate(i, j, cfg.Width, cfg.Height, centerR, centerI, zoom, cfg.BasePower2)
			iteration, val, mag2 := m.MandelbrotKernel(cr, ci, threshold2, cfg.MaxIter, cfg.BasePower2)
			//fmt.Println(i0, j0, i, j,
			//	cfg.Width, cfg.Height,
			//	centerR.String(), centerI.String(), threshold2, zoom.String(),
			//	cfg.MaxIter, cfg.BasePower2)
			fmt.Println("lambda_tag1", i, j,
				cfg.Width, cfg.Height,
				centerR.String(), centerI.String(), zoom.String(), cfg.BasePower2)
			fmt.Println("lambda_tag2", i, j,
				cr.String(), ci.String(), threshold2, cfg.BasePower2)
			fmt.Println("lambda_tag3", i, j, iteration, val, mag2.String())

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
