package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//RegTiempo tabla de registros de tiempo
type RegTiempo struct {
	gorm.Model
	EstacionID  uint
	EquipoID    uint
	HoraLlegada time.Time `json:"HoraLlegada"`
}
