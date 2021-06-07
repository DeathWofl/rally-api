package preguntas

import (
	"github.com/spinales/rally-api/internal/platform/logger"
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

type Preguntas struct {
	logHandler logger.Logger
	store      store
}

func (prs *Preguntas) PreguntaPorID(id uint) (*Pregunta, error) {
	p, err := prs.PreguntaPorID(id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (prs *Preguntas) TodasPreguntas() (*[]Pregunta, error) {
	p, err := prs.store.All()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (prs *Preguntas) CrearPregunta(p *Pregunta) (*Pregunta, error) {
	p, err := prs.store.Create(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (prs *Preguntas) ActualizarPregunta(id uint, p *Pregunta) (*Pregunta, error) {
	p, err := prs.store.Update(id, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (prs *Preguntas) Delete(id uint) error {
	err := prs.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewService(l logger.Logger, db *gorm.DB) (*Preguntas, error) {
	pstore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Preguntas{
		logHandler: l,
		store:      pstore,
	}, nil
}
