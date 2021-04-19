// server1 é um servidor de "eco" minimo

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) //cada requisição chama handler
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//handler ecoa o componente path do url requisitado
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
