package estaciones

import (
	"gorm.io/gorm"
)

// estacionStore metodos disponibles para Estacion
type store interface {
	Estacion(id uint) (*Estacion, error)
	Estaciones() (*[]Estacion, error)
	CreateEstacion(e *Estacion) (*Estacion, error)
	UpdateEstacion(id uint, e *Estacion) (*Estacion, error)
	DeleteEstacion(id uint) error
}

// estacionStore implementando metodos
type estacionStore struct {
	db *gorm.DB
}

// Estacion busca el equipo solicitado en la BD
func (s *estacionStore) Estacion(id uint) (*Estacion, error) {
	var estacion Estacion
	s.db.First(&estacion, id)
	return &estacion, nil
}

// Estaciones retorna todos los equipos en la BD
func (s *estacionStore) Estaciones() (*[]Estacion, error) {
	var estaciones []Estacion
	s.db.Find(&estaciones)
	return &estaciones, nil
}

// CreateEstacion agreaga un nuevo equipo a la BD
func (s *estacionStore) CreateEstacion(e *Estacion) (*Estacion, error) {
	s.db.Create(e)
	return e, nil
}

// UpdateEstacion actualiza una estacion en la BD
func (s *estacionStore) UpdateEstacion(id uint, e *Estacion) (*Estacion, error) {
	var estacion Estacion
	s.db.Find(&estacion, id)
	s.db.Model(&estacion).Updates(e)
	return &estacion, nil
}

// DeleteEstacion elimina una estacion de la BD
func (s *estacionStore) DeleteEstacion(id uint) error {
	s.db.Delete(&Estacion{}, id)
	return nil
}
