package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//Retorna todas las Estacion
func GetAllEstacion(c echo.Context) error {
	DB := db.DBManager()
	estacion := []models.Estacion{}
	DB.Find(&estacion)
	return c.JSON(http.StatusOK, estacion)
}

//Retorna Estacion por ID
func GetEstacionID(c echo.Context) error{
	DB := db.DBManager()	
	estacion := models.Estacion{}
	id := c.Param("id")
	DB.Find(&estacion, id)
	return c.JSON(http.StatusOK, estacion)
}

//Registra un Estacion
func PostEstacion(c echo.Context) error{
	DB := db.DBManager()
	estacion := models.Estacion{}
	err := c.Bind(&estacion)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Create(&estacion)
	return c.JSON(http.StatusOK, estacion)
}

//Actualiza un Estacion
func PutEstacion(c echo.Context) error{
	DB := db.DBManager()
	estacion := models.Estacion{}
	id := c.Param("id")
	DB.Find(&estacion, id)
	putestacion := new(models.Estacion)
	if err := c.Bind(putestacion); err !=nil{
		panic(err)
	}
	DB.Model(&estacion).Updates(&putestacion)
	return c.JSON(http.StatusOK, estacion)
}

//Elimina un Estacion
func DeleteEstacion(c echo.Context) error{
	DB := db.DBManager()
	estacion := models.Estacion{}
	id := c.Param("id")
	DB.Delete(&estacion, id)
	if err := c.Bind(&estacion); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Estacion eliminada")
	}
	return c.String(http.StatusOK, "Estacion eliminada")
}





