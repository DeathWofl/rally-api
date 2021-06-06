package regtieps

import (
	"time"

	"gorm.io/gorm"
)

// RegTiempo representa el horaria que un equipo llego a una estacion
// RegTiempo = Registro de Tiempo
type RegTiempo struct {
	gorm.Model
	EstacionID  uint      `json:"EstacionID"`
	EquipoID    uint      `json:"EquipoID"`
	CodigoGrupo string    `json:"CodigoGrupo" gorm:"not null;type:varchar(100);"`
	HoraLlegada time.Time `json:"HoraLlegada"`
}
