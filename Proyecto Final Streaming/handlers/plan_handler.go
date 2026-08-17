package handlers

import (
	"encoding/json"
	"net/http"
	"streaming-go/models"
	"streaming-go/services"
)

type PlanHandler struct {
	service *services.PlanService
}

func NuevoPlanHandler(service *services.PlanService) *PlanHandler {
	return &PlanHandler{service: service}
}

func (h *PlanHandler) Crear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var plan models.Plan

	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.Crear(plan); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(plan)
}

func (h *PlanHandler) Listar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.Listar())
}
