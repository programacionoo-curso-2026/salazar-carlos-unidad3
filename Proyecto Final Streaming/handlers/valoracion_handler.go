package handlers

import (
	"encoding/json"
	"net/http"
	"streaming-go/models"
	"streaming-go/services"
)

type ValoracionHandler struct {
	service *services.ValoracionService
}

func NuevaValoracionHandler(service *services.ValoracionService) *ValoracionHandler {
	return &ValoracionHandler{service: service}
}

func (h *ValoracionHandler) Crear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var valoracion models.Valoracion

	if err := json.NewDecoder(r.Body).Decode(&valoracion); err != nil {
		http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.Crear(valoracion); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(valoracion)
}

func (h *ValoracionHandler) Listar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.Listar())
}
