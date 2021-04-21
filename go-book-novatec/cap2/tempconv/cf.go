// Cf converte seu argumento numérico para Celsius e Fahrenheit
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/praiakov/tconv/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		m := tempconv.Metros(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s", m, tempconv.EToM(m))
	}
}
