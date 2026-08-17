package handlers

import (
	"encoding/json"
	"net/http"

	"streaming-go/models"
	"streaming-go/services"
)

type ComentarioHandler struct {
	service *services.ComentarioService
}

func NuevoComentarioHandler(service *services.ComentarioService) *ComentarioHandler {
	return &ComentarioHandler{
		service: service,
	}
}

func (h *ComentarioHandler) Listar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	comentarios := h.service.Listar()

	json.NewEncoder(w).Encode(comentarios)
}

func (h *ComentarioHandler) Crear(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var comentario models.Comentario

	if err := json.NewDecoder(r.Body).Decode(&comentario); err != nil {
		http.Error(
			w,
			`{"error":"JSON inválido"}`,
			http.StatusBadRequest,
		)
		return
	}

	if err := h.service.Crear(comentario); err != nil {
		http.Error(
			w,
			`{"error":"`+err.Error()+`"}`,
			http.StatusBadRequest,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(comentario)
}
