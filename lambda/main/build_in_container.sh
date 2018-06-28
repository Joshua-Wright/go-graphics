#!/usr/bin/env bash

docker run -it -v /home/j0sh/Documents/code/go-graphics/src:/go/src go-gmp-static /bin/sh -c "cd /go/src/github.com/joshua-wright/go-graphics/lambda && go build --ldflags '-linkmode external -extldflags \"-static\"' mandelbrot_fixed_point.go"
zip mandelbrot_fixed_point.zip mandelbrot_fixed_point
