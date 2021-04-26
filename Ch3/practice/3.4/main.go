package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	blue          = "#0000ff"
	red           = "#ff0000"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	//@example: localhost:8000?colorH=00ff00&colorL=ff00ff
	http.HandleFunc("/", handle)
	log.Println(http.ListenAndServe("localhost:8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	var realWidth, realHeight, realColorH, realColorL = width, height, red, blue

	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}
	if r.Form["width"] != nil {
		realWidth, _ = strconv.Atoi(r.Form["width"][0])
	}
	if r.Form["height"] != nil {
		realHeight, _ = strconv.Atoi(r.Form["height"][0])
	}
	if r.Form["colorH"] != nil {
		realColorH = "#" + r.Form["colorH"][0]
	}
	if r.Form["colorL"] != nil {
		realColorL = "#" + r.Form["colorL"][0]
	}
	draw(w, realWidth, realHeight, realColorH, realColorL)
}

func draw(out io.Writer, width int, height int, colorH string, colorL string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, color := corner(i+1, j, colorH, colorL)
			bx, by, _ := corner(i, j, colorH, colorL)
			cx, cy, _ := corner(i, j+1, colorH, colorL)
			dx, dy, _ := corner(i+1, j+1, colorH, colorL)
			if isInfPoint(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			fmt.Fprintf(out, "<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, colorH, colorL string) (float64, float64, string) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	if z > 0 {
		return sx, sy, colorH
	}
	return sx, sy, colorL
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func isInfPoint(ax, ay, bx, by, cx, cy, dx, dy float64) bool {
	if math.IsInf(ax, 0) || math.IsInf(ay, 0) ||
		math.IsInf(bx, 0) || math.IsInf(by, 0) ||
		math.IsInf(cx, 0) || math.IsInf(cy, 0) ||
		math.IsInf(dx, 0) || math.IsInf(dy, 0) {
		return true
	}
	return false
}
