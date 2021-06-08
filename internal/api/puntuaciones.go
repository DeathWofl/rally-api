package api

import (
	"fmt"

	"github.com/spinales/rally-api/internal/puntuaciones"
)

func (a *API) TodasPuntuaciones() (*[]puntuaciones.Puntuacion, error) {
	pts, err := a.puntuaciones.TodasPuntuaciones()
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando puntuaciones: %q", err))
		return nil, err
	}

	return pts, nil
}
