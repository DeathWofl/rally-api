package route

import (
	"net/http"
	"time"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//PostRegTiempo agregar un registro de tiempo
func PostRegTiempo(c echo.Context) error {
	DB := db.DBManager()

	Restiem := models.RegTiempo{}
	err := c.Bind(&Restiem)
	if err != nil {
		panic(err)
	}

	Restiem.HoraLlegada = time.Now()
	DB.Create(&Restiem)
	return c.JSON(http.StatusOK, Restiem)
}

//GetAllRegsTiempo todos los registros de tiempo
func GetAllRegsTiempo(c echo.Context) error {
	DB := db.DBManager()

	Regs := []models.RegTiempo{}
	DB.Find(&Regs)

	return c.JSON(http.StatusOK, Regs)
}

//GetRegTiempo buscar un registro de tiempo por su ID
func GetRegTiempo(c echo.Context) error {
	DB := db.DBManager()
	ID := c.Param("id")

	Reg := models.RegTiempo{}
	DB.Find(&Reg, ID)

	return c.JSON(http.StatusOK, Reg)
}
