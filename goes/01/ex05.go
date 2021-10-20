package main

import (
	"fmt"
)

type numero int

var x numero
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("Tipo de x é: %T\n", x)

	x = 42
	fmt.Println(x)
	y = int(x)

	fmt.Println(y)
	fmt.Printf("Tipo de x é: %T\n", y)
}
