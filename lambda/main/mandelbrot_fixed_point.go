package main

import (
	l "github.com/joshua-wright/go-graphics/lambda"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(l.MandelbrotPixel)
	//resp, err := MandelbrotPixel(events.APIGatewayProxyRequest{
	//	//Body:`{"i": 0,"j": 0,"width": 1024,"height": 1024,"maxIter": 5120,"centerR": "-0.74364085","centerI": "0.13182733","zoom": "120188","threshold2": "32.0","basePower2": 128}`,
	//	//Body: `{"imin": 0,"imax": 64,"jmin": 0,"jmax": 1,"width": 64,"height": 64,"maxIter": 5120,"centerR": "-0.74364085","centerI": "0.13182733","zoom": "120188","threshold2": "1024.0","basePower2": 128,"returnIteration": false,"returnVal": true,"returnMag2": false}`,
	//	Body: `{"imin": 0,"imax": 64,"jmin": 0,"jmax": 1,"width": 64,"height": 64,"maxIter": 5120,"centerR": "0.0","centerI": "0.0","zoom": "1","threshold2": "1024.0","basePower2": 128,"returnIteration": false,"returnVal": true,"returnMag2": false}`,
	//})
	//fmt.Println(resp, err)
}
