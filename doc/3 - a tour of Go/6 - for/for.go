package main

import "fmt"

func main() {
	sum := 0
	cont := 0
	while := 1

	for i := 0; i < 10; i++ {
		sum += 1
	}

	fmt.Println(sum)

	//continued
	cont = continued()
	fmt.Println(cont)

	//for is gos while
	for while < 1000 {
		while += while
	}
	fmt.Println(while)
}
