package services

import (
	"errors"

	"streaming-go/models"
	"streaming-go/repository"
)

type ContenidoService struct {
	repo *repository.Memoria
}

func NuevoContenidoService(repo *repository.Memoria) *ContenidoService {
	return &ContenidoService{repo: repo}
}

func (s *ContenidoService) Crear(c models.Contenido) error {

	if c.ContenidoID == "" {
		return errors.New("el ID es obligatorio")
	}

	if c.Titulo == "" {
		return errors.New("el título es obligatorio")
	}

	if _, existe := s.repo.Contenidos[c.ContenidoID]; existe {
		return errors.New("el contenido ya existe")
	}

	s.repo.Contenidos[c.ContenidoID] = c

	return nil
}

func (s *ContenidoService) Listar() []models.Contenido {

	resultado := make([]models.Contenido, 0)

	for _, contenido := range s.repo.Contenidos {
		resultado = append(resultado, contenido)
	}

	return resultado
}

func (s *ContenidoService) Obtener(id string) (models.Contenido, error) {

	contenido, existe := s.repo.Contenidos[id]

	if !existe {
		return models.Contenido{}, errors.New("contenido no encontrado")
	}

	return contenido, nil
}

func (s *ContenidoService) Eliminar(id string) error {

	if _, existe := s.repo.Contenidos[id]; !existe {
		return errors.New("contenido no encontrado")
	}

	delete(s.repo.Contenidos, id)

	return nil
}

func (s *ContenidoService) Actualizar(id string, nuevo models.Contenido) error {

	if _, existe := s.repo.Contenidos[id]; !existe {
		return errors.New("contenido no encontrado")
	}

	if nuevo.Titulo == "" {
		return errors.New("el título es obligatorio")
	}

	nuevo.ContenidoID = id

	s.repo.Contenidos[id] = nuevo

	return nil
}
