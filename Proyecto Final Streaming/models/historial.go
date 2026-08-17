package models

import "time"

type Historial struct {
	HistorialID        string    `json:"historial_id"`
	PerfilID           string    `json:"perfil_id"`
	ContenidoID        string    `json:"contenido_id"`
	Progreso           float64   `json:"progreso"`
	TiempoActual       int       `json:"tiempo_actual"`
	FechaVisualizacion time.Time `json:"fecha_visualizacion"`
}
