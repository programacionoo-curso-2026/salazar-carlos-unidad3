package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"streaming-go/models"
	"streaming-go/services"
)

func (h *ContenidoHandler) Obtener(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/api/contenidos/")

	contenido, err := h.service.Obtener(id)

	if err != nil {
		http.Error(
			w,
			`{"error":"`+err.Error()+`"}`,
			http.StatusNotFound,
		)
		return
	}

	json.NewEncoder(w).Encode(contenido)
}

type ContenidoHandler struct {
	service *services.ContenidoService
}

func NuevoContenidoHandler(service *services.ContenidoService) *ContenidoHandler {
	return &ContenidoHandler{service: service}
}

func (h *ContenidoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contenido models.Contenido

	err := json.NewDecoder(r.Body).Decode(&contenido)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	err = h.service.Crear(contenido)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contenido)
}

func (h *ContenidoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contenidos := h.service.Listar()

	json.NewEncoder(w).Encode(contenidos)
}
func (h *ContenidoHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/api/contenidos/")

	err := h.service.Eliminar(id)

	if err != nil {
		http.Error(
			w,
			`{"error":"`+err.Error()+`"}`,
			http.StatusNotFound,
		)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Contenido eliminado correctamente",
	})
}
func (h *ContenidoHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/api/contenidos/")

	var contenido models.Contenido

	if err := json.NewDecoder(r.Body).Decode(&contenido); err != nil {
		http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.Actualizar(id, contenido); err != nil {
		http.Error(
			w,
			`{"error":"`+err.Error()+`"}`,
			http.StatusBadRequest,
		)
		return
	}

	resultado, _ := h.service.Obtener(id)

	json.NewEncoder(w).Encode(resultado)
}
