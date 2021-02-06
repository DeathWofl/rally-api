package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DeathWofl/rally-api/pkg/models"
	"github.com/labstack/echo"
)

//GetAllEstacion Retorna todas las Estacion
func (s *Service) GetAllEstacion(c echo.Context) error {
	// trayendo estaciones
	estaciones, err := s.EstacionService.Estaciones()
	if err != nil {
		log.Fatalln("buscando estaciones: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, estaciones)
}

//GetEstacion Retorna Estacion por ID
func (s *Service) GetEstacion(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// trayendo estacion
	estacion, err := s.EstacionService.Estacion(uint(id))
	if err != nil {
		log.Fatalln("buscando estacion: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, estacion)
}

//PostEstacion Registra un Estacion
func (s *Service) PostEstacion(c echo.Context) error {
	// recibiendo el usuario a crear
	estacion := models.Estacion{}
	if err := c.Bind(&estacion); err != nil {
		log.Fatalln("procesando estacion: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	cestacion, err := s.EstacionService.CreateEstacion(&estacion)
	if err != nil {
		log.Fatalln("creando estacion: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, cestacion)
}

//PutEstacion Actualiza un Estacion
func (s *Service) PutEstacion(c echo.Context) error {

	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// procesando estacion
	var estacion models.Estacion
	if err := c.Bind(&estacion); err != nil {
		log.Fatalln("procesando estacion: %w", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// actualizando
	updateestacion, err := s.EstacionService.UpdateEstacion(uint(id), &estacion)
	if err != nil {
		log.Fatalln("actualizando estacion: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, updateestacion)
}

//DeleteEstacion Elimina un Estacion
func (s *Service) DeleteEstacion(c echo.Context) error {

	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// eliminado estacion
	err = s.EstacionService.DeleteEstacion(uint(id))
	if err != nil {
		log.Fatalln("eliminando estacion: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
