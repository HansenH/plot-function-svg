# plot-function-svg
Plot 3D math equation z=f(x, y) with SVG format.

Some codes are referred from https://github.com/adonovan/gopl.io  
licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>.<br/>
<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png"/></a>

![image](https://github.com/HansenH/plot-function-svg/blob/main/examples/example1.png)  

## APIs
```go
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
func PlotSVG(f func(x, y float64) float64, w io.Writer, cfg *PlotConfig) 

```
	