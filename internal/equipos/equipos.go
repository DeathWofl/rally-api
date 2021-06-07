package equipos

import (
	"github.com/spinales/rally-api/internal/platform/logger"
	regresp "github.com/spinales/rally-api/internal/registro_respuestas"
	regtieps "github.com/spinales/rally-api/internal/registro_tiempos"
	"gorm.io/gorm"
)

// Equipo representa cada equipo del rally
type Equipo struct {
	gorm.Model
	MatriculaE1 string `json:"MatriculaE1" gorm:"not null;type:varchar(15);"`
	MatriculaE2 string `json:"MatriculaE2" gorm:"not null;type:varchar(15);"`
	MatriculaE3 string `json:"MatriculaE3" gorm:"not null;type:varchar(15);"`
	CodigoGrupo string `json:"CodigoGrupo" gorm:"not null;type:varchar(100);" sql:"index"`
	ContraGrupo string `json:"ContraGrupo" gorm:"not null;type:varchar(100);"`
	RegResps    []regresp.RegResp
	RegTiempos  []regtieps.RegTiempo
}

type Equipos struct {
	logHandler logger.Logger
	store      store
}

func (es *Equipos) EquipoPorID(id uint) (*Equipo, error) {
	e, err := es.store.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (es *Equipos) TodosEquipos() (*[]Equipo, error) {
	e, err := es.store.All()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (es *Equipos) CrearEquipo(e *Equipo) (*Equipo, error) {
	e, err := es.store.Create(e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (es *Equipos) ActualizarEquipo(id uint, e *Equipo) (*Equipo, error) {
	e, err := es.store.Update(id, e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (es *Equipos) EliminarEquipo(id uint) error {
	err := es.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewService(l logger.Logger, db *gorm.DB) (*Equipos, error) {
	ustore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Equipos{
		logHandler: l,
		store:      ustore,
	}, nil
}
