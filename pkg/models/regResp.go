package models

import "github.com/jinzhu/gorm"

//RegResp representa cada vez que un equipo responde una
// pregunta con su puntuacion
// RegResp = Registro de Respuestas
type RegResp struct {
	gorm.Model
	PreguntaID   uint `json:"PreguntaID"`
	Calificacion uint `json:"Calificacion"`
	EquipoID     uint `json:"EquipoID"`
}

// RegRespService metodos disponibles para RegRes
type RegRespService interface {
	RegResp(id uint) (*RegResp, error)
	RegResps() (*[]RegResp, error)
	CreateRegResp(r *RegResp) (*RegResp, error)
	UpdateRegResp(id uint, r *RegResp) (*RegResp, error)
	DeleteRegResp(id uint) error
}
