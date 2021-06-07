package estaciones

import (
	"github.com/spinales/rally-api/internal/platform/logger"
	"github.com/spinales/rally-api/internal/preguntas"
	regtieps "github.com/spinales/rally-api/internal/registro_tiempos"
	"gorm.io/gorm"
)

// Estacion tabla de estaciones del rally
type Estacion struct {
	gorm.Model
	Nombre      string `json:"nombre" gorm:"not null;type:varchar(20)"`
	Description string `json:"Descripcion" gorm:"type:varchar(250)"`
	Preguntas   []preguntas.Pregunta
	RegTiempos  []regtieps.RegTiempo
}

type Estaciones struct {
	logHandler logger.Logger
	store      store
}

func (es *Estaciones) EstacionPorID(id uint) (*Estacion, error) {
	e, err := es.store.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (es *Estaciones) TodasEstaciones() (*[]Estacion, error) {
	e, err := es.store.All()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (es *Estaciones) CrearEstacion(e *Estacion) (*Estacion, error) {
	e, err := es.store.Create(e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (es *Estaciones) ActualizarEstacion(id uint, e *Estacion) (*Estacion, error) {
	e, err := es.store.Update(id, e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (es *Estaciones) EliminarEstacion(id uint) error {
	err := es.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewService(l logger.Logger, db *gorm.DB) (*Estaciones, error) {
	estore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Estaciones{
		logHandler: l,
		store:      estore,
	}, nil
}
