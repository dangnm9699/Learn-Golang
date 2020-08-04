package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	width, height = 1024, 1024
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}
	var (
		xmin float64
		ymin float64
		xmax float64
		ymax float64
		x    float64 = 0
		y    float64 = 0
		zoom float64 = 0
	)
	if len(r.Form["x"]) >= 1 {
		x, _ = strconv.ParseFloat(r.Form["x"][0], 64)
	}
	if len(r.Form["y"]) >= 1 {
		y, _ = strconv.ParseFloat(r.Form["y"][0], 64)
	}
	if len(r.Form["zoom"]) >= 1 {
		zoom, _ = strconv.ParseFloat(r.Form["zoom"][0], 64)
	}
	dif := math.Pow(2, (1 - zoom))
	xmin, xmax = x-dif, x+dif
	ymin, ymax = y-dif, y+dif
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot64(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
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
