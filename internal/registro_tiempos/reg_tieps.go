package regtieps

import (
	"time"

	"github.com/spinales/rally-api/internal/platform/logger"
	"gorm.io/gorm"
)

// RegTiempo representa el horaria que un equipo llego a una estacion
// RegTiempo = Registro de Tiempo
type RegTiempo struct {
	gorm.Model
	EstacionID  uint      `json:"EstacionID"`
	EquipoID    uint      `json:"EquipoID"`
	CodigoGrupo string    `json:"CodigoGrupo" gorm:"not null;type:varchar(100);"`
	HoraLlegada time.Time `json:"HoraLlegada"`
}

type RegTiempos struct {
	logHandler logger.Logger
	store      store
}

func (rts *RegTiempos) RegTiempoPorID(id uint) (*RegTiempo, error) {
	rt, err := rts.store.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (rts *RegTiempos) TodosRegTiempos() (*[]RegTiempo, error) {
	rt, err := rts.store.All()
	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (rts *RegTiempos) CrearRegTiempo(r *RegTiempo) (*RegTiempo, error) {
	rt, err := rts.store.Create(r)
	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (rts *RegTiempos) ActualizarRegTiempo(id uint, r *RegTiempo) (*RegTiempo, error) {
	rt, err := rts.store.Update(id, r)
	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (rts *RegTiempos) EliminarRegTiempo(id uint) error {
	err := rts.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewService(l logger.Logger, db *gorm.DB) (*RegTiempos, error) {
	regstore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &RegTiempos{
		logHandler: l,
		store:      regstore,
	}, nil
}
