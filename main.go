package main

import (
	"Qoute_api/handlers"
	"Qoute_api/router"
	"Qoute_api/storage"
	"log"
	"net/http"
)

func main() {
	store := storage.NewMemoryStorage()
	handler := &handlers.QuoteHandler{Store: store}
	r := router.NewRouter(handler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
