package models

import "github.com/jinzhu/gorm"

//Estacion tabla de estaciones del rally
type Estacion struct {
	gorm.Model
	Nombre string `json:"nombre" gorm:"not null;type:varchar(20)"`
}
