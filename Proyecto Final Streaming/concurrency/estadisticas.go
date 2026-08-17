package concurrency

import "streaming-go/repository"

type Estadisticas struct {
	Contenidos      int `json:"contenidos"`
	Favoritos       int `json:"favoritos"`
	Visualizaciones int `json:"visualizaciones"`
}

func ObtenerEstadisticas(repo *repository.Memoria) Estadisticas {

	contenidosChan := make(chan int)
	favoritosChan := make(chan int)
	historialChan := make(chan int)

	go func() {
		contenidosChan <- len(repo.Contenidos)
	}()

	go func() {
		favoritosChan <- len(repo.Favoritos)
	}()

	go func() {
		historialChan <- len(repo.Historiales)
	}()

	return Estadisticas{
		Contenidos:      <-contenidosChan,
		Favoritos:       <-favoritosChan,
		Visualizaciones: <-historialChan,
	}
}
