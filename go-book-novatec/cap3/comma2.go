package main

import (
	"bytes"
	"fmt"
)

// comma insere vírgulas em uma string que representa um inteiro decimal
// não negativo
func comma(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	var buf bytes.Buffer
	for _, s := range str {
		fmt.Fprintf(&buf, "%c", s)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("et21332"))
}
