package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DeathWofl/rally-api/pkg/models"
	"github.com/labstack/echo"
)

//PostRegRespuesta agregar registro de respuesta
func (s *Service) PostRegRespuesta(c echo.Context) error {
	// recibiendo el usuario a crear
	reg := models.RegResp{}
	if err := c.Bind(&reg); err != nil {
		log.Fatalln("procesando registro de respuesta: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	creg, err := s.RegRespService.CreateRegResp(&reg)
	if err != nil {
		log.Fatalln("creando registro de respuesta: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, creg)
}

//PostAllRegRespuesta agregar registro de respuesta
func (s *Service) PostAllRegRespuesta(c echo.Context) error {
	rr := []models.RegResp{}
	if err := c.Bind(&rr); err != nil {
		log.Fatalln("procesando registro de respuestas: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	for _, r := range rr {
		_, err := s.RegRespService.CreateRegResp(&r)
		if err != nil {
			log.Fatalln("creando registro de respuesta: %w", err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusOK)
}

//GetAllRegRespuesta retorna todas los registros de respuestas
func (s *Service) GetAllRegRespuesta(c echo.Context) error {
	// trayendo registros de respuestas
	regs, err := s.RegRespService.RegResps()
	if err != nil {
		log.Fatalln("buscando registro de respuesta: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, regs)
}

//GetRegRespuesta busca regrespuesta por su ID
func (s *Service) GetRegRespuesta(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// trayendo registros de respuestas
	reg, err := s.RegRespService.RegResp(uint(id))
	if err != nil {
		log.Fatalln("buscando registro de respuesta: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, reg)
}
