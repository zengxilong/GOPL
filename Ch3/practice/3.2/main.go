package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isInfPoint(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f2(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
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

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func f2(x, y float64) float64 {
	return math.Pow(math.Sin(x/xyrange*3*math.Pi), 2) * math.Cos(y/xyrange*3*math.Pi)
}
func f3(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 10
}
