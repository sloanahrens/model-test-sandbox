package main

import (
	"fmt"
	"log"
	"net/http"

	"sandbox/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /api/items", handler.ListItems)
	mux.HandleFunc("POST /api/items", handler.CreateItem)

	addr := ":8080"
	fmt.Printf("Server starting on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
