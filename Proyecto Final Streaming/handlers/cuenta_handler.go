package handlers

import (
	"encoding/json"
	"net/http"
	"streaming-go/models"
	"streaming-go/services"
	"strings"
)

type CuentaHandler struct {
	service *services.CuentaService
}

func NuevaCuentaHandler(service *services.CuentaService) *CuentaHandler {
	return &CuentaHandler{service: service}
}

func (h *CuentaHandler) Crear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cuenta models.Cuenta

	if err := json.NewDecoder(r.Body).Decode(&cuenta); err != nil {
		http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.Crear(cuenta); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cuenta)
}

func (h *CuentaHandler) Listar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.Listar())
}

func (h *CuentaHandler) Obtener(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/api/cuentas/")

	cuenta, err := h.service.Obtener(id)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cuenta)
}
