package api

import (
	"fmt"

	regresp "github.com/spinales/rally-api/internal/registro_respuestas"
)

func (a *API) TodosRegResp() (*[]regresp.RegResp, error) {
	rs, err := a.regresp.TodosRegResp()
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando registro de respuesta: %q", err))
		return nil, err
	}

	return rs, nil
}

func (a *API) RegRespPorID(id uint) (*regresp.RegResp, error) {
	r, err := a.regresp.RegRespPorID(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando registro de respuesta por ID: %q", err))
		return nil, err
	}

	return r, nil
}

func (a *API) CrearRegResp(r *regresp.RegResp) (*regresp.RegResp, error) {
	r, err := a.regresp.CrearRegResp(r)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Creando registro de respuesta: %q", err))
		return nil, err
	}

	return r, nil
}

func (a *API) ActualizarRegResp(r *regresp.RegResp, id uint) (*regresp.RegResp, error) {
	r, err := a.regresp.ActualizarRegResp(id, r)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Actualizando registro de respuesta: %q", err))
		return nil, err
	}

	return r, nil
}

func (a *API) EliminarRegResp(id uint) error {
	err := a.regresp.EliminarRegResp(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Elimando registro de respuesta: %q", err))
		return err
	}

	return nil
}
