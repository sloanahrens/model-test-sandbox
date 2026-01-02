package handler

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	items   = make([]Item, 0)
	itemsMu sync.RWMutex
)

func ListItems(w http.ResponseWriter, r *http.Request) {
	itemsMu.RLock()
	defer itemsMu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	item := Item{
		ID:        generateID(),
		Name:      input.Name,
		CreatedAt: time.Now(),
	}

	itemsMu.Lock()
	items = append(items, item)
	itemsMu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func generateID() string {
	return time.Now().Format("20060102150405")
}
