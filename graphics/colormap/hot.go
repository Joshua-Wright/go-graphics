package colormap

import "image/color"

func HotColormap(x float64) color.RGBA {
	/* expects 0 <= x <= 1 */
	d := x * 255.0;
	pix := color.RGBA{};
	/* red */
	if (d > 94) {
		pix.R = 0xff;
	} else {
		pix.R = uint8(51.0 * d / 19.0);
	}

	/* green */
	if (d > 190) {
		pix.G = 0xff;
	} else if (d > 95) {
		pix.G = uint8(85.0*d/32.0 - 8075.0/32.0);
	} else {
		pix.G = 0;
	}

	/* blue */
	if (d > 191) {
		pix.B = uint8(255.0*d/64.0 - 48705.0/64.0);
	} else {
		pix.B = 0;
	}
	return pix;
}
