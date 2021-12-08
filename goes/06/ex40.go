package main

import "fmt"

func main() {
	func() {
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("Listo!")

}
