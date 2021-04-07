package main

import "fmt"

func main() {
	creature := "shark"
	pointer := &creature

	fmt.Println("creature=", creature)
	fmt.Println("pointer=", pointer)
	fmt.Println("*pointer=", *pointer)

	*pointer = "jellyfish"
	fmt.Println("*pointer=", *pointer)

	fmt.Println("creature =", creature)
}
