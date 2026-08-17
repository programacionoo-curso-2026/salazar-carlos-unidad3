package handlers

import (
	"encoding/json"
	"net/http"

	"streaming-go/concurrency"
	"streaming-go/repository"
)

type EstadisticasHandler struct {
	repo *repository.Memoria
}

func NuevoEstadisticasHandler(repo *repository.Memoria) *EstadisticasHandler {
	return &EstadisticasHandler{
		repo: repo,
	}
}

func (h *EstadisticasHandler) Obtener(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	estadisticas := concurrency.ObtenerEstadisticas(h.repo)

	json.NewEncoder(w).Encode(estadisticas)
}
