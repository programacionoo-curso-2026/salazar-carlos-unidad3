package models

type Pelicula struct {
	Contenido
}

func (p Pelicula) Reproducir() string {
	return "Reproduciendo película: " + p.Titulo
}
