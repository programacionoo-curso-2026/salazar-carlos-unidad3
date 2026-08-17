package services

import (
	"streaming-go/interfaces"
)

func IniciarReproduccion(contenido interfaces.Reproducible) string {
	return contenido.Reproducir()
}
