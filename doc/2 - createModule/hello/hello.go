package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Praia")
	fmt.Println(message)
}
