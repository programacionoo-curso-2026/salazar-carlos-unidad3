package main

import (
	"fmt"
	"net/http"

	"streaming-go/handlers"
	"streaming-go/repository"
	"streaming-go/services"
)

func main() {

	// =========================
	// REPOSITORIO
	// =========================

	repo := repository.NuevaMemoria()
	repository.CargarDatosIniciales(repo)

	// =========================
	// SERVICIOS
	// =========================

	cuentaService := services.NuevaCuentaService(repo)
	perfilService := services.NuevoPerfilService(repo)
	contenidoService := services.NuevoContenidoService(repo)
	favoritoService := services.NuevoFavoritoService(repo)
	comentarioService := services.NuevoComentarioService(repo)
	valoracionService := services.NuevaValoracionService(repo)
	historialService := services.NuevoHistorialService(repo)
	planService := services.NuevoPlanService(repo)
	proveedorService := services.NuevoProveedorService(repo)
	pagoService := services.NuevoPagoService(repo)

	// =========================
	// HANDLERS
	// =========================

	cuentaHandler := handlers.NuevaCuentaHandler(cuentaService)
	perfilHandler := handlers.NuevoPerfilHandler(perfilService)
	contenidoHandler := handlers.NuevoContenidoHandler(contenidoService)
	favoritoHandler := handlers.NuevoFavoritoHandler(favoritoService)
	comentarioHandler := handlers.NuevoComentarioHandler(comentarioService)
	valoracionHandler := handlers.NuevaValoracionHandler(valoracionService)
	historialHandler := handlers.NuevoHistorialHandler(historialService)
	planHandler := handlers.NuevoPlanHandler(planService)
	proveedorHandler := handlers.NuevoProveedorHandler(proveedorService)
	pagoHandler := handlers.NuevoPagoHandler(pagoService)
	estadisticasHandler := handlers.NuevoEstadisticasHandler(repo)
	// =========================
	// CUENTAS
	// =========================

	http.HandleFunc("/api/cuentas", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			cuentaHandler.Listar(w, r)

		case http.MethodPost:
			cuentaHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/cuentas/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			cuentaHandler.Obtener(w, r)
			return
		}

		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	})

	// =========================
	// PERFILES
	// =========================

	http.HandleFunc("/api/perfiles", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			perfilHandler.Listar(w, r)

		case http.MethodPost:
			perfilHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	// CONTENIDOS
	// =========================

	http.HandleFunc("/api/contenidos", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			contenidoHandler.Listar(w, r)

		case http.MethodPost:
			contenidoHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/contenidos/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			contenidoHandler.Obtener(w, r)

		case http.MethodPut:
			contenidoHandler.Actualizar(w, r)

		case http.MethodDelete:
			contenidoHandler.Eliminar(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	// =========================
	// FAVORITOS
	// =========================

	http.HandleFunc("/api/favoritos", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			favoritoHandler.Listar(w, r)

		case http.MethodPost:
			favoritoHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	// COMENTARIOS
	// =========================

	http.HandleFunc("/api/comentarios", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			comentarioHandler.Listar(w, r)

		case http.MethodPost:
			comentarioHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	// VALORACIONES
	// =========================

	http.HandleFunc("/api/valoraciones", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			valoracionHandler.Listar(w, r)

		case http.MethodPost:
			valoracionHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	// HISTORIAL
	// =========================

	http.HandleFunc("/api/historial", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			historialHandler.Listar(w, r)

		case http.MethodPost:
			historialHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	// PLANES
	// =========================

	http.HandleFunc("/api/planes", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			planHandler.Listar(w, r)

		case http.MethodPost:
			planHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	// PROVEEDORES
	// =========================

	http.HandleFunc("/api/proveedores", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			proveedorHandler.Listar(w, r)

		case http.MethodPost:
			proveedorHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	// PAGOS
	// =========================

	http.HandleFunc("/api/pagos", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			pagoHandler.Listar(w, r)

		case http.MethodPost:
			pagoHandler.Crear(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	// =========================
	// ESTADÍSTICAS
	// =========================

	http.HandleFunc("/api/estadisticas", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		estadisticasHandler.Obtener(w, r)
	})

	// =========================
	// SERVIDOR
	// =========================

	fmt.Println("==========================================")
	fmt.Println("     PLATAFORMA DE STREAMING - GO")
	fmt.Println("==========================================")
	fmt.Println("Servidor: http://localhost:8080")
	fmt.Println("")
	fmt.Println("Servicios disponibles:")
	fmt.Println("GET/POST  /api/cuentas")
	fmt.Println("GET       /api/cuentas/{id}")
	fmt.Println("GET/POST  /api/perfiles")
	fmt.Println("GET/POST  /api/contenidos")
	fmt.Println("GET/POST  /api/favoritos")
	fmt.Println("GET/POST  /api/comentarios")
	fmt.Println("GET/POST  /api/valoraciones")
	fmt.Println("GET/POST  /api/historial")
	fmt.Println("GET/POST  /api/planes")
	fmt.Println("GET/POST  /api/proveedores")
	fmt.Println("GET/POST  /api/pagos")
	fmt.Println("GET       /api/estadisticas")
	fmt.Println("==========================================")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
