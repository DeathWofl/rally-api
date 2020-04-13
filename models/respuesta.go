package models

//Respuesta tabla de respuestas
type Respuesta struct {
	Resp  string `json:"Resp" gorm:"type:varchar(200);not null"`
	valor int    `json:"Valor" gorm:"not null;"`
}
