// Some codes are referred from https://github.com/adonovan/gopl.io
// Copyright © 2021 HansenH https://github.com/HansenH
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type PlotConfig struct {
	Width       int
	Height      int
	Cells       int
	XYrange     float64
	Zscale      float64
	CameraAngle float64
	Color       bool
}

func DefaultPlotConfig() *PlotConfig {
	return &PlotConfig{
		Width:       600,
		Height:      320,
		Cells:       100,
		XYrange:     30.0,
		Zscale:      0.4,
		CameraAngle: math.Pi / 6,
		Color:       true,
	}
}

// Draw an SVG to io.Writer. Default w (use nil) is os.Stdout. Default cfg (use nil) is DefaultPlotConfig().
func PlotSVG(f func(x, y float64) float64, w io.Writer, cfg *PlotConfig) {
	if w == nil {
		w = os.Stdout
	}
	if cfg == nil {
		cfg = DefaultPlotConfig()
	}
	zmin, zmax := minmax(f, cfg)
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.5' "+
		"width='%d' height='%d'>", cfg.Width, cfg.Height)
	for i := 0; i < cfg.Cells; i++ {
		for j := 0; j < cfg.Cells; j++ {
			ax, ay := cfg.projectToCanvas(cfg.getXYZ(f, i, j))
			bx, by := cfg.projectToCanvas(cfg.getXYZ(f, i+1, j))
			cx, cy := cfg.projectToCanvas(cfg.getXYZ(f, i+1, j+1))
			dx, dy := cfg.projectToCanvas(cfg.getXYZ(f, i, j+1))
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			color := "grey"
			if cfg.Color && zmax != zmin {
				_, _, z := cfg.getXYZ(f, i, j)
				color = getColor((z - zmin) / (zmax - zmin))
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

// Find point (x,y) and z=f(x, y) at corner of cell (i,j).
func (cfg *PlotConfig) getXYZ(f func(x, y float64) float64, i, j int) (x, y, z float64) {
	x = cfg.XYrange * (float64(i)/float64(cfg.Cells) - 0.5)
	y = cfg.XYrange * (float64(j)/float64(cfg.Cells) - 0.5)
	z = f(x, y)
	return
}

// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
func (cfg *PlotConfig) projectToCanvas(x, y, z float64) (float64, float64) {
	sinAngle, cosAngle := math.Sin(cfg.CameraAngle), math.Cos(cfg.CameraAngle)
	xyscale := float64(cfg.Width) / 2 / cfg.XYrange
	zCoeff := cfg.Zscale * float64(cfg.Height)
	sx := float64(cfg.Width)/2 + (x-y)*cosAngle*xyscale
	sy := float64(cfg.Height)/2 + (x+y)*sinAngle*xyscale - z*zCoeff
	return sx, sy
}

// 0 <= epsilon <=1
func getColor(epsilon float64) string {
	red := int(0xff * epsilon)
	blue := int(0xff * (1 - epsilon))
	return fmt.Sprintf("#%02x00%02x", red, blue)
}

func minmax(f func(x, y float64) float64, cfg *PlotConfig) (min, max float64) {
	min = math.Inf(1)
	max = math.Inf(-1)
	for i := 0; i < cfg.Cells; i++ {
		for j := 0; j < cfg.Cells; j++ {
			_, _, z := cfg.getXYZ(f, i, j)
			if z < min {
				min = z
			}
			if z > max {
				max = z
			}
		}
	}
	return
}

func main() {
	f := func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return math.Sin(r) / r
	}
	PlotSVG(f, nil, nil)
}
