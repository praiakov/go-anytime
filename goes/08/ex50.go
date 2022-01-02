package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 4, 41, 12, 2}
	xs := []string{"adr", "pr", "ea", "gt"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
