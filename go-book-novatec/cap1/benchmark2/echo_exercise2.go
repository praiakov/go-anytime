package main

import (
	"fmt"
	"os"
)

func Echo2() {
	for i, arg := range os.Args[0:] {
		fmt.Println(i, arg)
	}
}

func main() {
	fmt.Println("Hello World")
}
