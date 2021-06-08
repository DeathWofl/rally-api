package api

import (
	"time"

	"github.com/spinales/rally-api/internal/equipos"
	"github.com/spinales/rally-api/internal/estaciones"
	"github.com/spinales/rally-api/internal/platform/logger"
	"github.com/spinales/rally-api/internal/preguntas"
	"github.com/spinales/rally-api/internal/puntuaciones"
	regresp "github.com/spinales/rally-api/internal/registro_respuestas"
	regtieps "github.com/spinales/rally-api/internal/registro_tiempos"
	"github.com/spinales/rally-api/internal/respuestas"
	"github.com/spinales/rally-api/internal/usuarios"
)

type API struct {
	logger       logger.Logger
	equipos      *equipos.Equipos
	estaciones   *estaciones.Estaciones
	preguntas    *preguntas.Preguntas
	puntuaciones *puntuaciones.Puntuaciones
	regresp      *regresp.RegResps
	regtieps     *regtieps.RegTiempos
	respuestas   *respuestas.Respuestas
	usuarios     *usuarios.Usuarios
}

func (a *API) Health() (map[string]interface{}, error) {
	return map[string]interface{}{
		"env":        "testing",
		"version":    "v0.1.0",
		"commit":     "<git commit hash>",
		"status":     "all systems up and running",
		"startedAt":  time.Now().String(),
		"releasedOn": time.Now().String(),
	}, nil
}

func NewService(l logger.Logger, eqs *equipos.Equipos, etns *estaciones.Estaciones, prs *preguntas.Preguntas, pns *puntuaciones.Puntuaciones, rresp *regresp.RegResps, rtieps *regtieps.RegTiempos, rs *respuestas.Respuestas, uss *usuarios.Usuarios) (*API, error) {
	return &API{
		logger:       l,
		equipos:      eqs,
		estaciones:   etns,
		preguntas:    prs,
		puntuaciones: pns,
		regresp:      rresp,
		regtieps:     rtieps,
		respuestas:   rs,
		usuarios:     uss,
	}, nil
}
