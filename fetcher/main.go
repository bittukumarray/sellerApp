package main

import (
	"fetcher/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/crawl", handler.Fetch).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8000", r))
}
