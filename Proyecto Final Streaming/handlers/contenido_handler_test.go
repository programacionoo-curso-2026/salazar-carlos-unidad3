package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"streaming-go/repository"
	"streaming-go/services"
)

func TestContenidoListar(t *testing.T) {

	repo := repository.NuevaMemoria()
	service := services.NuevoContenidoService(repo)
	handler := NuevoContenidoHandler(service)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/contenidos",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.Listar(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("se esperaba status 200, se obtuvo %d", rec.Code)
	}
}

func TestContenidoCrear(t *testing.T) {

	repo := repository.NuevaMemoria()
	service := services.NuevoContenidoService(repo)
	handler := NuevoContenidoHandler(service)

	jsonContenido := `{
		"contenido_id": "CONT_TEST",
		"titulo": "Contenido de prueba",
		"sinopsis": "Prueba del servicio web",
		"duracion": 120,
		"anio": 2026,
		"clasificacion": "PG-13",
		"genero_id": "GEN001",
		"categoria_id": "CAT001",
		"proveedor_id": "PROV001",
		"video_id": "VID_TEST",
		"url_reproduccion": "https://example.com/video"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/contenidos",
		strings.NewReader(jsonContenido),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.Crear(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("se esperaba status 201, se obtuvo %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "CONT_TEST") {
		t.Fatal("el contenido creado no aparece en la respuesta")
	}
}

func TestContenidoCrearJSONInvalido(t *testing.T) {

	repo := repository.NuevaMemoria()
	service := services.NuevoContenidoService(repo)
	handler := NuevoContenidoHandler(service)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/contenidos",
		strings.NewReader(`{"contenido_id":`),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.Crear(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("se esperaba status 400, se obtuvo %d", rec.Code)
	}
}

func TestContenidoObtener(t *testing.T) {

	repo := repository.NuevaMemoria()
	service := services.NuevoContenidoService(repo)
	handler := NuevoContenidoHandler(service)

	jsonContenido := `{
		"contenido_id": "CONT_GET",
		"titulo": "Interestelar",
		"sinopsis": "Una prueba de contenido",
		"duracion": 169,
		"anio": 2014,
		"clasificacion": "PG-13",
		"genero_id": "GEN001",
		"categoria_id": "CAT001",
		"proveedor_id": "PROV001",
		"video_id": "VID001",
		"url_reproduccion": "https://example.com/video"
	}`

	reqCrear := httptest.NewRequest(
		http.MethodPost,
		"/api/contenidos",
		strings.NewReader(jsonContenido),
	)

	recCrear := httptest.NewRecorder()

	handler.Crear(recCrear, reqCrear)

	reqObtener := httptest.NewRequest(
		http.MethodGet,
		"/api/contenidos/CONT_GET",
		nil,
	)

	recObtener := httptest.NewRecorder()

	handler.Obtener(recObtener, reqObtener)

	if recObtener.Code != http.StatusOK {
		t.Fatalf("se esperaba status 200, se obtuvo %d", recObtener.Code)
	}

	if !strings.Contains(recObtener.Body.String(), "CONT_GET") {
		t.Fatal("no se encontró el contenido solicitado")
	}
}

func TestContenidoActualizar(t *testing.T) {

	repo := repository.NuevaMemoria()
	service := services.NuevoContenidoService(repo)
	handler := NuevoContenidoHandler(service)

	jsonCrear := `{
		"contenido_id": "CONT_UPDATE",
		"titulo": "Película Original",
		"sinopsis": "Contenido original",
		"duracion": 100,
		"anio": 2020,
		"clasificacion": "PG",
		"genero_id": "GEN001",
		"categoria_id": "CAT001",
		"proveedor_id": "PROV001",
		"video_id": "VID_UPDATE",
		"url_reproduccion": "https://example.com/video"
	}`

	reqCrear := httptest.NewRequest(
		http.MethodPost,
		"/api/contenidos",
		strings.NewReader(jsonCrear),
	)

	recCrear := httptest.NewRecorder()

	handler.Crear(recCrear, reqCrear)

	jsonActualizar := `{
		"titulo": "Película Actualizada",
		"sinopsis": "Contenido actualizado",
		"duracion": 120,
		"anio": 2026,
		"clasificacion": "PG-13",
		"genero_id": "GEN001",
		"categoria_id": "CAT001",
		"proveedor_id": "PROV001",
		"video_id": "VID_UPDATE",
		"url_reproduccion": "https://example.com/video"
	}`

	reqActualizar := httptest.NewRequest(
		http.MethodPut,
		"/api/contenidos/CONT_UPDATE",
		strings.NewReader(jsonActualizar),
	)

	recActualizar := httptest.NewRecorder()

	handler.Actualizar(recActualizar, reqActualizar)

	if recActualizar.Code != http.StatusOK {
		t.Fatalf("se esperaba status 200, se obtuvo %d", recActualizar.Code)
	}

	if !strings.Contains(
		recActualizar.Body.String(),
		"Película Actualizada",
	) {
		t.Fatal("el contenido no fue actualizado correctamente")
	}
}

func TestContenidoEliminar(t *testing.T) {

	repo := repository.NuevaMemoria()
	service := services.NuevoContenidoService(repo)
	handler := NuevoContenidoHandler(service)

	jsonContenido := `{
		"contenido_id": "CONT_DELETE",
		"titulo": "Contenido para eliminar",
		"sinopsis": "Prueba de eliminación",
		"duracion": 90,
		"anio": 2026,
		"clasificacion": "PG",
		"genero_id": "GEN001",
		"categoria_id": "CAT001",
		"proveedor_id": "PROV001",
		"video_id": "VID_DELETE",
		"url_reproduccion": "https://example.com/video"
	}`

	reqCrear := httptest.NewRequest(
		http.MethodPost,
		"/api/contenidos",
		strings.NewReader(jsonContenido),
	)

	recCrear := httptest.NewRecorder()

	handler.Crear(recCrear, reqCrear)

	reqEliminar := httptest.NewRequest(
		http.MethodDelete,
		"/api/contenidos/CONT_DELETE",
		nil,
	)

	recEliminar := httptest.NewRecorder()

	handler.Eliminar(recEliminar, reqEliminar)

	if recEliminar.Code != http.StatusOK {
		t.Fatalf("se esperaba status 200, se obtuvo %d", recEliminar.Code)
	}

	if !strings.Contains(
		recEliminar.Body.String(),
		"eliminado correctamente",
	) {
		t.Fatal("el contenido no fue eliminado correctamente")
	}
}
