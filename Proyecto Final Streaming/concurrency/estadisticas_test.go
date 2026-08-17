package concurrency

import (
	"testing"

	"streaming-go/models"
	"streaming-go/repository"
)

func TestObtenerEstadisticas(t *testing.T) {

	repo := repository.NuevaMemoria()

	repo.Contenidos["CONT001"] = models.Contenido{
		ContenidoID: "CONT001",
		Titulo:      "Contenido de prueba",
	}

	repo.Favoritos["FAV001"] = models.Favorito{}

	repo.Historiales["HIS001"] = models.Historial{}

	resultado := ObtenerEstadisticas(repo)

	if resultado.Contenidos != 1 {
		t.Fatalf(
			"se esperaban 1 contenido, se obtuvieron %d",
			resultado.Contenidos,
		)
	}

	if resultado.Favoritos != 1 {
		t.Fatalf(
			"se esperaba 1 favorito, se obtuvieron %d",
			resultado.Favoritos,
		)
	}

	if resultado.Visualizaciones != 1 {
		t.Fatalf(
			"se esperaba 1 visualización, se obtuvieron %d",
			resultado.Visualizaciones,
		)
	}
}
