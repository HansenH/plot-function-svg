package main

import (
	"math"
	"os"
	"plot-function-svg/plot"
)

func main() {
	cfg := &plot.PlotConfig{
		Width:       600,
		Height:      400,
		Cells:       100,
		XYrange:     2,
		Xoffset:     0.4,
		Yoffset:     0.3,
		Zscale:      0.1,
		CameraAngle: math.Pi / 6,
		Color:       true,
		RightHand:   false,
	}
	f := func(x, y float64) float64 {
		return 2*math.Pow(x, 2) + 0.5*math.Sin(-7*y)
	}
	file, _ := os.Create("example2.svg")
	plot.PlotSVG(f, file, cfg)
}
