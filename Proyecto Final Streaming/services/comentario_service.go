package services

import (
	"errors"
	"streaming-go/models"
	"streaming-go/repository"
)

type ComentarioService struct {
	repo *repository.Memoria
}

func NuevoComentarioService(repo *repository.Memoria) *ComentarioService {
	return &ComentarioService{repo: repo}
}

func (s *ComentarioService) Crear(c models.Comentario) error {
	if c.ComentarioID == "" {
		return errors.New("el ID es obligatorio")
	}

	if c.Texto == "" {
		return errors.New("el comentario no puede estar vacío")
	}

	if _, existe := s.repo.Comentarios[c.ComentarioID]; existe {
		return errors.New("el comentario ya existe")
	}

	s.repo.Comentarios[c.ComentarioID] = c
	return nil
}

func (s *ComentarioService) Listar() []models.Comentario {
	resultado := make([]models.Comentario, 0)

	for _, comentario := range s.repo.Comentarios {
		resultado = append(resultado, comentario)
	}

	return resultado
}
