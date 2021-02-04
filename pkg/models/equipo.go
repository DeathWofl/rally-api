// Package models define cada entidad que existe en la app
// tambien los metodos correspondientes a cada entidad
package models

import "github.com/jinzhu/gorm"

//Equipo representa cada equipo del rally
type Equipo struct {
	gorm.Model
	MatriculaE1 string `json:"MatriculaE1" gorm:"not null;type:varchar(15);"`
	MatriculaE2 string `json:"MatriculaE2" gorm:"not null;type:varchar(15);"`
	MatriculaE3 string `json:"MatriculaE3" gorm:"not null;type:varchar(15);"`
	CodigoGrupo string `json:"CodigoGrupo" gorm:"not null;type:varchar(100);" sql:"index"`
	ContraGrupo string `json:"ContraGrupo" gorm:"not null;type:varchar(100);"`
	RegResps    []RegResp
	RegTiempos  []RegTiempo
}

// EquipoService metodos disponibles para Equipo
type EquipoService interface {
	Equipo(id uint) (*Equipo, error)
	Equipos() (*[]Equipo, error)
	CreateEquipo(u *Equipo) (*Equipo, error)
	UpdateEquipo(id uint, u *Equipo) (*Equipo, error)
	DeleteEquipo(id uint) error
}
