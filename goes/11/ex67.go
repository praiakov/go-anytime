package main

import (
	"fmt"
)

type errorPer struct {
	info string
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("aquí está el error: %v", ep.info)
}

func main() {
	e1 := errorPer{
		info: "necesito dormir más",
	}

	foo(e1)
}

func foo(e error) {
	fmt.Println("foo corrió -", e)
	fmt.Println()
}
