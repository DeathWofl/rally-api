package models

//RegResp Registros de las respuestas correctas
type RegResp struct {
	ID           int `json:"IDRegResp" gorm:"not null;primary key;auto_increment"`
	Pregunta     Pregunta
	Calificacion int `json:"Calificacion"`
	Equipo       Equipo
}
