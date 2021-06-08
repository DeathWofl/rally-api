package api

import (
	"fmt"

	"github.com/spinales/rally-api/internal/estaciones"
)

func (a *API) TodosEstaciones() (*[]estaciones.Estacion, error) {
	ets, err := a.estaciones.TodasEstaciones()
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando estaciones: %q", err))
		return nil, err
	}

	return ets, nil
}

func (a *API) EstacionPorID(id uint) (*estaciones.Estacion, error) {
	et, err := a.estaciones.EstacionPorID(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando estacion por ID: %q", err))
		return nil, err
	}

	return et, nil
}

func (a *API) CrearEstacion(et *estaciones.Estacion) (*estaciones.Estacion, error) {
	et, err := a.estaciones.CrearEstacion(et)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Creando estacion: %q", err))
		return nil, err
	}

	return et, nil
}

func (a *API) ActualizarEstacion(et *estaciones.Estacion, id uint) (*estaciones.Estacion, error) {
	et, err := a.estaciones.ActualizarEstacion(id, et)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Actualizando estacion: %q", err))
		return nil, err
	}

	return et, nil
}

func (a *API) EliminarEstacion(id uint) error {
	err := a.estaciones.EliminarEstacion(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Elimando estacion: %q", err))
		return err
	}

	return nil
}
