package preguntas

import "gorm.io/gorm"

// store metodos disponibles para Pregunta
type store interface {
	Pregunta(id uint) (*Pregunta, error)
	Preguntas() (*[]Pregunta, error)
	CreatePregunta(p *Pregunta) (*Pregunta, error)
	UpdatePregunta(id uint, p *Pregunta) (*Pregunta, error)
	DeletePregunta(id uint) error
}

// PreguntaService implementando metodos
type preguntaStore struct {
	db *gorm.DB
}

// Pregunta retorna la pregunta pedida en la BD
func (s *preguntaStore) Pregunta(id uint) (*Pregunta, error) {
	var pregunta Pregunta
	s.db.First(&pregunta, id)
	return &pregunta, nil
}

// Preguntas retorna todas las preguntas de la BD
func (s *preguntaStore) Preguntas() (*[]Pregunta, error) {
	var preguntas []Pregunta
	s.db.Find(&preguntas)
	return &preguntas, nil
}

// CreatePregunta agrega una pregunta a la BD
func (s *preguntaStore) CreatePregunta(p *Pregunta) (*Pregunta, error) {
	s.db.Create(p)
	return p, nil
}

// UpdatePregunta actualiza una pregunta en la BD
func (s *preguntaStore) UpdatePregunta(id uint, p *Pregunta) (*Pregunta, error) {
	var pregunta Pregunta
	s.db.Find(&pregunta, id)
	s.db.Model(&pregunta).Updates(p)
	return &pregunta, nil
}

// DeletePregunta elimina una pregunta de la BD
func (s *preguntaStore) DeletePregunta(id uint) error {
	s.db.Delete(&Pregunta{}, id)
	return nil
}
