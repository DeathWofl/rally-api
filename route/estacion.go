package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//GetAllEstacion Retorna todas las Estacion
func GetAllEstacion(c echo.Context) error {
	DB := db.DBManager()
	estacion := []models.Estacion{}
	DB.Preload("pregunta").Preload("reg_tiempos").Find(&estacion)
	return c.JSON(http.StatusOK, estacion)
}

//GetEstacion Retorna Estacion por ID
func GetEstacion(c echo.Context) error {
	DB := db.DBManager()
	id := c.Param("id")
	estacion := models.Estacion{}
	DB.First(&estacion, id)
	if estacion.ID == 0 {
		return c.String(http.StatusOK, "La estacion no existe")
	}
	// DB.Preload("pregunta").Preload("reg_tiempos").First(&estacion, id)
	return c.JSON(http.StatusOK, estacion)
}

//PostEstacion Registra un Estacion
func PostEstacion(c echo.Context) error {
	DB := db.DBManager()
	estacion := models.Estacion{}
	err := c.Bind(&estacion)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Create(&estacion)
	return c.JSON(http.StatusOK, estacion)
}

//PutEstacion Actualiza un Estacion
func PutEstacion(c echo.Context) error {
	DB := db.DBManager()
	estacion := models.Estacion{}
	id := c.Param("id")
	DB.Find(&estacion, id)
	putestacion := new(models.Estacion)
	if err := c.Bind(putestacion); err != nil {
		panic(err)
	}
	DB.Model(&estacion).Updates(&putestacion)
	return c.JSON(http.StatusOK, estacion)
}

//DeleteEstacion Elimina un Estacion
func DeleteEstacion(c echo.Context) error {
	DB := db.DBManager()
	estacion := models.Estacion{}
	id := c.Param("id")
	DB.Delete(&estacion, id)
	if err := c.Bind(&estacion); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Estacion eliminada")
	}
	return c.NoContent(http.StatusOK)
}
