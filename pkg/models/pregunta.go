package models

import "github.com/jinzhu/gorm"

//Pregunta representa cada pregunta del proceso.
type Pregunta struct {
	gorm.Model
	Preg       string `json:"Preg" gorm:"type:varchar(250);not null"`
	Respuestas []Respuesta
	RegResps   []RegResp
	EstacionID uint
}

// PreguntaService metodos disponibles para Pregunta
type PreguntaService interface {
	Pregunta(id uint) (*Pregunta, error)
	Preguntas() (*[]Pregunta, error)
	CreatePregunta(p *Pregunta) (*Pregunta, error)
	UpdatePregunta(id uint, p *Pregunta) (*Pregunta, error)
	DeletePregunta(id uint) error
}
