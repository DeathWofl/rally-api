package models

import "github.com/jinzhu/gorm"

//Pregunta tabla de preguntas
type Pregunta struct {
	gorm.Model
	Preg       string `json:"Preg" gorm:"type:varchar(250);not null"`
	Respuestas []Respuesta
	RegResps   []RegResp
	EstacionID uint
}
