package models

import "github.com/jinzhu/gorm"

//Equipo tabla de equipos del rally
type Equipo struct {
	gorm.Model
	MatriculaE1 string `json:"MatriculaE1" gorm:"not null;type:varchar(15);"`
	MatriculaE2 string `json:"MatriculaE2" gorm:"not null;type:varchar(15);"`
	MatriculaE3 string `json:"MatriculaE3" gorm:"not null;type:varchar(15);"`
	CodigoGrupo string `json:"CodigoGrupo" gorm:"not null;type:varchar(100);" sql:"index"`
	ContraGrupo string `json:"ContraGrupo" gorm:"not null;type:varchar(100);"`
	RegResps    []RegResp
	RegTiempos  []RegTiempo
}
