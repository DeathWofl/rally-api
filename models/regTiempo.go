package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//RegTiempo tabla de registros de tiempo
type RegTiempo struct {
	gorm.Model
	Estacion    Estacion
	Equipo      Equipo
	HoraLlegada time.Time `json:"HoraLlegada"`
}
