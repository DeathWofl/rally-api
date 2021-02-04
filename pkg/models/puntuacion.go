package models

import (
	"github.com/jinzhu/gorm"
)

// Puntuacion representa cada puntuacion de cada equipo
// antiguamente representado como ganador
type Puntuacion struct {
	gorm.Model
	Transcurso string
	Puntaje    float32
}

// PuntuacionService metodos disponibles para Puntuacion
type PuntuacionService interface {
	Puntuaciones() (*[]Puntuacion, error)
}
