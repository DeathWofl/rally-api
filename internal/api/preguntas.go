package api

import (
	"fmt"

	"github.com/spinales/rally-api/internal/preguntas"
)

func (a *API) TodasPreguntas() (*[]preguntas.Pregunta, error) {
	pts, err := a.preguntas.TodasPreguntas()
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando preguntas: %q", err))
		return nil, err
	}

	return pts, nil
}

func (a *API) PreguntaPorID(id uint) (*preguntas.Pregunta, error) {
	pt, err := a.preguntas.PreguntaPorID(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Buscando pregunta por ID: %q", err))
		return nil, err
	}

	return pt, nil
}

func (a *API) CrearPregunta(pt *preguntas.Pregunta) (*preguntas.Pregunta, error) {
	pt, err := a.preguntas.CrearPregunta(pt)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Creando pregunta: %q", err))
		return nil, err
	}

	return pt, nil
}

func (a *API) ActualizarPregunta(pt *preguntas.Pregunta, id uint) (*preguntas.Pregunta, error) {
	pt, err := a.preguntas.ActualizarPregunta(id, pt)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Actualizando pregunta: %q", err))
		return nil, err
	}

	return pt, nil
}

func (a *API) EliminarPregunta(id uint) error {
	err := a.preguntas.EliminarPregunta(id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Elimando pregunta: %q", err))
		return err
	}

	return nil
}
