package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	store *Store
}

type CreateIncidentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) listIncidents(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.List())
}

func (h *Handler) createIncident(w http.ResponseWriter, r *http.Request) {
	var req CreateIncidentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
		return
	}

	if req.Title == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "o campo título é obrigatório"})
		return
	}

	incident := h.store.Create(req.Title, req.Description)
	writeJSON(w, http.StatusCreated, incident)
}
