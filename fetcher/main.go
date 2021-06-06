package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fetcher/handler"
)



func main()  {
	r:=mux.NewRouter()
	r.HandleFunc("/crawl", handler.Fetch).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8000", r))
}