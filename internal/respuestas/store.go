package respuestas

import (
	"gorm.io/gorm"
)

type store interface {
	ReadByID(id uint) (*Respuesta, error)
	All() (*[]Respuesta, error)
	Create(r *Respuesta) (*Respuesta, error)
	Update(id uint, r *Respuesta) (*Respuesta, error)
	Delete(id uint) error
}

type respuestaStore struct {
	db *gorm.DB
}

func (s *respuestaStore) ReadByID(id uint) (*Respuesta, error) {
	var respuesta Respuesta
	s.db.First(&respuesta, id)
	return &respuesta, nil
}

func (s *respuestaStore) All() (*[]Respuesta, error) {
	var respuestas []Respuesta
	s.db.Find(&respuestas)
	return &respuestas, nil
}

func (s *respuestaStore) Create(r *Respuesta) (*Respuesta, error) {
	s.db.Create(r)
	return r, nil
}

func (s *respuestaStore) Update(id uint, r *Respuesta) (*Respuesta, error) {
	var respuesta Respuesta
	s.db.Find(&respuesta, id)
	s.db.Model(&respuesta).Updates(r)
	return &respuesta, nil
}

func (s *respuestaStore) Delete(id uint) error {
	s.db.Delete(&Respuesta{}, id)
	return nil
}

func newStore(db *gorm.DB) (*respuestaStore, error) {
	return &respuestaStore{
		db: db,
	}, nil
}
