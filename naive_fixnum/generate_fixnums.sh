#!/usr/bin/env bash
# generate_fixnums.sh


min_words=2
max_words=64

for (( i = $min_words; i < $max_words; i++ )); do
	mkdir -p generated/naive_fixnum_$i
	cp _fixnum.go      generated/naive_fixnum_$i/fixnum.go
	cp _mandelbrot.go  generated/naive_fixnum_$i/mandelbrot.go
	cp _bench_test.go  generated/naive_fixnum_$i/bench_test.go
	cp _fixnum_test.go generated/naive_fixnum_$i/fixnum_test.go
	sed -i "s/package naive_fixnum/package naive_fixnum_$i/g" generated/naive_fixnum_$i/*.go
	echo "package naive_fixnum_$i" > generated/naive_fixnum_$i/fpwords.go
	echo "const FpWords = 4" >> generated/naive_fixnum_$i/fpwords.go
done

