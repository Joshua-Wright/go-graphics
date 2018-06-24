package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	m "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	"encoding/json"
	"github.com/ncw/gmp"
	"log"
)

func MandelbrotPixel(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var cfg m.MandelbrotSinglePixelConfig
	err := json.Unmarshal([]byte(request.Body), &cfg)

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

	cr, ci := m.MandelbrotCoordinate(cfg.I, cfg.J, cfg.Width, cfg.Height, centerR, centerI, zoom, cfg.BasePower2)

	iteration, val, mag2 := m.MandelbrotKernel(cr, ci, threshold2, cfg.MaxIter, cfg.BasePower2)

	resp := m.MandelbrotSinglePixelResponse{
		I:         cfg.I,
		J:         cfg.J,
		Iteration: iteration,
		Val:       val,
		Mag2:      mag2.String(),
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
}
