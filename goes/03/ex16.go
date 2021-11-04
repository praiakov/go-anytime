package main

import "fmt"

const (
	endFor = 2021
)

func main() {
	na := 1992
	for {
		if na > endFor {
			break
		}
		fmt.Println(na)
		na++
	}
}
