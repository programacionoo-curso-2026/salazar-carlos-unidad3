package models

type Contenido struct {
	ContenidoID     string `json:"contenido_id"`
	Titulo          string `json:"titulo"`
	Sinopsis        string `json:"sinopsis"`
	Duracion        int    `json:"duracion"`
	Anio            int    `json:"anio"`
	Clasificacion   string `json:"clasificacion"`
	GeneroID        string `json:"genero_id"`
	CategoriaID     string `json:"categoria_id"`
	ProveedorID     string `json:"proveedor_id"`
	VideoID         string `json:"video_id"`
	URLReproduccion string `json:"url_reproduccion"`
}
