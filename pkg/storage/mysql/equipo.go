// Package mysql representa todas las interraciones que
// se haran entre la base de datos y la appp
// el nombre del paquete hace referencia al motor de base de datos
package mysql

import (
	"github.com/DeathWofl/rally-api/pkg/models"
	"gorm.io/gorm"
)

type EquipoService struct {
	DB *gorm.DB
}

// BD = base de datos

// Equipo trae un equipo por el id pedido
func (s *EquipoService) Equipo(id uint) (*models.Equipo, error) {
	var equipo models.Equipo
	s.DB.First(&equipo, id)
	return &equipo, nil
}

// Equipos retorna todos los equipos que existan en la BD
func (s *EquipoService) Equipos() (*[]models.Equipo, error) {
	var equipos []models.Equipo
	s.DB.Find(&equipos)
	return &equipos, nil
}

// CreateEquipo agrega un nuevo equipo a la BD
func (s *EquipoService) CreateEquipo(u *models.Equipo) (*models.Equipo, error) {
	s.DB.Create(u)
	return u, nil
}

// UpdateEquipo actualiza un equipo en la BD
func (s *EquipoService) UpdateEquipo(id uint, u *models.Equipo) (*models.Equipo, error) {
	var equipo models.Equipo
	s.DB.Find(&equipo, id)
	s.DB.Model(&equipo).Updates(u)
	return &equipo, nil
}

// DeleteEquipo elimina un equipo de la BD
func (s *EquipoService) DeleteEquipo(id uint) error {
	s.DB.Delete(&models.Equipo{}, id)
	return nil
}
