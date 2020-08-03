package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"net/http"
	"sync"

	// "math/rand"
	// "os"
	"strconv"
)

var mu sync.Mutex
var count int

func main() {
	// routers
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		var (
			cycles  int
			res     float64
			size    int
			nframes int
			delay   int
		)
		// number of complete x oscillator revolutions
		if len(r.Form["cycles"]) < 1 {
			cycles = 5
		} else {
			cycles, _ = strconv.Atoi(r.Form["cycles"][0])
		}
		// angular resolution
		if len(r.Form["res"]) < 1 {
			res = 0.001
		} else {
			res, _ = strconv.ParseFloat(r.Form["res"][0], 64)
		}
		// image canvas covers [-size..+size]
		if len(r.Form["size"]) < 1 {
			size = 100
		} else {
			size, _ = strconv.Atoi(r.Form["size"][0])
		}
		// number of animation frames
		if len(r.Form["nframes"]) < 1 {
			nframes = 64
		} else {
			nframes, _ = strconv.Atoi(r.Form["nframes"][0])
		}
		// delay between frames in 10ms units
		if len(r.Form["delay"]) < 1 {
			delay = 8
		} else {
			delay, _ = strconv.Atoi(r.Form["delay"][0])
		}
		lissajous(w, cycles, res, size, nframes, delay)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(w io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := 0.3101 * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	palette := make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{0, 0, 0, 255})
	for i := 0; i < nframes; i++ {
		ratio := float64(i+1) / float64(nframes)
		col := color.RGBA{uint8(math.Abs(255*ratio - 0)), uint8(math.Abs(255*ratio - 85)), uint8(math.Abs(255*ratio - 170)), 255}
		palette = append(palette, col)
	}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Cos(t)
			y := math.Cos(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8((i+1)%nframes))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
