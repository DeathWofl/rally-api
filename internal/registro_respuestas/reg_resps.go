package regresp

import "gorm.io/gorm"

// RegResp representa cada vez que un equipo responde una
// pregunta con su puntuacion
// RegResp = Registro de Respuestas
type RegResp struct {
	gorm.Model
	PreguntaID   uint `json:"PreguntaID"`
	Calificacion uint `json:"Calificacion"`
	EquipoID     uint `json:"EquipoID"`
}
