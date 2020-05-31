package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//PostRegRespuesta agregar registro de respuesta
func PostRegRespuesta(c echo.Context) error {
	DB := db.DBManager()

	Reg := models.RegResp{}
	c.Bind(&Reg)

	DB.Create(&Reg)

	return c.JSON(http.StatusOK, Reg)
}

//GetAllRegRespuesta retorna todas los registros de respuestas
func GetAllRegRespuesta(c echo.Context) error {
	DB := db.DBManager()

	Regs := []models.RegResp{}
	DB.Find(&Regs)

	return c.JSON(http.StatusOK, Regs)
}

//GetRegRespuesta busca regrespuesta por su ID
func GetRegRespuesta(c echo.Context) error {
	DB := db.DBManager()
	ID := c.Param("id")

	Reg := models.RegResp{}

	DB.Find(&Reg, ID)

	return c.JSON(http.StatusOK, Reg)
}
