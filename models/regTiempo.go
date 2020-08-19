package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//RegTiempo tabla de registros de tiempo
type RegTiempo struct {
	gorm.Model
	EstacionID  uint      `json:"EstacionID"`
	EquipoID    uint      `json:"EquipoID"`
	CodigoGrupo string `json:"CodigoGrupo" gorm:"not null;type:varchar(100);"`
	HoraLlegada time.Time `json:"HoraLlegada"`
}
