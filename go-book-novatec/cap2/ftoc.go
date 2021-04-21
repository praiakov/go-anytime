// Ftoc exibe duas conversões de Fahrenheit para Celsius
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C \n", freezingF, ftoC(freezingF)) // 32F = 0 C
	fmt.Printf("%g F = %g C \n", boilingF, ftoC(boilingF))   // 212 F = 100 C
}

func ftoC(f float64) float64 {
	return (f - 32) * 5 / 9
}
