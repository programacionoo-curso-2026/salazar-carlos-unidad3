package models

type Favorito struct {
	FavoritoID  string `json:"favorito_id"`
	PerfilID    string `json:"perfil_id"`
	ContenidoID string `json:"contenido_id"`
}
