package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//GetRespuesta retorna especificamente una respuesta
func GetRespuesta(c echo.Context) {

	DB := db.DBManager()

	id := c.Param("id")
	Resp := models.Respuesta
	DB.Find(&Resp, id)
	return c.JSON(http.StatusOK, Resp)
}
