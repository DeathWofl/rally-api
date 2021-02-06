package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// GetPuntuaciones retorna todos los equipos por orden de puntaje y tiempo
func (s *Service) GetPuntuaciones(c echo.Context) error {
	puntuaciones, err := s.PuntuacionService.Puntuaciones()
	if err != nil {
		log.Fatalln("buscando puntuaciones: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, puntuaciones)
}
