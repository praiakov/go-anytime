package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura password=macacoaurelio sslmode=disable"

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temple = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	selectProdutos, err := db.Query("Select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	temple.ExecuteTemplate(w, "Index", produtos)

	defer db.Close()
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
