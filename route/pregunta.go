package route

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//GetAllQuestion Todas las preguntas
func GetAllQuestion(c echo.Context) error {
	DB := db.DBManager()
	question := []models.Pregunta{}
	DB.Find(&question)
	return c.JSON(http.StatusOK, question)
}

//GetQuestion Select one question
func GetQuestion(c echo.Context) error {

	//DB
	DB := db.DBManager()

	//busqueda
	question := models.Pregunta{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	ID := uint(id)
	fmt.Println(ID)
	DB.Find(&question, ID)
	DB.Where(&models.Respuesta{PreguntaID: ID}).Find(&question.Respuestas)
	DB.Where(&models.RegResp{PreguntaID: ID}).Find(&question.RegResps)

	//retornando
	if question.ID == 0 {
		return c.JSON(http.StatusNotFound, "Pregunta no existente")
	}
	return c.JSON(http.StatusOK, question)
}

//PostQuestion borrar pregunta
func PostQuestion(c echo.Context) error {
	DB := db.DBManager()
	question := models.Pregunta{}
	err := c.Bind(&question)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Create(&question)
	return c.JSON(http.StatusOK, question)
}

//DeleteQuestion Delete question
func DeleteQuestion(c echo.Context) error {
	DB := db.DBManager()
	question := models.Pregunta{}
	id := c.Param("id")
	DB.Delete(&question, id)
	if id == "" {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusOK)
}

//PutQuestion Actualizar pregunta
func PutQuestion(c echo.Context) error {
	DB := db.DBManager()
	id := c.Param("id")
	question := models.Pregunta{}
	DB.Find(&question, id)
	putquestion := new(models.Pregunta)
	if err := c.Bind(putquestion); err != nil {
		panic(err)
	}
	DB.Model(&question).Updates(&putquestion)
	return c.JSON(http.StatusOK, question)
}
