package repository

import "streaming-go/models"

type Memoria struct {
	Cuentas      map[string]models.Cuenta
	Perfiles     map[string]models.Perfil
	Contenidos   map[string]models.Contenido
	Favoritos    map[string]models.Favorito
	Historiales  map[string]models.Historial
	Comentarios  map[string]models.Comentario
	Valoraciones map[string]models.Valoracion
	Planes       map[string]models.Plan
	Proveedores  map[string]models.ProveedorVideo
	Pagos        map[string]models.Pago
	Generos      map[string]models.Genero
	Categorias   map[string]models.Categoria
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		Cuentas:      make(map[string]models.Cuenta),
		Perfiles:     make(map[string]models.Perfil),
		Contenidos:   make(map[string]models.Contenido),
		Favoritos:    make(map[string]models.Favorito),
		Historiales:  make(map[string]models.Historial),
		Comentarios:  make(map[string]models.Comentario),
		Valoraciones: make(map[string]models.Valoracion),
		Planes:       make(map[string]models.Plan),
		Proveedores:  make(map[string]models.ProveedorVideo),
		Pagos:        make(map[string]models.Pago),
		Generos:      make(map[string]models.Genero),
		Categorias:   make(map[string]models.Categoria),
	}
}
