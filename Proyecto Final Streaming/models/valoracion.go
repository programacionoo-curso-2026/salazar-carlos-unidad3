package models

type Valoracion struct {
	ValoracionID string `json:"valoracion_id"`
	PerfilID     string `json:"perfil_id"`
	ContenidoID  string `json:"contenido_id"`
	Puntuacion   int    `json:"puntuacion"`
}
