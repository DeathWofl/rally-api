package mysql

import (
	"github.com/DeathWofl/rally-api/pkg/models"
	"gorm.io/gorm"
)

type PreguntaService struct {
	DB *gorm.DB
}

// Pregunta retorna la pregunta pedida en la BD
func (s *PreguntaService) Pregunta(id uint) (*models.Pregunta, error) {
	var pregunta models.Pregunta
	s.DB.First(&pregunta, id)
	return &pregunta, nil
}

// Preguntas retorna todas las preguntas de la BD
func (s *PreguntaService) Preguntas() (*[]models.Pregunta, error) {
	var preguntas []models.Pregunta
	s.DB.Find(&preguntas)
	return &preguntas, nil
}

// CreatePregunta agrega una pregunta a la BD
func (s *PreguntaService) CreatePregunta(p *models.Pregunta) (*models.Pregunta, error) {
	s.DB.Create(p)
	return p, nil
}

// UpdatePregunta actualiza una pregunta en la BD
func (s *PreguntaService) UpdatePregunta(id uint, p *models.Pregunta) (*models.Pregunta, error) {
	var pregunta models.Pregunta
	s.DB.Find(&pregunta, id)
	s.DB.Model(&pregunta).Updates(p)
	return &pregunta, nil
}

// DeletePregunta elimina una pregunta de la BD
func (s *PreguntaService) DeletePregunta(id uint) error {
	s.DB.Delete(&models.Pregunta{}, id)
	return nil
}
