// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// Plan9
/**
Plan9 is a 256-color palette that partitions the 24-bit RGB space into 4×4×4 subdivision, with 4 shades in each subcube.
Compared to the WebSafe, the idea is to reduce the color resolution by dicing the color cube into fewer cells, and to
use the extra space to increase the intensity resolution. This results in 16 gray shades (4 gray subcubes with 4
samples in each), 13 shades of each primary and secondary color (3 subcubes with 4 samples plus black) and a reasonable
selection of colors covering the rest of the color cube. The advantage is better representation of continuous tones.

This palette was used in the Plan 9 Operating System, described at https://9p.io/magic/man2html/6/color
*/
var Plan9 = []color.Color{
	color.RGBA{A: 0xff},
	color.RGBA{B: 0x44, A: 0xff},
	color.RGBA{B: 0x88, A: 0xff},
	color.RGBA{B: 0xcc, A: 0xff},
	color.RGBA{G: 0x44, A: 0xff},
	color.RGBA{0x00, 0x44, 0x44, 0xff},
	color.RGBA{0x00, 0x44, 0x88, 0xff},
	color.RGBA{0x00, 0x44, 0xcc, 0xff},
	color.RGBA{0x00, 0x88, 0x00, 0xff},
	color.RGBA{0x00, 0x88, 0x44, 0xff},
	color.RGBA{0x00, 0x88, 0x88, 0xff},
	color.RGBA{0x00, 0x88, 0xcc, 0xff},
	color.RGBA{0x00, 0xcc, 0x00, 0xff},
	color.RGBA{0x00, 0xcc, 0x44, 0xff},
	color.RGBA{0x00, 0xcc, 0x88, 0xff},
	color.RGBA{0x00, 0xcc, 0xcc, 0xff},
	color.RGBA{0x00, 0xdd, 0xdd, 0xff},
	color.RGBA{0x11, 0x11, 0x11, 0xff},
	color.RGBA{0x00, 0x00, 0x55, 0xff},
	color.RGBA{0x00, 0x00, 0x99, 0xff},
	color.RGBA{0x00, 0x00, 0xdd, 0xff},
	color.RGBA{0x00, 0x55, 0x00, 0xff},
	color.RGBA{0x00, 0x55, 0x55, 0xff},
	color.RGBA{0x00, 0x4c, 0x99, 0xff},
	color.RGBA{0x00, 0x49, 0xdd, 0xff},
	color.RGBA{0x00, 0x99, 0x00, 0xff},
	color.RGBA{0x00, 0x99, 0x4c, 0xff},
	color.RGBA{0x00, 0x99, 0x99, 0xff},
	color.RGBA{0x00, 0x93, 0xdd, 0xff},
	color.RGBA{0x00, 0xdd, 0x00, 0xff},
	color.RGBA{0x00, 0xdd, 0x49, 0xff},
	color.RGBA{0x00, 0xdd, 0x93, 0xff},
	color.RGBA{0x00, 0xee, 0x9e, 0xff},
	color.RGBA{0x00, 0xee, 0xee, 0xff},
	color.RGBA{0x22, 0x22, 0x22, 0xff},
	color.RGBA{0x00, 0x00, 0x66, 0xff},
	color.RGBA{0x00, 0x00, 0xaa, 0xff},
	color.RGBA{0x00, 0x00, 0xee, 0xff},
	color.RGBA{0x00, 0x66, 0x00, 0xff},
	color.RGBA{0x00, 0x66, 0x66, 0xff},
	color.RGBA{0x00, 0x55, 0xaa, 0xff},
	color.RGBA{0x00, 0x4f, 0xee, 0xff},
	color.RGBA{0x00, 0xaa, 0x00, 0xff},
	color.RGBA{0x00, 0xaa, 0x55, 0xff},
	color.RGBA{0x00, 0xaa, 0xaa, 0xff},
	color.RGBA{0x00, 0x9e, 0xee, 0xff},
	color.RGBA{0x00, 0xee, 0x00, 0xff},
	color.RGBA{0x00, 0xee, 0x4f, 0xff},
	color.RGBA{0x00, 0xff, 0x55, 0xff},
	color.RGBA{0x00, 0xff, 0xaa, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0x33, 0x33, 0x33, 0xff},
	color.RGBA{0x00, 0x00, 0x77, 0xff},
	color.RGBA{0x00, 0x00, 0xbb, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0x77, 0x00, 0xff},
	color.RGBA{0x00, 0x77, 0x77, 0xff},
	color.RGBA{0x00, 0x5d, 0xbb, 0xff},
	color.RGBA{0x00, 0x55, 0xff, 0xff},
	color.RGBA{0x00, 0xbb, 0x00, 0xff},
	color.RGBA{0x00, 0xbb, 0x5d, 0xff},
	color.RGBA{0x00, 0xbb, 0xbb, 0xff},
	color.RGBA{0x00, 0xaa, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x44, 0x00, 0x44, 0xff},
	color.RGBA{0x44, 0x00, 0x88, 0xff},
	color.RGBA{0x44, 0x00, 0xcc, 0xff},
	color.RGBA{0x44, 0x44, 0x00, 0xff},
	color.RGBA{0x44, 0x44, 0x44, 0xff},
	color.RGBA{0x44, 0x44, 0x88, 0xff},
	color.RGBA{0x44, 0x44, 0xcc, 0xff},
	color.RGBA{0x44, 0x88, 0x00, 0xff},
	color.RGBA{0x44, 0x88, 0x44, 0xff},
	color.RGBA{0x44, 0x88, 0x88, 0xff},
	color.RGBA{0x44, 0x88, 0xcc, 0xff},
	color.RGBA{0x44, 0xcc, 0x00, 0xff},
	color.RGBA{0x44, 0xcc, 0x44, 0xff},
	color.RGBA{0x44, 0xcc, 0x88, 0xff},
	color.RGBA{0x44, 0xcc, 0xcc, 0xff},
	color.RGBA{0x44, 0x00, 0x00, 0xff},
	color.RGBA{0x55, 0x00, 0x00, 0xff},
	color.RGBA{0x55, 0x00, 0x55, 0xff},
	color.RGBA{0x4c, 0x00, 0x99, 0xff},
	color.RGBA{0x49, 0x00, 0xdd, 0xff},
	color.RGBA{0x55, 0x55, 0x00, 0xff},
	color.RGBA{0x55, 0x55, 0x55, 0xff},
	color.RGBA{0x4c, 0x4c, 0x99, 0xff},
	color.RGBA{0x49, 0x49, 0xdd, 0xff},
	color.RGBA{0x4c, 0x99, 0x00, 0xff},
	color.RGBA{0x4c, 0x99, 0x4c, 0xff},
	color.RGBA{0x4c, 0x99, 0x99, 0xff},
	color.RGBA{0x49, 0x93, 0xdd, 0xff},
	color.RGBA{0x49, 0xdd, 0x00, 0xff},
	color.RGBA{0x49, 0xdd, 0x49, 0xff},
	color.RGBA{0x49, 0xdd, 0x93, 0xff},
	color.RGBA{0x49, 0xdd, 0xdd, 0xff},
	color.RGBA{0x4f, 0xee, 0xee, 0xff},
	color.RGBA{0x66, 0x00, 0x00, 0xff},
	color.RGBA{0x66, 0x00, 0x66, 0xff},
	color.RGBA{0x55, 0x00, 0xaa, 0xff},
	color.RGBA{0x4f, 0x00, 0xee, 0xff},
	color.RGBA{0x66, 0x66, 0x00, 0xff},
	color.RGBA{0x66, 0x66, 0x66, 0xff},
	color.RGBA{0x55, 0x55, 0xaa, 0xff},
	color.RGBA{0x4f, 0x4f, 0xee, 0xff},
	color.RGBA{0x55, 0xaa, 0x00, 0xff},
	color.RGBA{0x55, 0xaa, 0x55, 0xff},
	color.RGBA{0x55, 0xaa, 0xaa, 0xff},
	color.RGBA{0x4f, 0x9e, 0xee, 0xff},
	color.RGBA{0x4f, 0xee, 0x00, 0xff},
	color.RGBA{0x4f, 0xee, 0x4f, 0xff},
	color.RGBA{0x4f, 0xee, 0x9e, 0xff},
	color.RGBA{0x55, 0xff, 0xaa, 0xff},
	color.RGBA{0x55, 0xff, 0xff, 0xff},
	color.RGBA{0x77, 0x00, 0x00, 0xff},
	color.RGBA{0x77, 0x00, 0x77, 0xff},
	color.RGBA{0x5d, 0x00, 0xbb, 0xff},
	color.RGBA{0x55, 0x00, 0xff, 0xff},
	color.RGBA{0x77, 0x77, 0x00, 0xff},
	color.RGBA{0x77, 0x77, 0x77, 0xff},
	color.RGBA{0x5d, 0x5d, 0xbb, 0xff},
	color.RGBA{0x55, 0x55, 0xff, 0xff},
	color.RGBA{0x5d, 0xbb, 0x00, 0xff},
	color.RGBA{0x5d, 0xbb, 0x5d, 0xff},
	color.RGBA{0x5d, 0xbb, 0xbb, 0xff},
	color.RGBA{0x55, 0xaa, 0xff, 0xff},
	color.RGBA{0x55, 0xff, 0x00, 0xff},
	color.RGBA{0x55, 0xff, 0x55, 0xff},
	color.RGBA{0x88, 0x00, 0x88, 0xff},
	color.RGBA{0x88, 0x00, 0xcc, 0xff},
	color.RGBA{0x88, 0x44, 0x00, 0xff},
	color.RGBA{0x88, 0x44, 0x44, 0xff},
	color.RGBA{0x88, 0x44, 0x88, 0xff},
	color.RGBA{0x88, 0x44, 0xcc, 0xff},
	color.RGBA{0x88, 0x88, 0x00, 0xff},
	color.RGBA{0x88, 0x88, 0x44, 0xff},
	color.RGBA{0x88, 0x88, 0x88, 0xff},
	color.RGBA{0x88, 0x88, 0xcc, 0xff},
	color.RGBA{0x88, 0xcc, 0x00, 0xff},
	color.RGBA{0x88, 0xcc, 0x44, 0xff},
	color.RGBA{0x88, 0xcc, 0x88, 0xff},
	color.RGBA{0x88, 0xcc, 0xcc, 0xff},
	color.RGBA{0x88, 0x00, 0x00, 0xff},
	color.RGBA{0x88, 0x00, 0x44, 0xff},
	color.RGBA{0x99, 0x00, 0x4c, 0xff},
	color.RGBA{0x99, 0x00, 0x99, 0xff},
	color.RGBA{0x93, 0x00, 0xdd, 0xff},
	color.RGBA{0x99, 0x4c, 0x00, 0xff},
	color.RGBA{0x99, 0x4c, 0x4c, 0xff},
	color.RGBA{0x99, 0x4c, 0x99, 0xff},
	color.RGBA{0x93, 0x49, 0xdd, 0xff},
	color.RGBA{0x99, 0x99, 0x00, 0xff},
	color.RGBA{0x99, 0x99, 0x4c, 0xff},
	color.RGBA{0x99, 0x99, 0x99, 0xff},
	color.RGBA{0x93, 0x93, 0xdd, 0xff},
	color.RGBA{0x93, 0xdd, 0x00, 0xff},
	color.RGBA{0x93, 0xdd, 0x49, 0xff},
	color.RGBA{0x93, 0xdd, 0x93, 0xff},
	color.RGBA{0x93, 0xdd, 0xdd, 0xff},
	color.RGBA{0x99, 0x00, 0x00, 0xff},
	color.RGBA{0xaa, 0x00, 0x00, 0xff},
	color.RGBA{0xaa, 0x00, 0x55, 0xff},
	color.RGBA{0xaa, 0x00, 0xaa, 0xff},
	color.RGBA{0x9e, 0x00, 0xee, 0xff},
	color.RGBA{0xaa, 0x55, 0x00, 0xff},
	color.RGBA{0xaa, 0x55, 0x55, 0xff},
	color.RGBA{0xaa, 0x55, 0xaa, 0xff},
	color.RGBA{0x9e, 0x4f, 0xee, 0xff},
	color.RGBA{0xaa, 0xaa, 0x00, 0xff},
	color.RGBA{0xaa, 0xaa, 0x55, 0xff},
	color.RGBA{0xaa, 0xaa, 0xaa, 0xff},
	color.RGBA{0x9e, 0x9e, 0xee, 0xff},
	color.RGBA{0x9e, 0xee, 0x00, 0xff},
	color.RGBA{0x9e, 0xee, 0x4f, 0xff},
	color.RGBA{0x9e, 0xee, 0x9e, 0xff},
	color.RGBA{0x9e, 0xee, 0xee, 0xff},
	color.RGBA{0xaa, 0xff, 0xff, 0xff},
	color.RGBA{0xbb, 0x00, 0x00, 0xff},
	color.RGBA{0xbb, 0x00, 0x5d, 0xff},
	color.RGBA{0xbb, 0x00, 0xbb, 0xff},
	color.RGBA{0xaa, 0x00, 0xff, 0xff},
	color.RGBA{0xbb, 0x5d, 0x00, 0xff},
	color.RGBA{0xbb, 0x5d, 0x5d, 0xff},
	color.RGBA{0xbb, 0x5d, 0xbb, 0xff},
	color.RGBA{0xaa, 0x55, 0xff, 0xff},
	color.RGBA{0xbb, 0xbb, 0x00, 0xff},
	color.RGBA{0xbb, 0xbb, 0x5d, 0xff},
	color.RGBA{0xbb, 0xbb, 0xbb, 0xff},
	color.RGBA{0xaa, 0xaa, 0xff, 0xff},
	color.RGBA{0xaa, 0xff, 0x00, 0xff},
	color.RGBA{0xaa, 0xff, 0x55, 0xff},
	color.RGBA{0xaa, 0xff, 0xaa, 0xff},
	color.RGBA{0xcc, 0x00, 0xcc, 0xff},
	color.RGBA{0xcc, 0x44, 0x00, 0xff},
	color.RGBA{0xcc, 0x44, 0x44, 0xff},
	color.RGBA{0xcc, 0x44, 0x88, 0xff},
	color.RGBA{0xcc, 0x44, 0xcc, 0xff},
	color.RGBA{0xcc, 0x88, 0x00, 0xff},
	color.RGBA{0xcc, 0x88, 0x44, 0xff},
	color.RGBA{0xcc, 0x88, 0x88, 0xff},
	color.RGBA{0xcc, 0x88, 0xcc, 0xff},
	color.RGBA{0xcc, 0xcc, 0x00, 0xff},
	color.RGBA{0xcc, 0xcc, 0x44, 0xff},
	color.RGBA{0xcc, 0xcc, 0x88, 0xff},
	color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	color.RGBA{0xcc, 0x00, 0x00, 0xff},
	color.RGBA{0xcc, 0x00, 0x44, 0xff},
	color.RGBA{0xcc, 0x00, 0x88, 0xff},
	color.RGBA{0xdd, 0x00, 0x93, 0xff},
	color.RGBA{0xdd, 0x00, 0xdd, 0xff},
	color.RGBA{0xdd, 0x49, 0x00, 0xff},
	color.RGBA{0xdd, 0x49, 0x49, 0xff},
	color.RGBA{0xdd, 0x49, 0x93, 0xff},
	color.RGBA{0xdd, 0x49, 0xdd, 0xff},
	color.RGBA{0xdd, 0x93, 0x00, 0xff},
	color.RGBA{0xdd, 0x93, 0x49, 0xff},
	color.RGBA{0xdd, 0x93, 0x93, 0xff},
	color.RGBA{0xdd, 0x93, 0xdd, 0xff},
	color.RGBA{0xdd, 0xdd, 0x00, 0xff},
	color.RGBA{0xdd, 0xdd, 0x49, 0xff},
	color.RGBA{0xdd, 0xdd, 0x93, 0xff},
	color.RGBA{0xdd, 0xdd, 0xdd, 0xff},
	color.RGBA{0xdd, 0x00, 0x00, 0xff},
	color.RGBA{0xdd, 0x00, 0x49, 0xff},
	color.RGBA{0xee, 0x00, 0x4f, 0xff},
	color.RGBA{0xee, 0x00, 0x9e, 0xff},
	color.RGBA{0xee, 0x00, 0xee, 0xff},
	color.RGBA{0xee, 0x4f, 0x00, 0xff},
	color.RGBA{0xee, 0x4f, 0x4f, 0xff},
	color.RGBA{0xee, 0x4f, 0x9e, 0xff},
	color.RGBA{0xee, 0x4f, 0xee, 0xff},
	color.RGBA{0xee, 0x9e, 0x00, 0xff},
	color.RGBA{0xee, 0x9e, 0x4f, 0xff},
	color.RGBA{0xee, 0x9e, 0x9e, 0xff},
	color.RGBA{0xee, 0x9e, 0xee, 0xff},
	color.RGBA{0xee, 0xee, 0x00, 0xff},
	color.RGBA{0xee, 0xee, 0x4f, 0xff},
	color.RGBA{0xee, 0xee, 0x9e, 0xff},
	color.RGBA{0xee, 0xee, 0xee, 0xff},
	color.RGBA{0xee, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x55, 0xff},
	color.RGBA{0xff, 0x00, 0xaa, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x55, 0x00, 0xff},
	color.RGBA{0xff, 0x55, 0x55, 0xff},
	color.RGBA{0xff, 0x55, 0xaa, 0xff},
	color.RGBA{0xff, 0x55, 0xff, 0xff},
	color.RGBA{0xff, 0xaa, 0x00, 0xff},
	color.RGBA{0xff, 0xaa, 0x55, 0xff},
	color.RGBA{0xff, 0xaa, 0xaa, 0xff},
	color.RGBA{0xff, 0xaa, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x55, 0xff},
	color.RGBA{0xff, 0xff, 0xaa, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{Y: 255 - contrast*n}
			return Plan9[255-contrast*n]
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{Y: 192, Cb: blue, Cr: red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{Y: 128, Cb: blue, Cr: red}
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{Y: 255 - contrast*i}
		}
	}
	return color.Black
}