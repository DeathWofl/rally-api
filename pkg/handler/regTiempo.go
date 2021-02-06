package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/DeathWofl/rally-api/pkg/models"
	"github.com/labstack/echo"
)

//PostRegTiempo agregar un registro de tiempo
func (s *Service) PostRegTiempo(c echo.Context) error {
	// recibiendo el usuario a crear
	reg := models.RegTiempo{}
	if err := c.Bind(&reg); err != nil {
		log.Fatalln("procesando registro de tiempo: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	reg.HoraLlegada = time.Now()
	creg, err := s.RegTiempoService.CreateRegTiempo(&reg)
	if err != nil {
		log.Fatalln("creando registro de tiempo: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, creg)
}

//GetAllRegsTiempo todos los registros de tiempo
func (s *Service) GetAllRegsTiempo(c echo.Context) error {
	// trayendo registro de tiempo
	reg, err := s.RegTiempoService.RegTiempos()
	if err != nil {
		log.Fatalln("buscando registro de tiempos: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, reg)
}

//GetRegTiempo buscar un registro de tiempo por su ID
func (s *Service) GetRegTiempo(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// trayendo registro de tiempo
	reg, err := s.RegTiempoService.RegTiempo(uint(id))
	if err != nil {
		log.Fatalln("buscando registro de tiempo: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, reg)
}
