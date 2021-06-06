package estaciones

import (
	"github.com/jinzhu/gorm"
	"github.com/spinales/rally-api/internal/preguntas"
	regtieps "github.com/spinales/rally-api/internal/registro_tiempos"
)

// Estacion tabla de estaciones del rally
type Estacion struct {
	gorm.Model
	Nombre      string `json:"nombre" gorm:"not null;type:varchar(20)"`
	Description string `json:"Descripcion" gorm:"type:varchar(250)"`
	Preguntas   []preguntas.Pregunta
	RegTiempos  []regtieps.RegTiempo
}
