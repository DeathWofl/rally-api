package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DeathWofl/rally-api/pkg/models"
	"github.com/labstack/echo"
)

//GetAllQuestion Todas las preguntas
func (s *Service) GetAllQuestion(c echo.Context) error {
	// trayendo preguntas
	pregs, err := s.PreguntaService.Preguntas()
	if err != nil {
		log.Fatalln("buscando preguntas: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, pregs)
}

//GetQuestion Select one question
func (s *Service) GetQuestion(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// trayendo pregunta
	preg, err := s.PreguntaService.Pregunta(uint(id))
	if err != nil {
		log.Fatalln("buscando pregunta: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, preg)
}

//PostQuestion agregar pregunta
func (s *Service) PostQuestion(c echo.Context) error {
	// recibiendo el pregunta a crear
	preg := models.Pregunta{}
	if err := c.Bind(&preg); err != nil {
		log.Fatalln("procesando pregunta: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	cuser, err := s.PreguntaService.CreatePregunta(&preg)
	if err != nil {
		log.Fatalln("creando pregunta: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, cuser)
}

//DeleteQuestion Delete question
func (s *Service) DeleteQuestion(c echo.Context) error {

	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// eliminado pregunta
	err = s.PreguntaService.DeletePregunta(uint(id))
	if err != nil {
		log.Fatalln("eliminando pregunta: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

//PutQuestion Actualizar pregunta
func (s *Service) PutQuestion(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// procesando pregunta
	var preg models.Pregunta
	if err := c.Bind(preg); err != nil {
		log.Fatalln("procesando pregunta: %w", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// actualizando pregunta
	updatepreg, err := s.PreguntaService.UpdatePregunta(uint(id), &preg)
	if err != nil {
		log.Fatalln("actualizando pregunta: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, updatepreg)
}
