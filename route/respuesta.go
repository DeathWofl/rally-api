package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//GetRespuesta retorna especificamente una respuesta por su ID
func GetRespuesta(c echo.Context) error {
	DB := db.DBManager()
	id := c.Param("id")
	Resp := models.Respuesta{}
	DB.First(&Resp, id)
	DB.First(&Resp.Preg, Resp.PreguntaID)
	return c.JSON(http.StatusOK, Resp)
}

//GetAllRespuestas retorna todas las respuestas existentes
func GetAllRespuestas(c echo.Context) error {
	DB := db.DBManager()
	resp := []models.Respuesta{}
	DB.Find(&resp)
	return c.JSON(http.StatusOK, resp)
}

//PutRespuesta actualizar una respuesta
func PutRespuesta(c echo.Context) error {
	DB := db.DBManager()
	resp := models.Respuesta{}
	id := c.Param("id")
	DB.Find(&resp, id)
	putequipo := new(models.Respuesta)
	if err := c.Bind(putequipo); err != nil {
		panic(err)
	}
	DB.Model(&resp).Updates(&putequipo)
	return c.JSON(http.StatusOK, resp)
}

//PostRespuesta Agregar respuesta
func PostRespuesta(c echo.Context) error {
	DB := db.DBManager()

	Resp := models.Respuesta{}
	err := c.Bind(&Resp)
	if err != nil {
		panic(err)
	}

	DB.Create(&Resp)
	return c.JSON(http.StatusOK, Resp)
}

//DeleteRespuesta Elimina un Estacion
func DeleteRespuesta(c echo.Context) error {
	DB := db.DBManager()
	respuesta := models.Respuesta{}
	id := c.Param("id")
	DB.Delete(&respuesta, id)
	if err := c.Bind(&respuesta); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Estacion eliminada")
	}
	return c.NoContent(http.StatusOK)
}
