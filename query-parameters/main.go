package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8880", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")

	// if !ok || len(keys[0]) < 1 {
	// 	log.Println("Url Param 'key' is missing")
	// 	return
	// }

	fmt.Fprintf(w, "Url Param 'key' is: "+string(key))
}
