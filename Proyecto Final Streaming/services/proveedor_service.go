package services

import (
	"errors"
	"streaming-go/models"
	"streaming-go/repository"
)

type ProveedorService struct {
	repo *repository.Memoria
}

func NuevoProveedorService(repo *repository.Memoria) *ProveedorService {
	return &ProveedorService{repo: repo}
}

func (s *ProveedorService) Crear(p models.ProveedorVideo) error {
	if p.ProveedorID == "" {
		return errors.New("el ID del proveedor es obligatorio")
	}

	if p.NombreProveedor == "" {
		return errors.New("el nombre del proveedor es obligatorio")
	}

	if _, existe := s.repo.Proveedores[p.ProveedorID]; existe {
		return errors.New("el proveedor ya existe")
	}

	s.repo.Proveedores[p.ProveedorID] = p

	return nil
}

func (s *ProveedorService) Listar() []models.ProveedorVideo {
	proveedores := make([]models.ProveedorVideo, 0)

	for _, proveedor := range s.repo.Proveedores {
		proveedores = append(proveedores, proveedor)
	}

	return proveedores
}
