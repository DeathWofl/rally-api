package models

import "github.com/jinzhu/gorm"

//Respuesta tabla de respuestas
type Respuesta struct {
	gorm.Model
	Resp       string `json:"Resp" gorm:"type:varchar(200);not null"`
	Valor      int
	PreguntaID uint
	Preg       Pregunta
}
