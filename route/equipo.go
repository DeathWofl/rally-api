package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//Retorna todos los equipos
func GetAllEquipos(c echo.Context) error {
	DB := db.DBManager()
	equipo := []models.Equipo{}
	DB.Find(&equipo)
	return c.JSON(http.StatusOK, equipo)
}

//Retorna equipo por ID
func GetEquiposID(c echo.Context) error{
	DB := db.DBManager()	
	equipo := models.Equipo{}
	id := c.Param("id")
	DB.Find(&equipo, id)
	return c.JSON(http.StatusOK, equipo)
}

//Registra un equipo
func PostEquipo(c echo.Context) error{
	DB := db.DBManager()
	equipo := models.Equipo{}
	err := c.Bind(&equipo)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Create(&equipo)
	return c.JSON(http.StatusOK, equipo)
}

//Actualiza un equipo
func PutEquipo(c echo.Context) error{
	DB := db.DBManager()
	equipo := models.Equipo{}
	id := c.Param("id")
	DB.Find(&equipo, id)
	putequipo := new(models.Equipo)
	if err := c.Bind(putequipo); err !=nil{
		panic(err)
	}
	DB.Model(&equipo).Updates(&putequipo)
	return c.JSON(http.StatusOK, equipo)
}

//Elimina un equipo
func DeleteEquipo(c echo.Context) error{
	DB := db.DBManager()
	equipo := models.Equipo{}
	id := c.Param("id")
	DB.Delete(&equipo, id)
	if err := c.Bind(&equipo); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Equipo eliminado")
	}
	return c.String(http.StatusOK, "Equipo eliminado")
}





