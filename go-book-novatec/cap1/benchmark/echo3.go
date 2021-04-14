package main

import (
	"fmt"
	"os"
	"strings"
)

func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	fmt.Println("Hello baby")
}
