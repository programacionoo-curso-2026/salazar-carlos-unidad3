package services

import (
	"errors"
	"streaming-go/models"
	"streaming-go/repository"
)

type FavoritoService struct {
	repo *repository.Memoria
}

func NuevoFavoritoService(repo *repository.Memoria) *FavoritoService {
	return &FavoritoService{repo: repo}
}

func (s *FavoritoService) Crear(f models.Favorito) error {
	if f.FavoritoID == "" {
		return errors.New("el ID es obligatorio")
	}

	if _, existe := s.repo.Favoritos[f.FavoritoID]; existe {
		return errors.New("el favorito ya existe")
	}

	if _, existe := s.repo.Perfiles[f.PerfilID]; !existe {
		return errors.New("el perfil no existe")
	}

	if _, existe := s.repo.Contenidos[f.ContenidoID]; !existe {
		return errors.New("el contenido no existe")
	}

	s.repo.Favoritos[f.FavoritoID] = f

	return nil
}

func (s *FavoritoService) Listar() []models.Favorito {
	favoritos := make([]models.Favorito, 0)

	for _, favorito := range s.repo.Favoritos {
		favoritos = append(favoritos, favorito)
	}

	return favoritos
}
