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
	id := c.Param("ID")
	Resp := models.Respuesta{}
	DB.Find(&Resp, id)
	return c.JSON(http.StatusOK, Resp)
}

//Respuestas retorna todas las respuestas existentes
func Respuestas(c echo.Context) error {
	DB := db.DBManager()
	Resp := []models.Respuesta{}
	DB.Find(&Resp)
	return c.JSON(http.StatusOK, Resp)
}

//PutRespuesta actualizar una respuesta
func PutRespuesta(c echo.Context) error {

	//abro conexion
	DB := db.DBManager()

	// tomo el ID de parametro
	ID := c.Param("ID")

	// Respuesta antes de actualizar
	Resp := models.Respuesta{}
	DB.Find(&Resp, ID)

	//Datos de los campos a actualizar
	PutResp := models.Respuesta{}
	if err := c.Bind(PutUser); err != nil {
		panic(err)
	}

	//Actualizando Datos
	DB.Model(&Resp).Updates(PutResp)

	return c.JSON(http.StatusOK, PutUser)
}

//PostRespuesta Agregar respuesta
func PostRespuesta(c echo.Context) error {
	DB := db.DBManager()

	Resp := models.Respuesta{}
	err := c.Bind(Resp)
	if err != nil {
		panic(err)
	}

	DB.Create(&Resp)
	return c.JSON(http.StatusOK, Resp)
}

//BuscarRespuestas busca respuestas que correspondan a una pregunta
func BuscarRespuestas(c echo.Context) error {
	DB := db.DBManager()

	Pregunt := models.Pregunta{}
	c.Bind(Pregunt)

	Respues := []models.Respuesta{}

	err := DB.Model(&Respues).Related(&Pregunt)
	if err != nil {
		panic(err)
	}

	if Respues == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, Respues)
}
