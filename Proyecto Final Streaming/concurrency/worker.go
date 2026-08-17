package concurrency

import (
	"fmt"
	"streaming-go/models"
)

func ProcesarHistorial(canal <-chan models.Historial) {
	for historial := range canal {
		fmt.Println(
			"Procesando historial:",
			historial.HistorialID,
		)
	}
}

func RegistrarHistorial(canal chan<- models.Historial, historial models.Historial) {
	canal <- historial
}
