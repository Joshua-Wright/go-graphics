#!/usr/bin/env bash

mains=(
	dragon_curve
	dragon_curve_pixel_turtle
	dragon_curve_wallpaper
	golden_dragon
	main
	sierpinski_arrowhead_curve
	sierpinski_arrowhead_depth
	sierpinski_arrowhead_smoothed_range
	sierpinski_arrowhead_smoothing
	sierpinski
	test_bspline
)

mkdir -p build/bin
cd build

echo > run_log.txt
for main in ${mains[@]}; do
	echo $main
	cd bin
	if ! go build "../../main/${main}.go"; then
		echo "${main} failed to build"
		echo "${main} failed to build" >> run_log.txt
		continue
	fi
	cd ..
	if ! "./bin/${main}"; then
		echo "${main} failed to run"
		echo "${main} failed to run" >> run_log.txt
		continue
	fi
done
