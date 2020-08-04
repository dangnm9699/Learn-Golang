package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
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
			// img.Set(px, py, mandelbrot(z))
			img.Set(px, py, mandelbrot64(z))
			// img.Set(px, py, mandelbrotBigFloat(z))
			// img.Set(px, py, mandelbrotBigRat(z))
		}
	}
	f, _ := os.Create("out.png")
	defer f.Close()
	png.Encode(f, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - n*contrast}
		}
	}
	return color.Black
}

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - n*contrast}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	realZ := (&big.Float{}).SetFloat64(real(z))
	imagZ := (&big.Float{}).SetFloat64(imag(z))
	var realV, imagV = &big.Float{}, &big.Float{}
	for n := uint8(0); n < iterations; n++ {
		fmt.Printf("...%d\n", n)
		realV2, imagV2 := &big.Float{}, &big.Float{}
		realV2.Mul(realV, realV).Sub(realV2, (&big.Float{}).Mul(imagV, imagV)).Add(realV2, realZ)
		imagV2.Mul(realV, imagV).Mul(imagV2, big.NewFloat(2)).Add(imagV2, imagZ)
		realV, imagV = realV2, imagV2
		cmplxAbsV := &big.Float{}
		cmplxAbsV.Mul(realV, realV).Add(cmplxAbsV, (&big.Float{}).Mul(imagV, imagV))
		if cmplxAbsV.Sqrt(cmplxAbsV).Cmp(big.NewFloat(2)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotBigRat(z complex128) color.Color {
	const iterations = 15
	const contrast = 15
	realZ := (&big.Rat{}).SetFloat64(real(z))
	imagZ := (&big.Rat{}).SetFloat64(imag(z))
	var realV, imagV = &big.Rat{}, &big.Rat{}
	for n := uint8(0); n < iterations; n++ {
		fmt.Printf("...%d\n", n)
		realV2, imagV2 := &big.Rat{}, &big.Rat{}
		realV2.Mul(realV, realV).Sub(realV2, (&big.Rat{}).Mul(imagV, imagV)).Add(realV2, realZ)
		imagV2.Mul(realV, imagV).Mul(imagV2, big.NewRat(2, 1)).Add(imagV2, imagZ)
		realV, imagV = realV2, imagV2
		cmplxAbsV := &big.Rat{}
		cmplxAbsV.Mul(realV, realV).Add(cmplxAbsV, (&big.Rat{}).Mul(imagV, imagV))
		if cmplxAbsV.Cmp(big.NewRat(4, 1)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
