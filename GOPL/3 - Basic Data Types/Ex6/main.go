package main

import (
	"image"
	"image/color"
	"image/png"
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
		y0 := float64(py)/height*(ymax-ymin) + ymin
		y1 := (float64(py)+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x0 := float64(px)/width*(xmax-xmin) + xmin
			x1 := (float64(px)+1)/width*(xmax-xmin) + xmin
			z1 := complex(x0, y0)
			z2 := complex(x0, y1)
			z3 := complex(x1, y0)
			z4 := complex(x1, y1)
			//
			colorr := getColorAverage([]complex128{z1, z2, z3, z4})
			// Image point (px, py) represents complex value z.
			img.Set(px, py, colorr)
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
			return color.Gray{255 - n*contrast}
		}
	}
	return color.Black
}

func getColorAverage(cplx []complex128) color.Color {
	var r, g, b, a float64
	for _, n := range cplx {
		cl := mandelbrot(n)
		rm, gm, bm, am := cl.RGBA()
		r += float64(rm) / float64(len(cplx))
		g += float64(gm) / float64(len(cplx))
		b += float64(bm) / float64(len(cplx))
		a += float64(am) / float64(len(cplx))
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
