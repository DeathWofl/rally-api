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
	c.Bind(Reg)

	DB.Create(&Reg)

	return c.JSON(http.StatusOK, Reg)
}

//RegRespuestas retorna todas los registros de respuestas
func RegRespuestas(c echo.Context) error {
	DB := db.DBManager()

	Regs := []models.RegResp{}
	DB.Find(&Regs)

	return c.JSON(http.StatusOK, Regs)
}

//RegRespuesta busca regrespuesta por su ID
func RegRespuesta(c echo.Context) error {
	DB := db.DBManager()
	ID := c.Param("ID")

	Reg := models.RegResp{}

	DB.Find(&Reg, ID)

	return c.JSON(http.StatusOK, Reg)
}

//BuscarRegRespuesta retorna todos los registros de respuestas de un equipo
func BuscarRegRespuesta(c echo.Context) error {
	DB := db.DBManager()

	equipo := models.Equipo{}
	c.Bind(equipo)

	Registros := []models.RegResp{}
	err := DB.Model(&Registros).Related(&equipo)
	if err != nil {
		panic(err)
	}

	if Registros == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, Registros)
}
