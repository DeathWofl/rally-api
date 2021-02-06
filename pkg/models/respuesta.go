package models

import "github.com/jinzhu/gorm"

//Respuesta representa cada respuesta correspodiente a una pregunta
type Respuesta struct {
	gorm.Model
	Resp       string `json:"Resp" gorm:"type:varchar(200);not null"`
	Valor      int
	PreguntaID uint
	Pregunta   Pregunta
}

// RespuestaService metodos disponlibles para Respuesta
type RespuestaService interface {
	Respuesta(id uint) (*Respuesta, error)
	Respuestas() (*[]Respuesta, error)
	CreateRespuesta(r *Respuesta) (*Respuesta, error)
	UpdateRespuesta(id uint, r *Respuesta) (*Respuesta, error)
	DeleteRespuesta(id uint) error
}
