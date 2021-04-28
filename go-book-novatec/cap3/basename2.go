package main

import (
	"fmt"
	"strings"
)

// basename remove componentes de diretorio e um sufixo
// exemplos: a=> a, a, a/b/c.go => c, a/b.c.go => b.c

func basenames(s string) string {
	slash := strings.LastIndex(s, "/") // - 1 se "/" não fo encontrada
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

func main() {
	fmt.Println(basenames("a/b/c.go"))
}
