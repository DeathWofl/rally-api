package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

func GetAllQuestion(c echo.Context) error {
	DB := db.Init()
	question := models.Pregunta{}
	DB.Find(&question)
	return c.JSON(http.StatusOK, question)
}

func GetQuestion(c echo.Context) error {
	DB := db.Init()
	question := models.Pregunta{}
	id := c.Param("id")
	DB.Find(&question, id)
	return c.JSON(http.StatusOK, question)
}

func PostQuestion(c echo.Context) error {
	DB := db.Init()
	question := models.Pregunta{}
	err := c.Bind(&question)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Create(&question)
	return c.JSON(http.StatusOK, question)
}
