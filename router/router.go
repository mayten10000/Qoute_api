package router

import (
	"Qoute_api/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(h *handlers.QuoteHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/quotes", h.AddQuote).Methods("POST")
	r.HandleFunc("/quotes", h.GetAll).Methods("GET")
	r.HandleFunc("/quotes/random", h.GetRandom).Methods("GET")
	r.HandleFunc("/quotes/{id}", h.Delete).Methods("DELETE")

	return r
}
