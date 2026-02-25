package main

import (
	"log"
	"net/http"

	"upcycleconnect/internal/routes"
)

func main() {
	mux := routes.NewRouter()

	addr := ":8080"
	log.Printf("API listening on %s", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
