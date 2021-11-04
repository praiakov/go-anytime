package main

import "fmt"

const (
	endFor = 100
)

func main() {
	for i := 10; i < endFor; i++ {
		fmt.Printf("Se dividirmos %v entre 4, o resto é %v\n", i, i%4)
	}
}
