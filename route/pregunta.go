package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

func getAllQuestion(c echo.Context) error {
	DB := db.DBManager()
	question := models.Pregunta{}
	DB.Find(&question)
	return c.JSON(http.StatusOK, question)
}

func getQuestion(c echo.Context) error {
	DB := db.DBManager()
	question := models.Pregunta{}
	id := c.Param("id")
	DB.Find(&question, id)
	return c.JSON(http.StatusOK, question)
}

func postQuestion(c echo.Context) error {
	DB := db.DBManager()
	question := models.Pregunta{}
	err := c.Bind(&question)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Create(&question)
	return c.JSON(http.StatusOK, question)
}
