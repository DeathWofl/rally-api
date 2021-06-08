package api

import (
	"fmt"

	regtieps "github.com/spinales/rally-api/internal/registro_tiempos"
)

func (a *API) TodosRegTiempos() (*[]regtieps.RegTiempo, error) {
	rs, err := a.regtieps.TodosRegTiempos()
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando registros de tiempo: %q", err))
		return nil, err
	}

	return rs, nil
}

func (a *API) RegTiempoPorID(id uint) (*regtieps.RegTiempo, error) {
	r, err := a.regtieps.RegTiempoPorID(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando registro de tiempo por ID: %q", err))
		return nil, err
	}

	return r, nil
}

func (a *API) CrearRegTiempo(r *regtieps.RegTiempo) (*regtieps.RegTiempo, error) {
	r, err := a.regtieps.CrearRegTiempo(r)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Creando registro de tiempo: %q", err))
		return nil, err
	}

	return r, nil
}

func (a *API) ActualizarRegTiempo(r *regtieps.RegTiempo, id uint) (*regtieps.RegTiempo, error) {
	r, err := a.regtieps.ActualizarRegTiempo(id, r)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Actualizando registro de tiempo: %q", err))
		return nil, err
	}

	return r, nil
}

func (a *API) EliminarRegTiempo(id uint) error {
	err := a.regtieps.EliminarRegTiempo(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Elimando registro de tiempo: %q", err))
		return err
	}

	return nil
}
