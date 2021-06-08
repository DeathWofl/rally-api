package api

import (
	"fmt"

	"github.com/spinales/rally-api/internal/usuarios"
)

func (a *API) TodosUsuarios() (*[]usuarios.Usuario, error) {
	us, err := a.usuarios.TodosUsuarios()
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando usuarios: %q", err))
		return nil, err
	}

	return us, nil
}

func (a *API) UsuariosPorID(id uint) (*usuarios.Usuario, error) {
	u, err := a.usuarios.UsuarioPorID(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando usuario por ID: %q", err))
		return nil, err
	}

	return u, nil
}

func (a *API) CrearUsuario(u *usuarios.Usuario) (*usuarios.Usuario, error) {
	u, err := a.usuarios.CrearUsuario(u)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Creando usuario: %q", err))
		return nil, err
	}

	return u, nil
}

func (a *API) ActualizarUsuario(u *usuarios.Usuario, id uint) (*usuarios.Usuario, error) {
	u, err := a.usuarios.ActualizarUsuario(id, u)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Actualizando usuario: %q", err))
		return nil, err
	}

	return u, nil
}

func (a *API) EliminarUsuario(id uint) error {
	err := a.usuarios.EliminarUsuario(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Eliminando usuario: %q", err))
		return err
	}

	return nil
}
