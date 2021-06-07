package estaciones

import (
	"gorm.io/gorm"
)

type store interface {
	ReadByID(id uint) (*Estacion, error)
	All() (*[]Estacion, error)
	Create(e *Estacion) (*Estacion, error)
	Update(id uint, e *Estacion) (*Estacion, error)
	Delete(id uint) error
}

type estacionStore struct {
	db *gorm.DB
}

func (s *estacionStore) ReadByID(id uint) (*Estacion, error) {
	var estacion Estacion
	s.db.First(&estacion, id)
	return &estacion, nil
}

func (s *estacionStore) All() (*[]Estacion, error) {
	var estaciones []Estacion
	s.db.Find(&estaciones)
	return &estaciones, nil
}

func (s *estacionStore) Create(e *Estacion) (*Estacion, error) {
	s.db.Create(e)
	return e, nil
}

func (s *estacionStore) Update(id uint, e *Estacion) (*Estacion, error) {
	var estacion Estacion
	s.db.Find(&estacion, id)
	s.db.Model(&estacion).Updates(e)
	return &estacion, nil
}

func (s *estacionStore) Delete(id uint) error {
	s.db.Delete(&Estacion{}, id)
	return nil
}

func newStore(db *gorm.DB) (*estacionStore, error) {
	return &estacionStore{
		db: db,
	}, nil
}
