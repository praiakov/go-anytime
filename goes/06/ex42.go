package main

import "fmt"

var x int
var g func()

func main() {
	f := foo()

	fmt.Println(f())

}

func foo() func() int {
	return func() int {
		return 42
	}
}
