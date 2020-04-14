package models

import "github.com/jinzhu/gorm"

//RegResp Registros de las respuestas correctas
type RegResp struct {
	gorm.Model
	Pregunta     Pregunta
	Calificacion int `json:"Calificacion"`
	Equipo       Equipo
}
