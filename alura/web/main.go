package main

import (
	"net/http"

	"github.com/praiakov/web/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
