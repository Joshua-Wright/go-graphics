package main

import (
	l "github.com/joshua-wright/go-graphics/lambda"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(l.MandelbrotPixel)
}
