package main

import "fmt"

const (
	endFor = 90
)

func main() {
	for i := 65; i < endFor; i++ {
		fmt.Printf("%d", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
