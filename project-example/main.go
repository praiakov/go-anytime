package main

import (
	"fmt"

	"github.com/praiakov/project/models"
)

func main() {
	produto := models.Example{}
	produto.Name = "Hello World"

	fmt.Println("Welcome ", produto.Name)

}
