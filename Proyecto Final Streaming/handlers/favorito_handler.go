package handlers

import (
	"encoding/json"
	"net/http"
	"streaming-go/models"
	"streaming-go/services"
)

type FavoritoHandler struct {
	service *services.FavoritoService
}

func NuevoFavoritoHandler(service *services.FavoritoService) *FavoritoHandler {
	return &FavoritoHandler{service: service}
}

func (h *FavoritoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var favorito models.Favorito

	if err := json.NewDecoder(r.Body).Decode(&favorito); err != nil {
		http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.Crear(favorito); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(favorito)
}

func (h *FavoritoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.Listar())
}
