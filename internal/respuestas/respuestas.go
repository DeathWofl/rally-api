package respuestas

import (
	"gorm.io/gorm"
)

type Respuesta struct {
	gorm.Model
	Resp       string `json:"Resp" gorm:"type:varchar(200);not null"`
	Valor      int
	PreguntaID uint
}
