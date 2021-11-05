package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Não deveria imprimir")
	case true:
		fmt.Println("Deveria imprimir")
	}
}
