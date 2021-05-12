/*
From https://medium.com/@rafaelacioly/construindo-uma-api-restful-com-go-d6007e4faff6
*/
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praiakov/webapi/controllers"
	"github.com/praiakov/webapi/routes"
)

func main() {
	controllers.LoadData()
	router := mux.NewRouter()
	rot := routes.LoadRoutes(router)
	log.Fatal(http.ListenAndServe("localhost:8000", &rot))
}
