package preguntas

import (
	regresp "github.com/spinales/rally-api/internal/registro_respuestas"
	"github.com/spinales/rally-api/internal/respuestas"
	"gorm.io/gorm"
)

// Pregunta representa cada pregunta del proceso.
type Pregunta struct {
	gorm.Model
	Preg       string `json:"Preg" gorm:"type:varchar(250);not null"`
	Respuestas []respuestas.Respuesta
	RegResps   []regresp.RegResp
	EstacionID uint
}
