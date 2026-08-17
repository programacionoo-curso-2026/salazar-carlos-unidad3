package models

type Comentario struct {
	ComentarioID string `json:"comentario_id"`
	PerfilID     string `json:"perfil_id"`
	ContenidoID  string `json:"contenido_id"`
	Texto        string `json:"texto"`
}
