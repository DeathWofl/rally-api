package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//RegTiempo tabla de registros de tiempo
type RegTiempo struct {
	gorm.Model
	EstacionID   uint      `json:"EstacionID"`
	EquipoID     uint      `json:"EquipoID"`
	EquipoCodigo string    `json:"EquipoCodigo"`
	HoraLlegada  time.Time `json:"HoraLlegada"`
}
