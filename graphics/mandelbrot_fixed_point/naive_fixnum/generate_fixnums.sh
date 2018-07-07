#!/usr/bin/env bash
# generate_fixnums.sh


min_words=2
max_words=31

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


echo "package naive_fixnum" > naive_fixnum_map.go
echo >> naive_fixnum_map.go
echo "import (" >> naive_fixnum_map.go
for (( i = $min_words; i < $max_words; i++ )); do
	echo "	\"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_$i\"" >> naive_fixnum_map.go
done
echo ")" >> naive_fixnum_map.go
echo >> naive_fixnum_map.go
echo "var FixnumConstructorMap = map[uint64]MandelbrotPerPixelConstructor{" >> naive_fixnum_map.go
for (( i = $min_words; i < $max_words; i++ )); do
	echo "	$i : naive_fixnum_$i.NewMandelbrotPerPixel," >> naive_fixnum_map.go
done
echo "}" >> naive_fixnum_map.go

