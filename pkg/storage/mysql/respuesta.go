package mysql

import (
	"github.com/DeathWofl/rally-api/pkg/models"
	"gorm.io/gorm"
)

// RespuestaService implementando metodos
type RespuestaService struct {
	DB *gorm.DB
}

// Respuesta retorna la respuesta pedida en la BD
func (s *RespuestaService) Respuesta(id uint) (*models.Respuesta, error) {
	var respuesta models.Respuesta
	s.DB.First(&respuesta, id)
	return &respuesta, nil
}

// Respuestas retorna todas las respuestas de la BD
func (s *RespuestaService) Respuestas() (*[]models.Respuesta, error) {
	var respuestas []models.Respuesta
	s.DB.Find(&respuestas)
	return &respuestas, nil
}

// CreateRespuesta agrega una respuesta a la BD
func (s *RespuestaService) CreateRespuesta(r *models.Respuesta) (*models.Respuesta, error) {
	s.DB.Create(r)
	return r, nil
}

// UpdateRespuesta actualiza una respuesta en la BD
func (s *RespuestaService) UpdateRespuesta(id uint, r *models.Respuesta) (*models.Respuesta, error) {
	var respuesta models.Respuesta
	s.DB.Find(&respuesta, id)
	s.DB.Model(&respuesta).Updates(r)
	return &respuesta, nil
}

// DeleteRespuesta elimina una respuesta de la BD
func (s *RespuestaService) DeleteRespuesta(id uint) error {
	s.DB.Delete(&models.Respuesta{}, id)
	return nil
}
