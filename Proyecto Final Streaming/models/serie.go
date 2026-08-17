package models

type Serie struct {
	Contenido
	Temporadas int `json:"temporadas"`
}

func (s Serie) Reproducir() string {
	return "Reproduciendo serie: " + s.Titulo
}
