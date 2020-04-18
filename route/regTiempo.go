package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//PostRegTiempo agregar un registro de tiempo
func PostRegTiempo(c echo.Context) error {
	DB := db.DBManager()

	reg := models.RegTiempo{}
	c.Bind(reg)

	DB.Create(&reg)

	return c.JSON(http.StatusOK, reg)
}

//RegsTiempo todos los registros de tiempo
func RegsTiempo(c echo.Context) error {
	DB := db.DBManager()

	Regs := []models.RegTiempo{}
	DB.Find(&Regs)

	return c.JSON(http.StatusOK, Regs)
}

//RegTiempo buscar un registro de tiempo por su ID
func RegTiempo(c echo.Context) error {
	DB := db.DBManager()
	ID := c.Param("ID")

	Reg := models.RegTiempo{}
	DB.Find(&Reg, ID)

	return c.JSON(http.StatusOK, Reg)
}

//BuscarRegTiempo buscar los registros de tiempo pertenecientes a un grupo
func BuscarRegTiempo(c echo.Context) error {
	DB := db.DBManager()

	equipo := models.Equipo{}
	c.Bind(equipo)

	Regs := []models.RegTiempo{}
	err := DB.Model(Regs).Related(equipo)
	if err != nil {
		panic(err)
	}

	if Regs == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, Regs)
}
