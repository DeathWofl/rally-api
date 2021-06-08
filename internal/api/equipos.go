package api

import (
	"fmt"

	"github.com/spinales/rally-api/internal/equipos"
)

func (a *API) TodosEquipos() (*[]equipos.Equipo, error) {
	eqs, err := a.equipos.TodosEquipos()
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando equipos: %q", err))
		return nil, err
	}

	return eqs, nil
}

func (a *API) EquipoPorID(id uint) (*equipos.Equipo, error) {
	eq, err := a.equipos.EquipoPorID(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando equipo por ID: %q", err))
		return nil, err
	}

	return eq, nil
}

func (a *API) CrearEquipo(eq *equipos.Equipo) (*equipos.Equipo, error) {
	eq, err := a.equipos.CrearEquipo(eq)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Creando equipo: %q", err))
		return nil, err
	}

	return eq, nil
}

func (a *API) ActualizarEquipo(eq *equipos.Equipo, id uint) (*equipos.Equipo, error) {
	eq, err := a.equipos.ActualizarEquipo(id, eq)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Actualizando equipo: %q", err))
		return nil, err
	}

	return eq, nil
}

func (a *API) EliminarEquipo(id uint) error {
	err := a.equipos.EliminarEquipo(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Eliminando equipo: %q", err))
		return err
	}

	return nil
}
