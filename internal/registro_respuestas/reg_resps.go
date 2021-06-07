package regresp

import (
	"github.com/spinales/rally-api/internal/platform/logger"
	"gorm.io/gorm"
)

// RegResp representa cada vez que un equipo responde una
// pregunta con su puntuacion
// RegResp = Registro de Respuestas
type RegResp struct {
	gorm.Model
	PreguntaID   uint `json:"PreguntaID"`
	Calificacion uint `json:"Calificacion"`
	EquipoID     uint `json:"EquipoID"`
}

type RegResps struct {
	logHandler logger.Logger
	store      store
}

func (rrs *RegResps) RegRespPorID(id uint) (*RegResp, error) {
	rr, err := rrs.store.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return rr, nil
}

func (rrs *RegResps) TodosRegResp() (*[]RegResp, error) {
	rr, err := rrs.store.All()
	if err != nil {
		return nil, err
	}

	return rr, nil
}

func (rrs *RegResps) CrearRegResp(r *RegResp) (*RegResp, error) {
	rr, err := rrs.store.Create(r)
	if err != nil {
		return nil, err
	}

	return rr, nil
}

func (rrs *RegResps) ActualizarRegResp(id uint, r *RegResp) (*RegResp, error) {
	rr, err := rrs.store.Update(id, r)
	if err != nil {
		return nil, err
	}

	return rr, nil
}

func (rrs *RegResps) EliminarRegResp(id uint) error {
	err := rrs.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewService(l logger.Logger, db *gorm.DB) (*RegResps, error) {
	regstore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &RegResps{
		logHandler: l,
		store:      regstore,
	}, nil
}
