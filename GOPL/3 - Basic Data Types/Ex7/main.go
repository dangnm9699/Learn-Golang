package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

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
	f, _ := os.Create("out.png")
	defer f.Close()
	png.Encode(f, img) // NOTE: ignoring errors
}

// f(z) = z^4 - 1
// delta(z) = - f(z) / f'(z)
// delta(z) = - (z^4 - 1) / (4*z^3)
// delta(z) = - 1/4(z - 1/z^3)
func mandelbrot(z complex128) color.Color {
	const iterations = 20
	const contrast = 15
	for n := uint8(0); n < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			if (math.Abs(real(z)-1) < 1e-6) && (math.Abs(imag(z)-0) < 1e-6) {
				return color.RGBA{255 - contrast*n, 0, 0, 255}
			}
			if (math.Abs(real(z)+1) < 1e-6) && (math.Abs(imag(z)-0) < 1e-6) {
				return color.RGBA{0, 255 - contrast*n, 0, 255}
			}
			if (math.Abs(real(z)-0) < 1e-6) && (math.Abs(imag(z)-1) < 1e-6) {
				return color.RGBA{0, 0, 255 - contrast*n, 255}
			}
			if (math.Abs(real(z)-0) < 1e-6) && math.Abs(imag(z)+1) < 1e-6 {
				return color.RGBA{255 - contrast*n, 255 - contrast*n, 0, 255}
			}
		}
	}
	return color.Black
}
