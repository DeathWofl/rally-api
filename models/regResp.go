package models

import "github.com/jinzhu/gorm"

//RegResp Registros de las respuestas correctas
type RegResp struct {
	gorm.Model
	PreguntaID   uint
	Calificacion uint `json:"Calificacion"`
	EquipoID     uint
}
