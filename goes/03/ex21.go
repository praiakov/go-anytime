package main

import "fmt"

func main() {
	deporteFav := "futebol"
	switch deporteFav {
	case "futebol":
		fmt.Println("Futebol é bom")
	case "corrida":
		fmt.Println("Corrida não deveria imprimir")
	}
}
