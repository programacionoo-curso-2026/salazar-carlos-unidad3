package handlers

import (
	"encoding/json"
	"net/http"
	"streaming-go/models"
	"streaming-go/services"
)

type HistorialHandler struct {
	service *services.HistorialService
}

func NuevoHistorialHandler(service *services.HistorialService) *HistorialHandler {
	return &HistorialHandler{service: service}
}

func (h *HistorialHandler) Crear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var historial models.Historial

	if err := json.NewDecoder(r.Body).Decode(&historial); err != nil {
		http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.Crear(historial); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(historial)
}

func (h *HistorialHandler) Listar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.Listar())
}
