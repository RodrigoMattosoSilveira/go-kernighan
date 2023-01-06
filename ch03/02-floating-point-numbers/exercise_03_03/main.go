// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				break
			}
			bx, by, err := corner(i, j)
			if err != nil {
				break
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				break
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				break
			}
			//minZ := math.Max(az, bz, cz, dz)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='blue'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.

	z, err := f(x, y)
	if err != nil {
		return 0, 0, err
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) (float64, error) {
	if math.IsInf(math.Hypot(x, y), 0) || math.IsNaN(math.Hypot(x, y)) {
		return 0, fmt.Errorf("invalid poligon")
	}
	r := math.Hypot(x, y) // distance from (0,0)
	if math.IsInf(math.Sin(r)/r, 0) || math.IsNaN(math.Sin(r)/r) {
		return 0, fmt.Errorf("invalid poligon")
	}
	return math.Sin(r) / r, nil
}

func fill(z float64) string {
	color := "white"
	if z < 0.1 {
		color = "blue"
	} else {
		if z > 0.9 {
			color = "red"
		}
	}
	return color
}
