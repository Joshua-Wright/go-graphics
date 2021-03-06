#!/usr/bin/env bash
# generate_fixnums.sh


# min_words=2
# max_words=31

words="2 3 4 5 6 7 8 10 12 14 16 18 22 26 30"

for i in $words; do
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
echo "//go:generate bash generate_fixnums.sh" >> naive_fixnum_map.go
echo >> naive_fixnum_map.go
echo "import (" >> naive_fixnum_map.go
for i in $words; do
	echo "	\"github.com/joshua-wright/go-graphics/graphics/mandelbrot_fixed_point/naive_fixnum/generated/naive_fixnum_$i\"" >> naive_fixnum_map.go
done
echo "	\"github.com/joshua-wright/go-graphics/graphics/colormap\"" >> naive_fixnum_map.go
echo "	\"github.com/joshua-wright/go-graphics/graphics/memory_mapped\"" >> naive_fixnum_map.go
echo "	\"github.com/joshua-wright/go-graphics/graphics/per_pixel_image\"" >> naive_fixnum_map.go
echo ")" >> naive_fixnum_map.go
echo >> naive_fixnum_map.go

echo "type MandelbrotPerPixelConstructor func(width, height, maxIter int64," >> naive_fixnum_map.go
echo "	centerR, centerI, zoom, threshold string," >> naive_fixnum_map.go
echo "	Wrap, MaxVal float64, cmap colormap.ColorMap," >> naive_fixnum_map.go
echo "	OutImage *memory_mapped.PPMFile," >> naive_fixnum_map.go
echo "	OutIter *memory_mapped.Array2dFloat64) per_pixel_image.PixelFunction" >> naive_fixnum_map.go

echo >> naive_fixnum_map.go
echo "var FixnumConstructorWords = []uint{" >> naive_fixnum_map.go
for i in $words; do
	echo "	$i," >> naive_fixnum_map.go
done
echo "}" >> naive_fixnum_map.go

echo >> naive_fixnum_map.go
echo "var FixnumConstructors = []MandelbrotPerPixelConstructor{" >> naive_fixnum_map.go
for i in $words; do
	echo "	naive_fixnum_$i.NewMandelbrotPerPixel," >> naive_fixnum_map.go
done
echo "}" >> naive_fixnum_map.go
