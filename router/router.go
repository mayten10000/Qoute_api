package router

import (
	"github.com/gorilla/mux"
	"github.com/mayten10000/Qoute_api/handlers"
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
