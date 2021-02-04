package mysql

import (
	"github.com/DeathWofl/rally-api/pkg/models"
	"github.com/jinzhu/gorm"
)

type EstacionService struct {
	DB *gorm.DB
}

// Estacion busca el equipo solicitado en la BD
func (s *EstacionService) Estacion(id uint) (*models.Estacion, error) {
	var estacion models.Estacion
	s.DB.First(&estacion, id)
	return &estacion, nil
}

// Equipos retorna todos los equipos en la BD
func (s *EstacionService) Equipos() (*[]models.Estacion, error) {
	var estaciones []models.Estacion
	s.DB.Find(&estaciones)
	return &estaciones, nil
}

// CreateEstacion agreaga un nuevo equipo a la BD
func (s *EstacionService) CreateEstacion(e *models.Estacion) (*models.Estacion, error) {
	s.DB.Create(e)
	return e, nil
}

// UpdateEstacion actualiza una estacion en la BD
func (s *EstacionService) UpdateEstacion(id uint, e *models.Estacion) (*models.Estacion, error) {
	var estacion models.Estacion
	s.DB.Find(&estacion, id)
	s.DB.Model(&estacion).Updates(e)
	return &estacion, nil
}

// DeleteEstacion elimina una estacion de la BD
func (s *EstacionService) DeleteEstacion(id uint) error {
	s.DB.Delete(&models.Estacion{}, id)
	return nil
}
