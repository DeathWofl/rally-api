package equipos

import (
	"fmt"

	"gorm.io/gorm"
)

type store interface {
	ReadByID(id uint) (*Equipo, error)
	All() (*[]Equipo, error)
	Create(u *Equipo) (*Equipo, error)
	Update(id uint, u *Equipo) (*Equipo, error)
	Delete(id uint) error
}

type equipoStore struct {
	db *gorm.DB
}

func (s *equipoStore) ReadByID(id uint) (*Equipo, error) {
	var equipo Equipo
	s.db.First(&equipo, id)
	return &equipo, nil
}

func (s *equipoStore) All() (*[]Equipo, error) {
	var equipos []Equipo
	s.db.Find(&equipos)
	return &equipos, nil
}

func (s *equipoStore) Create(u *Equipo) (*Equipo, error) {
	// Busqueda para evitar que se repita una matricula
	result := Equipo{}

	// Confirma que la matricula de ese estudiante no este registrada con anterioridad en la BD
	// Especificamente del estudiante de la matricula 1
	s.db.Where(Equipo{MatriculaE1: u.MatriculaE1}).
		Or(Equipo{MatriculaE2: u.MatriculaE1}).
		Or(Equipo{MatriculaE3: u.MatriculaE1}).
		Find(&result)
	if result.MatriculaE1 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 1 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 2
	s.db.Where(Equipo{MatriculaE1: u.MatriculaE2}).
		Or(Equipo{MatriculaE2: u.MatriculaE2}).
		Or(Equipo{MatriculaE3: u.MatriculaE2}).
		Find(&result)
	if result.MatriculaE2 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 2 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 3
	s.db.Where(Equipo{MatriculaE1: u.MatriculaE3}).
		Or(Equipo{MatriculaE2: u.MatriculaE3}).
		Or(Equipo{MatriculaE3: u.MatriculaE3}).
		Find(&result)
	if result.MatriculaE1 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 3 ya esta registrada")
	}
	s.db.Create(u)
	return u, nil
}

func (s *equipoStore) Update(id uint, u *Equipo) (*Equipo, error) {
	var equipo Equipo

	validequipo := Equipo{}
	s.db.Where(Equipo{MatriculaE1: u.MatriculaE1}).
		Or(Equipo{MatriculaE2: u.MatriculaE1}).
		Or(Equipo{MatriculaE3: u.MatriculaE1}).
		Not(id).
		Find(&validequipo)
	if validequipo.MatriculaE1 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 1 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 2
	s.db.Where(Equipo{MatriculaE1: u.MatriculaE2}).
		Or(Equipo{MatriculaE2: u.MatriculaE2}).
		Or(Equipo{MatriculaE3: u.MatriculaE2}).
		Not(id).
		Find(&validequipo)
	if validequipo.MatriculaE2 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 2 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 3
	s.db.Where(Equipo{MatriculaE1: u.MatriculaE3}).
		Or(Equipo{MatriculaE2: u.MatriculaE3}).
		Or(Equipo{MatriculaE3: u.MatriculaE3}).
		Not(id).
		Find(&validequipo)
	if validequipo.MatriculaE1 != "" {
		return nil, fmt.Errorf("La Matricula del Estudiante 3 ya esta registrada")
	}

	s.db.Find(&equipo, id)
	s.db.Model(&equipo).Updates(u)
	return &equipo, nil
}

func (s *equipoStore) Delete(id uint) error {
	s.db.Delete(&Equipo{}, id)
	return nil
}

func newStore(db *gorm.DB) (*equipoStore, error) {
	return &equipoStore{
		db: db,
	}, nil
}
