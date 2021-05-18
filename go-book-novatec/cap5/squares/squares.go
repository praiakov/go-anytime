// squares devolve uma função que devolve
// o próximo quadrado perfeito sempre que ela é chamada

package main

import (
	"fmt"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
