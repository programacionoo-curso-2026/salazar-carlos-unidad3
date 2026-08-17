package services

import (
	"errors"
	"streaming-go/models"
	"streaming-go/repository"
)

type PerfilService struct {
	repo *repository.Memoria
}

func NuevoPerfilService(repo *repository.Memoria) *PerfilService {
	return &PerfilService{repo: repo}
}

func (s *PerfilService) Crear(p models.Perfil) error {
	if p.PerfilID == "" {
		return errors.New("el ID del perfil es obligatorio")
	}

	if p.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}

	if _, existe := s.repo.Perfiles[p.PerfilID]; existe {
		return errors.New("el perfil ya existe")
	}

	if _, existe := s.repo.Cuentas[p.CuentaID]; !existe {
		return errors.New("la cuenta no existe")
	}

	s.repo.Perfiles[p.PerfilID] = p

	return nil
}

func (s *PerfilService) Listar() []models.Perfil {
	perfiles := make([]models.Perfil, 0)

	for _, perfil := range s.repo.Perfiles {
		perfiles = append(perfiles, perfil)
	}

	return perfiles
}
