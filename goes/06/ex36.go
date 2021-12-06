package main

import "fmt"

func main() {
	ii := []int{1, 12, 3, 21, 12, 12, 3, 3}

	n := foo(ii...)
	fmt.Println(n)

	n2 := bar(ii)
	fmt.Println(n2)

}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}
