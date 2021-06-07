package puntuaciones

import (
	"github.com/spinales/rally-api/internal/platform/logger"
	"gorm.io/gorm"
)

// Puntuacion representa cada puntuacion de cada equipo
// antiguamente representado como ganador
type Puntuacion struct {
	gorm.Model
	Transcurso string
	Puntaje    float32
}

type Puntuaciones struct {
	logHandler logger.Logger
	store      store
}

func (ps *Puntuaciones) TodasPuntuaciones() (*[]Puntuacion, error) {
	p, err := ps.store.All()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func NewService(l logger.Logger, db *gorm.DB) (*Puntuaciones, error) {
	pstore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Puntuaciones{
		logHandler: l,
		store:      pstore,
	}, nil
}
