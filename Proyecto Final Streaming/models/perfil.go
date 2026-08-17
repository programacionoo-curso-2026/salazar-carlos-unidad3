package models

type Perfil struct {
	PerfilID   string `json:"perfil_id"`
	CuentaID   string `json:"cuenta_id"`
	Nombre     string `json:"nombre"`
	Avatar     string `json:"avatar"`
	TipoPerfil string `json:"tipo_perfil"`
}
