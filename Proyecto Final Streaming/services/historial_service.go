package services

import (
	"errors"
	"streaming-go/models"
	"streaming-go/repository"
)

type HistorialService struct {
	repo *repository.Memoria
}

func NuevoHistorialService(repo *repository.Memoria) *HistorialService {
	return &HistorialService{repo: repo}
}

func (s *HistorialService) Crear(h models.Historial) error {
	if h.HistorialID == "" {
		return errors.New("el ID es obligatorio")
	}

	if h.Progreso < 0 || h.Progreso > 100 {
		return errors.New("el progreso debe estar entre 0 y 100")
	}

	if _, existe := s.repo.Historiales[h.HistorialID]; existe {
		return errors.New("el historial ya existe")
	}

	s.repo.Historiales[h.HistorialID] = h

	return nil
}
func (s *HistorialService) Listar() []models.Historial {
	resultado := make([]models.Historial, 0)

	for _, historial := range s.repo.Historiales {
		resultado = append(resultado, historial)
	}

	return resultado
}
