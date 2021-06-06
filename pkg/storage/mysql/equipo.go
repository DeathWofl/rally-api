// Package mysql representa todas las interraciones que
// se haran entre la base de datos y la appp
// el nombre del paquete hace referencia al motor de base de datos
package mysql

import (
	"fmt"

	"github.com/spinales/rally-api/pkg/models"
	"gorm.io/gorm"
)

// EquipoService implementado metodos
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
	// Busqueda para evitar que se repita una matricula
	result := models.Equipo{}

	// Confirma que la matricula de ese estudiante no este registrada con anterioridad en la BD
	// Especificamente del estudiante de la matricula 1
	s.DB.Where(models.Equipo{MatriculaE1: u.MatriculaE1}).
		Or(models.Equipo{MatriculaE2: u.MatriculaE1}).
		Or(models.Equipo{MatriculaE3: u.MatriculaE1}).
		Find(&result)
	if result.MatriculaE1 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 1 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 2
	s.DB.Where(models.Equipo{MatriculaE1: u.MatriculaE2}).
		Or(models.Equipo{MatriculaE2: u.MatriculaE2}).
		Or(models.Equipo{MatriculaE3: u.MatriculaE2}).
		Find(&result)
	if result.MatriculaE2 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 2 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 3
	s.DB.Where(models.Equipo{MatriculaE1: u.MatriculaE3}).
		Or(models.Equipo{MatriculaE2: u.MatriculaE3}).
		Or(models.Equipo{MatriculaE3: u.MatriculaE3}).
		Find(&result)
	if result.MatriculaE1 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 3 ya esta registrada")
	}
	s.DB.Create(u)
	return u, nil
}

// UpdateEquipo actualiza un equipo en la BD
func (s *EquipoService) UpdateEquipo(id uint, u *models.Equipo) (*models.Equipo, error) {
	var equipo models.Equipo

	validequipo := models.Equipo{}
	s.DB.Where(models.Equipo{MatriculaE1: u.MatriculaE1}).
		Or(models.Equipo{MatriculaE2: u.MatriculaE1}).
		Or(models.Equipo{MatriculaE3: u.MatriculaE1}).
		Not(id).
		Find(&validequipo)
	if validequipo.MatriculaE1 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 1 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 2
	s.DB.Where(models.Equipo{MatriculaE1: u.MatriculaE2}).
		Or(models.Equipo{MatriculaE2: u.MatriculaE2}).
		Or(models.Equipo{MatriculaE3: u.MatriculaE2}).
		Not(id).
		Find(&validequipo)
	if validequipo.MatriculaE2 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 2 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 3
	s.DB.Where(models.Equipo{MatriculaE1: u.MatriculaE3}).
		Or(models.Equipo{MatriculaE2: u.MatriculaE3}).
		Or(models.Equipo{MatriculaE3: u.MatriculaE3}).
		Not(id).
		Find(&validequipo)
	if validequipo.MatriculaE1 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 3 ya esta registrada")
	}

	s.DB.Find(&equipo, id)
	s.DB.Model(&equipo).Updates(u)
	return &equipo, nil
}

// DeleteEquipo elimina un equipo de la BD
func (s *EquipoService) DeleteEquipo(id uint) error {
	s.DB.Delete(&models.Equipo{}, id)
	return nil
}
