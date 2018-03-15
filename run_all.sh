#!/usr/bin/env bash

mains=(
	dragon_curve
	dragon_curve_pixel_turtle
	golden_dragon
	main
	sierpinski_arrowhead_curve
	sierpinski_arrowhead_depth
	sierpinski_arrowhead_smoothed_range
	sierpinski_arrowhead_smoothing
	sierpinski
	spirograph
	test_bspline
	chua_circuit
	logistic_map/logistic_map
	# slow:
	# dragon_curve_wallpaper
)

mkdir -p build
cd build

echo > run_log.txt
for main in ${mains[@]}; do
	echo $main
	if ! go build ../main/${main}.go; then
		echo "${main} failed to build"
		echo "${main} failed to build" >> run_log.txt
		continue
	fi
	if ! "./$(basename ${main})"; then
		echo "${main} failed to run"
		echo "${main} failed to run" >> run_log.txt
		continue
	fi
done
