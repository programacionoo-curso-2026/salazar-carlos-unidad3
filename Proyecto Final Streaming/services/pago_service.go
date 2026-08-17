package services

import (
	"errors"
	"streaming-go/models"
	"streaming-go/repository"
)

type PagoService struct {
	repo *repository.Memoria
}

func NuevoPagoService(repo *repository.Memoria) *PagoService {
	return &PagoService{repo: repo}
}

func (s *PagoService) Crear(p models.Pago) error {
	if p.PagoID == "" {
		return errors.New("el ID del pago es obligatorio")
	}

	if p.Monto <= 0 {
		return errors.New("el monto debe ser mayor que cero")
	}

	if _, existe := s.repo.Cuentas[p.CuentaID]; !existe {
		return errors.New("la cuenta no existe")
	}

	s.repo.Pagos[p.PagoID] = p

	return nil
}

func (s *PagoService) Listar() []models.Pago {
	pagos := make([]models.Pago, 0)

	for _, pago := range s.repo.Pagos {
		pagos = append(pagos, pago)
	}

	return pagos
}
