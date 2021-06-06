package respuestas

import (
	"gorm.io/gorm"
)

// service metodos disponibles para Respuesta
type store interface {
	Respuesta(id uint) (*Respuesta, error)
	Respuestas() (*[]Respuesta, error)
	CreateRespuesta(r *Respuesta) (*Respuesta, error)
	UpdateRespuesta(id uint, r *Respuesta) (*Respuesta, error)
	DeleteRespuesta(id uint) error
}

// respuestaService implementando metodos
type respuestaStore struct {
	db *gorm.DB
}

// Respuesta retorna la respuesta pedida en la BD
func (s *respuestaStore) Respuesta(id uint) (*Respuesta, error) {
	var respuesta Respuesta
	s.db.First(&respuesta, id)
	return &respuesta, nil
}

// Respuestas retorna todas las respuestas de la BD
func (s *respuestaStore) Respuestas() (*[]Respuesta, error) {
	var respuestas []Respuesta
	s.db.Find(&respuestas)
	return &respuestas, nil
}

// CreateRespuesta agrega una respuesta a la BD
func (s *respuestaStore) CreateRespuesta(r *Respuesta) (*Respuesta, error) {
	s.db.Create(r)
	return r, nil
}

// UpdateRespuesta actualiza una respuesta en la BD
func (s *respuestaStore) UpdateRespuesta(id uint, r *Respuesta) (*Respuesta, error) {
	var respuesta Respuesta
	s.db.Find(&respuesta, id)
	s.db.Model(&respuesta).Updates(r)
	return &respuesta, nil
}

// DeleteRespuesta elimina una respuesta de la BD
func (s *respuestaStore) DeleteRespuesta(id uint) error {
	s.db.Delete(&Respuesta{}, id)
	return nil
}

func newService(db *gorm.DB) (*respuestaStore, error) {
	return &respuestaStore{
		db: db,
	}, nil
}
