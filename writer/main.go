package main

import (
	"log"
	"net/http"
	"writer/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/product", handler.Create).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":9000", r))

}
