package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

func GetAllQuestion(c echo.Context) error {
	DB := db.DBManager()
	question := []models.Pregunta{}
	DB.Find(&question)
	return c.JSON(http.StatusOK, question)
}

// Select one question
func GetQuestion(c echo.Context) error {
	DB := db.DBManager()
	question := models.Pregunta{}
	id := c.Param("id")
	DB.First(&question, id)
	if question.ID == 0 {
		return c.JSON(http.StatusNotFound, "Pregunta no existente")
	}
	return c.JSON(http.StatusOK, question)
}

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

// Delete question
func DeleteQuestion(c echo.Context) error {
	DB := db.DBManager()
	question := models.Pregunta{}
	id := c.Param("id")
	DB.Delete(&question, id)
	if question.ID == 0 {
		return c.JSON(http.StatusNotFound, "No se ha podido eliminar la pregunta.")
	}
	return c.String(http.StatusOK, "Pregunta eliminada")
}

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
