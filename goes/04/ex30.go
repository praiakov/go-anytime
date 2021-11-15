package main

import "fmt"

func main() {
	x := map[string][]string{
		`eduardo_tua`:  []string{`computadoras`, `montaña`, `playa`},
		`carlos_ramon`: []string{`leer`, `comprar`, `musica`},
		`juan_bimba`:   []string{`helado`, `pintar`, `bailar`},
	}

	x[`luis_perez`] = []string{`trabajar`, `playa`, `cerveza`}

	if v, ok := x[`carlos_ramon`]; ok {
		fmt.Println("Sé eliminará la llave con el valor:", v)
		delete(x, `carlos_ramon`)
	}

	for llave, valor := range x {
		fmt.Println("Registro: ", llave)
		for i, val := range valor {
			fmt.Println("\t", i, val)
		}
	}

}
