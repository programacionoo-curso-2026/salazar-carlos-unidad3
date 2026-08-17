package services

import (
	"testing"

	"streaming-go/models"
	"streaming-go/repository"
)

func TestCrearCuenta(t *testing.T) {
	repo := repository.NuevaMemoria()
	service := NuevaCuentaService(repo)

	cuenta := models.Cuenta{
		CuentaID: "C001",
		Email:    "usuario@gmail.com",
		Estado:   "activo",
		PlanID:   "PLAN01",
	}

	err := service.Crear(cuenta)

	if err != nil {
		t.Fatalf("No se esperaba un error: %v", err)
	}

	resultado, err := service.Obtener("C001")

	if err != nil {
		t.Fatalf("No se pudo obtener la cuenta: %v", err)
	}

	if resultado.Email != "usuario@gmail.com" {
		t.Errorf("Email incorrecto")
	}
}
func TestCuentaDuplicada(t *testing.T) {
	repo := repository.NuevaMemoria()
	service := NuevaCuentaService(repo)

	cuenta := models.Cuenta{
		CuentaID: "C001",
		Email:    "usuario@gmail.com",
	}

	service.Crear(cuenta)

	err := service.Crear(cuenta)

	if err == nil {
		t.Error("Se esperaba error por cuenta duplicada")
	}
}
