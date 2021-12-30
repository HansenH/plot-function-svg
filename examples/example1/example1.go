package main

import (
	"math"
	"os"
	"plot-function-svg/plot"
)

func main() {
	f := func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return math.Sin(r) / r
	}
	file, _ := os.Create("example1.svg")
	plot.PlotSVG(f, file, nil)
}
