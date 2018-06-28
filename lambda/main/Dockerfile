FROM golang:alpine

RUN apk update
RUN apk add alpine-sdk git ca-certificates gmp-dev zlib-dev

# docker run -it -v /home/j0sh/Documents/code/go-graphics/src:/go/src go-gmp-static /bin/sh
# cd /go/src/github.com/joshua-wright/go-graphics/lambda
# go build --ldflags '-linkmode external -extldflags "-static"' mandelbrot_fixed_point.go
