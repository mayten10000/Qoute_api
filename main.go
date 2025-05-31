package main

import (
	"log"
	"net/http"

	"github.com/mayten10000/Qoute_api/handlers"
	"github.com/mayten10000/Qoute_api/router"
	"github.com/mayten10000/Qoute_api/storage"
)

func main() {
	store := storage.NewMemoryStorage()
	handler := &handlers.QuoteHandler{Store: store}
	r := router.NewRouter(handler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
