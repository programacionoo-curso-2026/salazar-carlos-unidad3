package services

import (
	"errors"
	"streaming-go/models"
	"streaming-go/repository"
)

type ValoracionService struct {
	repo *repository.Memoria
}

func NuevaValoracionService(repo *repository.Memoria) *ValoracionService {
	return &ValoracionService{repo: repo}
}

func (s *ValoracionService) Crear(v models.Valoracion) error {
	if v.ValoracionID == "" {
		return errors.New("el ID es obligatorio")
	}

	if v.Puntuacion < 1 || v.Puntuacion > 5 {
		return errors.New("la puntuación debe estar entre 1 y 5")
	}

	if _, existe := s.repo.Valoraciones[v.ValoracionID]; existe {
		return errors.New("la valoración ya existe")
	}

	s.repo.Valoraciones[v.ValoracionID] = v
	return nil
}

func (s *ValoracionService) Listar() []models.Valoracion {
	resultado := make([]models.Valoracion, 0)

	for _, valoracion := range s.repo.Valoraciones {
		resultado = append(resultado, valoracion)
	}

	return resultado
}
