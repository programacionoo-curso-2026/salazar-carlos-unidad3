package main

import (
	"CompetenciasDocentes/dao"
	"CompetenciasDocentes/dataaccess"
	"log"
)

func main() {
	// Inicializar la base de datos
	db := dataaccess.InitDB()
	defer db.Close() // Importante: cerrar la conexión al final

	log.Println("Base de datos inicializada correctamente")

	// Crear el DAO
	docenteDAO := dao.NewDocenteDAO(db)

	// Crear la tabla
	if err := docenteDAO.CreateTable(); err != nil {
		log.Fatalf("Error al crear tabla: %v", err)
	}
}
```
