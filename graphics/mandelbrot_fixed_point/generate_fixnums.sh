#!/usr/bin/env bash
# generate_fixnums.sh


for (( i = 2; i < 64; i++ )); do
	mkdir naive-fixnum-$i
	cp -t naive-fixnum-$i fixnum.go mandelbrot.go bench_test.go fixnum_test.go
	echo "package naive-fixnum-$i" >> naive-fixnum-$i/fpwords.go
	echo "const FpWords = 4" >> naive-fixnum-$i/fpwords.go
done