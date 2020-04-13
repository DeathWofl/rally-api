package models

//Usuario tabla de usuarios para la aute.. de los maestros.
type Usuario struct {
	Nombre   string `json:"Nombre" gorm:"type:varchar(200);not null"`
	Username string `json:"Username" gorm:"type:varchar(100);not null"`
	Password string `json:"Password" gorm:"type:varchar(100); not null;"`
	Estacion Estacion
}
