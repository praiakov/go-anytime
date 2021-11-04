package main

import "fmt"

const (
	endFor = 2021
)

func main() {
	na := 1992
	for na <= endFor {
		fmt.Println(na)
		na++
	}
}
