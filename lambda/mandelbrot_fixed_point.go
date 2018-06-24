package main

import (
	"github.com/aws/aws-lambda-go/events"
	m "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	"encoding/json"
	"github.com/ncw/gmp"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
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

	resp := m.MandelbrotSPixelRangeResponse{}
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

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(MandelbrotPixel)
	//resp, err := MandelbrotPixel(events.APIGatewayProxyRequest{
	//	//Body:`{"i": 0,"j": 0,"width": 1024,"height": 1024,"maxIter": 5120,"centerR": "-0.74364085","centerI": "0.13182733","zoom": "120188","threshold2": "32.0","basePower2": 128}`,
	//	//Body: `{"imin": 0,"imax": 64,"jmin": 0,"jmax": 1,"width": 64,"height": 64,"maxIter": 5120,"centerR": "-0.74364085","centerI": "0.13182733","zoom": "120188","threshold2": "1024.0","basePower2": 128,"returnIteration": false,"returnVal": true,"returnMag2": false}`,
	//	Body: `{"imin": 0,"imax": 64,"jmin": 0,"jmax": 1,"width": 64,"height": 64,"maxIter": 5120,"centerR": "0.0","centerI": "0.0","zoom": "1","threshold2": "1024.0","basePower2": 128,"returnIteration": false,"returnVal": true,"returnMag2": false}`,
	//})
	//fmt.Println(resp, err)
}
