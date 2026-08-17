package services

import (
	"errors"
	"streaming-go/models"
	"streaming-go/repository"
)

type CuentaService struct {
	repo *repository.Memoria
}

func NuevaCuentaService(repo *repository.Memoria) *CuentaService {
	return &CuentaService{repo: repo}
}

func (s *CuentaService) Crear(c models.Cuenta) error {
	if c.CuentaID == "" {
		return errors.New("el ID de la cuenta es obligatorio")
	}

	if c.Email == "" {
		return errors.New("el email es obligatorio")
	}

	if _, existe := s.repo.Cuentas[c.CuentaID]; existe {
		return errors.New("la cuenta ya existe")
	}

	s.repo.Cuentas[c.CuentaID] = c

	return nil
}

func (s *CuentaService) Obtener(id string) (models.Cuenta, error) {
	cuenta, existe := s.repo.Cuentas[id]

	if !existe {
		return models.Cuenta{}, errors.New("cuenta no encontrada")
	}

	return cuenta, nil
}

func (s *CuentaService) Listar() []models.Cuenta {
	cuentas := make([]models.Cuenta, 0)

	for _, cuenta := range s.repo.Cuentas {
		cuentas = append(cuentas, cuenta)
	}

	return cuentas
}
