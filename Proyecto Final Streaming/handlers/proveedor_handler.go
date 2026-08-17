package handlers

import (
	"encoding/json"
	"net/http"
	"streaming-go/models"
	"streaming-go/services"
)

type ProveedorHandler struct {
	service *services.ProveedorService
}

func NuevoProveedorHandler(service *services.ProveedorService) *ProveedorHandler {
	return &ProveedorHandler{service: service}
}

func (h *ProveedorHandler) Crear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var proveedor models.ProveedorVideo

	if err := json.NewDecoder(r.Body).Decode(&proveedor); err != nil {
		http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.Crear(proveedor); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(proveedor)
}

func (h *ProveedorHandler) Listar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.Listar())
}
