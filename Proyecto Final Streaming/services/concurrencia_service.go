package services

import (
	"sync"

	"streaming-go/concurrency"
	"streaming-go/models"
)

func ProcesarHistorialConcurrente(historiales []models.Historial) {
	canal := make(chan models.Historial)

	var wg sync.WaitGroup

	go concurrency.ProcesarHistorial(canal)

	for _, historial := range historiales {
		wg.Add(1)

		go func(h models.Historial) {
			defer wg.Done()

			concurrency.RegistrarHistorial(canal, h)
		}(historial)
	}

	wg.Wait()
	close(canal)
}
