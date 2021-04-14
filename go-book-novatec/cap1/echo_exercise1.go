package main

import (
	"fmt"
	"os"
	"strings"
)

func Joiner() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func main() {
	Joiner()
}
