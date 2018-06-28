package lambda

import (
	m "github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"encoding/gob"
	"bytes"
	"github.com/ncw/gmp"
)

func MandelbrotPixel(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var cfg m.MandelbrotPixelRangeConfig
	//err := json.Unmarshal([]byte(request.Body), &cfg)
	decoder := gob.NewDecoder(bytes.NewBuffer([]byte(request.Body)))
	err := decoder.Decode(&cfg)

	if err != nil {
		log.Println("bad input: ", request.Body)
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
		resp.Mag2 = make([]*gmp.Int, size)
	}

	for j0 := int64(0); j0 < jRange; j0++ {
		for i0 := int64(0); i0 < iRange; i0++ {
			i := i0 + cfg.Imin
			j := j0 + cfg.Jmin

			cr, ci := m.MandelbrotCoordinate(i, j, cfg.Width, cfg.Height, cfg.CenterR, cfg.CenterI, cfg.Zoom, cfg.BasePower2)
			iteration, val, mag2 := m.MandelbrotKernel(cr, ci, cfg.Threshold2, cfg.MaxIter, cfg.BasePower2)

			idx := iRange*j0 + i0
			if cfg.ReturnIteration {
				resp.Iteration[idx] = iteration
			}
			if cfg.ReturnVal {
				resp.Val[idx] = val
			}
			if cfg.ReturnMag2 {
				resp.Mag2[idx] = mag2
			}
		}
	}

	outBuffer := new(bytes.Buffer)
	err = gob.NewEncoder(outBuffer).Encode(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "failed: " + err.Error(),
			StatusCode: 500,
		}, err
	}

	respStr := outBuffer.String()
	return events.APIGatewayProxyResponse{
		Body:       string(respStr),
		StatusCode: 200,
	}, err
}
