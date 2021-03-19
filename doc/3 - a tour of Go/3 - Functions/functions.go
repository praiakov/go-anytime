package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println((add(42, 13)))

	// multiple results
	a, b := swap("hello", "world")

	fmt.Println(a, b)

	//named results
	fmt.Println(split(17))

}
