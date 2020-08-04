package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = []color.Color{color.RGBA{255, 99, 71, 255},
	color.RGBA{224, 223, 213, 255},
	color.RGBA{73, 255, 68, 255},
	color.RGBA{241, 231, 193, 255},
	color.RGBA{43, 117, 255, 255},
	color.RGBA{209, 240, 182, 255}}

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
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 25
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[int(n)%len(palette)]
		}
	}
	return color.Black
}
