package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temple = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Camisa Hering", Preco: 39.99, Quantidade: 5},
		{Nome: "Notebook", Descricao: "Dell", Preco: 1939.99, Quantidade: 2},
		{"Tenis", "Nike", 399.99, 3},
	}

	temple.ExecuteTemplate(w, "Index", produtos)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
