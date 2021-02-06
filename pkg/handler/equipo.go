package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DeathWofl/rally-api/pkg/models"
	"github.com/labstack/echo"
)

//GetAllEquipos Retorna todos los equipos
func (s *Service) GetAllEquipos(c echo.Context) error {
	// trayendo equipos
	equipo, err := s.EquipoService.Equipos()
	if err != nil {
		log.Fatalln("buscando equipos: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, equipo)
}

//GetEquipo Retorna equipo por ID
func (s *Service) GetEquipo(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// trayendo equipo
	equipo, err := s.EquipoService.Equipo(uint(id))
	if err != nil {
		log.Fatalln("buscando equipo: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, equipo)
}

//PostEquipo Registra un equipo
func (s *Service) PostEquipo(c echo.Context) error {
	// recibiendo el equipo a crear
	equipo := models.Equipo{}
	if err := c.Bind(&equipo); err != nil {
		log.Fatalln("procesando equipo: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	cequipo, err := s.EquipoService.CreateEquipo(&equipo)
	if err != nil {
		log.Fatalln("creando equipo: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, cequipo)
}

//PutEquipo Actualiza un equipo
func (s *Service) PutEquipo(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// procesando equipo
	var equipo models.Equipo
	if err := c.Bind(equipo); err != nil {
		log.Fatalln("procesando equipo: %w", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// actualizando
	updateequipo, err := s.EquipoService.UpdateEquipo(uint(id), &equipo)
	if err != nil {
		log.Fatalln("actualizando equipo: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, updateequipo)
}

//DeleteEquipo Elimina un equipo
func (s *Service) DeleteEquipo(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// eliminado equipo
	err = s.EquipoService.DeleteEquipo(uint(id))
	if err != nil {
		log.Fatalln("eliminando equipo: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
