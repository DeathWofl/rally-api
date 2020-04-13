package models

//Pregunta tabla de preguntas
type Pregunta struct {
	Preg       string `json:"Preg" gorm:"type:varchar(250);not null"`
	Respuestas []Respuesta
	Estacion   Estacion
}
