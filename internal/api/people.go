package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	lastName := r.URL.Query().Get("last_name")

	if lastName == "" {
		http.Error(w, "missing last_name query param", http.StatusBadRequest)
		return
	}

	// TODO: передать в сервис и получить []Person

	log.Printf("Search people by last name: %s", lastName)

	// TODO: временная заглушка
	people := []any{
		map[string]any{
			"first_name":  "Иван",
			"last_name":   "Иванов",
			"middle_name": "Петрович",
			"age":         34,
			"gender":      "male",
			"nationality": "RU",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
