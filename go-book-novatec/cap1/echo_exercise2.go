package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Println(i, arg)
	}
}
