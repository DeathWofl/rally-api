package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//RegTiempo representa el horaria que un equipo llego a una estacion
// RegTiempo = Registro de Tiempo
type RegTiempo struct {
	gorm.Model
	EstacionID  uint      `json:"EstacionID"`
	EquipoID    uint      `json:"EquipoID"`
	CodigoGrupo string    `json:"CodigoGrupo" gorm:"not null;type:varchar(100);"`
	HoraLlegada time.Time `json:"HoraLlegada"`
}

// RegTiempoService metodos disponibles para RegTiempo
type RegTiempoService interface {
	RegTiempo(id uint) (*RegTiempo, error)
	RegTiempos() (*[]RegTiempo, error)
	CreateRegTiempo(r *RegTiempo) (*RegTiempo, error)
	UpdateRegTiempo(id uint, r *RegTiempo) (*RegTiempo, error)
	DeleteRegTiempo(id uint) error
}
