package repository

import (
	"streaming-go/models"
	"time"
)

func CargarDatosIniciales(repo *Memoria) {

	// PLANES
	repo.Planes["PLAN01"] = models.Plan{
		PlanID: "PLAN01",
		Nombre: "Básico",
		Precio: 5.99,
	}

	repo.Planes["PLAN02"] = models.Plan{
		PlanID: "PLAN02",
		Nombre: "Premium",
		Precio: 9.99,
	}

	// CUENTAS
	repo.Cuentas["CTA001"] = models.Cuenta{
		CuentaID:   "CTA001",
		Email:      "usuario@gmail.com",
		Contrasena: "123456",
		Estado:     "activo",
		PlanID:     "PLAN02",
	}

	// PERFILES
	repo.Perfiles["PER001"] = models.Perfil{
		PerfilID:   "PER001",
		CuentaID:   "CTA001",
		Nombre:     "Mateo",
		Avatar:     "avatar1.png",
		TipoPerfil: "adulto",
	}

	// GÉNEROS
	repo.Generos["GEN001"] = models.Genero{
		GeneroID: "GEN001",
		Nombre:   "Ciencia ficción",
	}

	repo.Generos["GEN002"] = models.Genero{
		GeneroID: "GEN002",
		Nombre:   "Acción",
	}

	// CATEGORÍAS
	repo.Categorias["CAT001"] = models.Categoria{
		CategoriaID: "CAT001",
		Nombre:      "Películas",
	}

	repo.Categorias["CAT002"] = models.Categoria{
		CategoriaID: "CAT002",
		Nombre:      "Series",
	}

	// PROVEEDORES
	repo.Proveedores["PROV001"] = models.ProveedorVideo{
		ProveedorID:     "PROV001",
		NombreProveedor: "StreamingGo",
		URLBase:         "https://streaminggo.example.com",
		Estado:          "activo",
	}

	// CONTENIDOS
	repo.Contenidos["CONT001"] = models.Contenido{
		ContenidoID:     "CONT001",
		Titulo:          "Interestelar",
		Sinopsis:        "Una misión espacial busca un nuevo hogar para la humanidad.",
		Duracion:        169,
		Anio:            2014,
		Clasificacion:   "PG-13",
		GeneroID:        "GEN001",
		CategoriaID:     "CAT001",
		ProveedorID:     "PROV001",
		VideoID:         "VID001",
		URLReproduccion: "https://streaminggo.example.com/video/VID001",
	}

	repo.Contenidos["CONT002"] = models.Contenido{
		ContenidoID:     "CONT002",
		Titulo:          "Avengers",
		Sinopsis:        "Un grupo de héroes se reúne para enfrentar una amenaza.",
		Duracion:        143,
		Anio:            2012,
		Clasificacion:   "PG-13",
		GeneroID:        "GEN002",
		CategoriaID:     "CAT001",
		ProveedorID:     "PROV001",
		VideoID:         "VID002",
		URLReproduccion: "https://streaminggo.example.com/video/VID002",
	}

	// HISTORIAL
	repo.Historiales["HIST001"] = models.Historial{
		HistorialID:        "HIST001",
		PerfilID:           "PER001",
		ContenidoID:        "CONT001",
		Progreso:           45,
		TiempoActual:       76,
		FechaVisualizacion: time.Now(),
	}
}
