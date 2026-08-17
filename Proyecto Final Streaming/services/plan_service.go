package services

import (
	"errors"
	"streaming-go/models"
	"streaming-go/repository"
)

type PlanService struct {
	repo *repository.Memoria
}

func NuevoPlanService(repo *repository.Memoria) *PlanService {
	return &PlanService{repo: repo}
}

func (s *PlanService) Crear(p models.Plan) error {
	if p.PlanID == "" {
		return errors.New("el ID del plan es obligatorio")
	}

	if p.Nombre == "" {
		return errors.New("el nombre del plan es obligatorio")
	}

	if p.Precio < 0 {
		return errors.New("el precio no puede ser negativo")
	}

	if _, existe := s.repo.Planes[p.PlanID]; existe {
		return errors.New("el plan ya existe")
	}

	s.repo.Planes[p.PlanID] = p

	return nil
}

func (s *PlanService) Listar() []models.Plan {
	planes := make([]models.Plan, 0)

	for _, plan := range s.repo.Planes {
		planes = append(planes, plan)
	}

	return planes
}
