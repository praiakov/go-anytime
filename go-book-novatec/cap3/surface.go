// Surface calcula uma renderização SVG de uma função de superficie 3D
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // tamanho do canvas em pixels
	cells         = 100                 // número de células da grade
	xyrange       = 30.0                // intervalos dos eixos (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels por unidade x ou y
	zscale        = height * 0.4        // pixels por unidade z
	angle         = math.Pi / 6         // angulo dos eixos x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

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
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Encontra o ponto (x,y) no canto da célula (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Calcula a altura Z da superficie
	z := f(x, y)

	// Faz uma projeção isometrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distância de (0,0)
	return math.Sin(r) / r
}
