package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DeathWofl/rally-api/pkg/models"
	"github.com/labstack/echo"
)

//GetAllUsers Select all user
func (s *Service) GetAllUsers(c echo.Context) error {
	// trayendo usuarios
	user, err := s.UsuarioService.Usuarios()
	if err != nil {
		log.Fatalln("buscando usuarios: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

//GetUser Selec one user
func (s *Service) GetUser(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// trayendo usuario
	user, err := s.UsuarioService.Usuario(uint(id))
	if err != nil {
		log.Fatalln("buscando usuario: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

//PostUser Create new user
func (s *Service) PostUser(c echo.Context) error {
	// recibiendo el usuario a crear
	user := models.Usuario{}
	if err := c.Bind(&user); err != nil {
		log.Fatalln("procesando usuario: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	cuser, err := s.UsuarioService.CreateUsuario(&user)
	if err != nil {
		log.Fatalln("creando usuario: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, cuser)
}

//DeleteUser user usuario
func (s *Service) DeleteUser(c echo.Context) error {

	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// eliminado equipo
	err = s.UsuarioService.DeleteUsuario(uint(id))
	if err != nil {
		log.Fatalln("eliminando usuario: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

//PutUser actualizar un usuario
func (s *Service) PutUser(c echo.Context) error {

	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// procesando usuario
	var user models.Usuario
	if err := c.Bind(user); err != nil {
		log.Fatalln("procesando usuario: %w", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// actualizando
	updateuser, err := s.UsuarioService.UpdateUsuario(uint(id), &user)
	if err != nil {
		log.Fatalln("actualizando usuario: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, updateuser)
}
