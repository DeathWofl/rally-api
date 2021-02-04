package models

import "github.com/jinzhu/gorm"

//Estacion tabla de estaciones del rally
type Estacion struct {
	gorm.Model
	Nombre      string `json:"nombre" gorm:"not null;type:varchar(20)"`
	Description string `json:"Descripcion" gorm:"type:varchar(250)"`
	Preguntas   []Pregunta
	RegTiempos  []RegTiempo
}

// EstacionService metodos disponibles para Estacion
type EstacionService interface {
	Estacion(id uint) (*Equipo, error)
	Equipos() (*[]Equipo, error)
	CreateEstacion(e *Estacion) (*Estacion, error)
	UpdateEstacion(id uint, e *Estacion) (*Estacion, error)
	DeleteEstacion(id uint) error
}
