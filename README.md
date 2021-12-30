# plot-function-svg
Plot 3D math equation z=f(x, y) with SVG format.

Some codes are referred from https://github.com/adonovan/gopl.io  
licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>.<br/>
<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png"/></a>

## APIs
```go
type PlotConfig struct {
	Width       int
	Height      int
	Cells       int
	XYrange     float64
    Xoffset     float64
	Yoffset     float64
	Zscale      float64
	CameraAngle float64
	Color       bool
    RightHand   bool
}

func DefaultPlotConfig() *PlotConfig {
	return &PlotConfig{
		Width:       600,
		Height:      320,
		Cells:       100,
		XYrange:     30.0,
		Xoffset:     0,
		Yoffset:     0,
		Zscale:      0.4,
		CameraAngle: math.Pi / 6,
		Color:       true,
        RightHand:   true,
	}
}

// Draw an SVG to io.Writer. Default w (use nil) is os.Stdout. Default cfg (use nil) is DefaultPlotConfig().
func PlotSVG(f func(x, y float64) float64, w io.Writer, cfg *PlotConfig) 

```
  
Right-hand coordinates (default):  
&nbsp;&nbsp;&nbsp;&nbsp;z  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ y  
x /  
<br>
Left-hand coordinates:  
&nbsp;&nbsp;&nbsp;&nbsp;z  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|__ x  
y /  

## Examples
Example1:  
```go
f := func(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
file, _ := os.Create("example1.svg")
plot.PlotSVG(f, file, nil)
```
![image](https://github.com/HansenH/plot-function-svg/blob/main/examples/example1/example1.png)  

Example2:  
```go
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
```
![image](https://github.com/HansenH/plot-function-svg/blob/main/examples/example2/example2.png)  
