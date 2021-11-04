package main

import "fmt"

const (
	endFor = 10000
)

func main() {
	for i := 0; i < endFor; i++ {
		fmt.Printf("number : %d", i)
	}
}
