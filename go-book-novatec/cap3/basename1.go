package main

import "fmt"

// basename remove componentes de diretorio e um sufixo
// exemplos: a=> a, a, a/b/c.go => c, a/b.c.go => b.c

func basenames(s string) string {
	// Descarta a última '/' e tudo que estiver antes
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserva tudo que estiver antes do último '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	fmt.Println(basenames("a/b/c.go"))
}
