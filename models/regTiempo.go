package models

import "time"

//RegTiempo tabla de registros de tiempo
type RegTiempo struct {
	Estacion    Estacion
	Equipo      Equipo
	HoraLlegada time.Time `json:"HoraLlegada"`
}
