package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DeathWofl/rally-api/pkg/models"
	"github.com/labstack/echo"
)

//GetRespuesta retorna especificamente una respuesta por su ID
func (s *Service) GetRespuesta(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// trayendo respuesta
	resp, err := s.RespuestaService.Respuesta(uint(id))
	if err != nil {
		log.Fatalln("buscando respuesta: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// en caso de que el usuario no exitas
	if resp.ID == 0 {
		return c.JSON(http.StatusBadRequest, "La respuesta no existe.")
	}
	return c.JSON(http.StatusOK, resp)
}

//GetAllRespuestas retorna todas las respuestas existentes
func (s *Service) GetAllRespuestas(c echo.Context) error {
	// trayendo respuestas
	resp, err := s.RespuestaService.Respuestas()
	if err != nil {
		log.Fatalln("buscando respuestas: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

//PutRespuesta actualizar una respuesta
func (s *Service) PutRespuesta(c echo.Context) error {
	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// procesando respuesta
	var resp models.Respuesta
	if err := c.Bind(resp); err != nil {
		log.Fatalln("procesando respuesta: %w", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// actualizando
	updateresp, err := s.RespuestaService.UpdateRespuesta(uint(id), &resp)
	if err != nil {
		log.Fatalln("actualizando respuesta: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, updateresp)
}

//PostRespuesta Agregar respuesta
func (s *Service) PostRespuesta(c echo.Context) error {
	// recibiendo el usuario a crear
	var resp models.Respuesta
	if err := c.Bind(&resp); err != nil {
		log.Fatalln("procesando respuesta: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	cresp, err := s.RespuestaService.CreateRespuesta(&resp)
	if err != nil {
		log.Fatalln("creando respuesta: %w", err)
	}

	return c.JSON(http.StatusOK, cresp)
}

//DeleteRespuesta Elimina un Estacion
func (s *Service) DeleteRespuesta(c echo.Context) error {

	// trayendo el parametro de la URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("convirtiendo ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// eliminado respuesta
	err = s.RespuestaService.DeleteRespuesta(uint(id))
	if err != nil {
		log.Fatalln("eliminando respuesta: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
