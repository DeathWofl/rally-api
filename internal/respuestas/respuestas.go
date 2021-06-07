package respuestas

import (
	"github.com/spinales/rally-api/internal/platform/logger"
	"gorm.io/gorm"
)

type Respuesta struct {
	gorm.Model
	Resp       string `json:"Resp" gorm:"type:varchar(200);not null"`
	Valor      int
	PreguntaID uint
}

type Respuestas struct {
	logHandler logger.Logger
	store      store
}

func (rs *Respuestas) RespuestaPorID(id uint) (*Respuesta, error) {
	rp, err := rs.store.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return rp, nil
}

func (rs *Respuestas) TodasRespuestas() (*[]Respuesta, error) {
	rps, err := rs.store.All()
	if err != nil {
		return nil, err
	}

	return rps, nil
}

func (rs *Respuestas) CrearRespuesta(r *Respuesta) (*Respuesta, error) {
	rp, err := rs.store.Create(r)
	if err != nil {
		return nil, err
	}

	return rp, nil
}

func (rs *Respuestas) ActualizarRespuesta(id uint, r *Respuesta) (*Respuesta, error) {
	rp, err := rs.store.Update(id, r)
	if err != nil {
		return nil, err
	}

	return rp, nil
}

func (rs *Respuestas) EliminarRespuesta(id uint) error {
	err := rs.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewService(l logger.Logger, db *gorm.DB) (*Respuestas, error) {
	respstore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Respuestas{
		logHandler: l,
		store:      respstore,
	}, nil
}
