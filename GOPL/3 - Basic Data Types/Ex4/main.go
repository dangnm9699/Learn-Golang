package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

var (
	width   int     // width size in pixels
	height  int     // height size in pixels
	cells   int     // number of grid cells
	xyrange float64 // axis ranges (-xyrange..+xyrange)
	xyscale float64 // pixels per x or y unit
	zscale  float64 // pixels per z unit
	angle   float64 // angle of x, y axes (=30°)
	sin30   float64
	cos30   float64
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/display", display)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "This is index page.\n<a href=\"/display\">Display</a>")
}

func display(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	if len(r.Form["width"]) < 1 {
		width = 600
	} else {
		width, _ = strconv.Atoi(r.Form["width"][0])
	}

	if len(r.Form["cells"]) < 1 {
		cells = 100
	} else {
		cells, _ = strconv.Atoi(r.Form["cells"][0])
	}

	if len(r.Form["height"]) < 1 {
		height = 320
	} else {
		height, _ = strconv.Atoi(r.Form["height"][0])
	}

	if len(r.Form["xyrange"]) < 1 {
		xyrange = 30.0
	} else {
		xyrange, _ = strconv.ParseFloat(r.Form["xyrange"][0], 64)
	}

	xyscale = float64(width) / 2 / xyrange
	zscale = float64(height) * 0.4
	angle = math.Pi / 6
	sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
	draw(w, width, height, cells, xyrange)
}

func draw(w io.Writer, width int, height int, cells int, xyrange float64) {
	var graph string
	graph += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	peak, valley := getPeakAndValley()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			z := getHeight(i, j)
			color := setColor(z, peak, valley)
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				graph += fmt.Sprintf("\t<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s;'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	graph += fmt.Sprint("</svg>")
	f, _ := os.Create("file.svg")
	defer f.Close()
	_, error := f.WriteString(graph)
	if error != nil {
		return
	}
	fmt.Fprintf(w, graph)
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// isFinite check if f is finite or non-finite
func isFinite(f float64) bool {
	if math.IsInf(f, 0) {
		return false
	}
	if math.IsNaN(f) {
		return false
	}
	return true
}

func getHeight(i, j int) float64 {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	// Compute surface height z.
	z := f(x, y)
	return z
}

func getPeakAndValley() (float64, float64) {
	peak := math.NaN()
	valley := math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := getHeight(i, j)
			if isFinite(z) {
				if math.IsNaN(peak) || z > peak {
					peak = z
				}
				if math.IsNaN(valley) || z < valley {
					valley = z
				}
			}
		}
	}
	return peak, valley
}

func setColor(height, peak, valley float64) string {
	var diff int
	if peak == valley {
		diff = 178
	} else {
		diff = int((height - valley) / (peak - valley) * 255)
	}
	red := fmt.Sprintf("%02x", diff)
	blue := fmt.Sprintf("%02x", 255-diff)
	return fmt.Sprintf("#%s00%s", red, blue)
}
