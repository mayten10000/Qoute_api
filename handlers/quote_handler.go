package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"quote-service/models"
	"quote-service/storage"
)

type QuoteHandler struct {
	Store *storage.MemoryStorage
}

func (h *QuoteHandler) AddQuote(w http.ResponseWriter, r *http.Request) {
	var q models.Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	created := h.Store.AddQuote(q)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func (h *QuoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	var quotes []models.Quote
	if author != "" {
		quotes = h.Store.GetByAuthor(author)
	} else {
		quotes = h.Store.GetAll()
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *QuoteHandler) GetRandom(w http.ResponseWriter, r *http.Request) {
	quote, ok := h.Store.GetRandom()
	if !ok {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func (h *QuoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if !h.Store.DeleteByID(id) {
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
