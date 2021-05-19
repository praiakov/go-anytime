package main

import (
	"fmt"
)

func max(values ...int) int {
	var max int
	var min int

	for _, i := range values {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}

	return max
}

func main() {
	fmt.Println(max(12, 14, 44))
}
