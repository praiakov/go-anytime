package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)
}
