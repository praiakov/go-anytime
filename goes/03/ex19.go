package main

import "fmt"

func main() {
	x := "Batman"
	if x == "Robin" {
		fmt.Println(x)
	} else if x == "Batman" {
		fmt.Println("batbatbat", x)
	} else {
		fmt.Println("Nenhum super heroi")
	}
}
