package api

import (
	"fmt"

	"github.com/spinales/rally-api/internal/respuestas"
)

func (a *API) TodosRespuestas() (*[]respuestas.Respuesta, error) {
	rps, err := a.respuestas.TodasRespuestas()
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando respuestas: %q", err))
		return nil, err
	}

	return rps, nil
}

func (a *API) RespuestaPorID(id uint) (*respuestas.Respuesta, error) {
	rp, err := a.respuestas.RespuestaPorID(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando respuesta por ID: %q", err))
		return nil, err
	}

	return rp, nil
}

func (a *API) CrearRespuesta(rp *respuestas.Respuesta) (*respuestas.Respuesta, error) {
	rp, err := a.respuestas.CrearRespuesta(rp)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Creando respuesta: %q", err))
		return nil, err
	}

	return rp, nil
}

func (a *API) ActualizarRespuesta(rp *respuestas.Respuesta, id uint) (*respuestas.Respuesta, error) {
	rp, err := a.respuestas.ActualizarRespuesta(id, rp)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Actualizando respuesta: %q", err))
		return nil, err
	}

	return rp, nil
}

func (a *API) EliminarRespuesta(id uint) error {
	err := a.respuestas.EliminarRespuesta(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Eliminando respuesta: %q", err))
		return err
	}

	return nil
}
