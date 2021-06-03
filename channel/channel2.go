package main

import (
	"fmt"
)

func main() {
	mensagens := make(chan string)

	go func() { mensagens <- "ping" }()

	msg := <-mensagens

	fmt.Println(msg)
}
