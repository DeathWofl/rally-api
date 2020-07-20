package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

//GetAllEquipos Retorna todos los equipos
func GetAllEquipos(c echo.Context) error {
	DB := db.DBManager()
	equipo := []models.Equipo{}
	DB.Preload("reg_resps").Preload("reg_tiempos").Find(&equipo)
	DB.Find(&equipo)
	return c.JSON(http.StatusOK, equipo)
}

//GetEquipo Retorna equipo por ID
func GetEquipo(c echo.Context) error {
	DB := db.DBManager()
	id := c.Param("id")
	equipo := models.Equipo{}
	DB.First(&equipo, id)
	if equipo.ID == 0 {
		return c.String(http.StatusOK, "El Equipo no existe")
	}
	// DB.Preload("reg_resps").Preload("reg_tiempos").Find(&equipo)
	return c.JSON(http.StatusOK, equipo)
}

//PostEquipo Registra un equipo
func PostEquipo(c echo.Context) error {

	DB := db.DBManager()

	// Captura los datos
	equipo := models.Equipo{}
	err := c.Bind(&equipo)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	// Busqueda para evitar que se repita una matricula
	result := models.Equipo{}

	// Confirma que la matricula de ese estudiante no este registrada con anterioridad en la BD
	// Especificamente del estudiante de la matricula 1
	DB.Where(models.Equipo{MatriculaE1: equipo.MatriculaE1}).
		Or(models.Equipo{MatriculaE2: equipo.MatriculaE1}).
		Or(models.Equipo{MatriculaE3: equipo.MatriculaE1}).
		Find(&result)
	if result.MatriculaE1 != "" {
		return c.String(http.StatusConflict, "La Matricula del Estudiante 1 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 2
	DB.Where(models.Equipo{MatriculaE1: equipo.MatriculaE2}).
		Or(models.Equipo{MatriculaE2: equipo.MatriculaE2}).
		Or(models.Equipo{MatriculaE3: equipo.MatriculaE2}).
		Find(&result)
	if result.MatriculaE2 != "" {
		return c.String(http.StatusConflict, "La Matricula del Estudiante 2 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 3
	DB.Where(models.Equipo{MatriculaE1: equipo.MatriculaE3}).
		Or(models.Equipo{MatriculaE2: equipo.MatriculaE3}).
		Or(models.Equipo{MatriculaE3: equipo.MatriculaE3}).
		Find(&result)
	if result.MatriculaE1 != "" {
		return c.String(http.StatusConflict, "La Matricula del Estudiante 3 ya esta registrada")
	}

	DB.Create(&equipo)
	return c.JSON(http.StatusOK, equipo)
}

//PutEquipo Actualiza un equipo
func PutEquipo(c echo.Context) error {
	DB := db.DBManager()
	id := c.Param("id")
	equipo := models.Equipo{}
	DB.Find(&equipo, id)
	putequipo := new(models.Equipo)
	if err := c.Bind(putequipo); err != nil {
		panic(err)
	}
	validequipo := models.Equipo{}
	DB.Where(models.Equipo{MatriculaE1: putequipo.MatriculaE1}).
		Or(models.Equipo{MatriculaE2: putequipo.MatriculaE1}).
		Or(models.Equipo{MatriculaE3: putequipo.MatriculaE1}).
		Not(id).
		Find(&validequipo)
	if validequipo.MatriculaE1 != "" {
		return c.String(http.StatusConflict, "La Matricula del Estudiante 1 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 2
	DB.Where(models.Equipo{MatriculaE1: putequipo.MatriculaE2}).
		Or(models.Equipo{MatriculaE2: putequipo.MatriculaE2}).
		Or(models.Equipo{MatriculaE3: putequipo.MatriculaE2}).
		Not(id).
		Find(&validequipo)
	if validequipo.MatriculaE2 != "" {
		return c.String(http.StatusConflict, "La Matricula del Estudiante 2 ya esta registrada")
	}

	// Especificamente del estudiante de la matricula 3
	DB.Where(models.Equipo{MatriculaE1: putequipo.MatriculaE3}).
		Or(models.Equipo{MatriculaE2: putequipo.MatriculaE3}).
		Or(models.Equipo{MatriculaE3: putequipo.MatriculaE3}).
		Not(id).
		Find(&validequipo)
	if validequipo.MatriculaE1 != "" {
		return c.String(http.StatusConflict, "La Matricula del Estudiante 3 ya esta registrada")
	}
	DB.Model(&equipo).Updates(&putequipo)
	return c.JSON(http.StatusOK, equipo)
}

//DeleteEquipo Elimina un equipo
func DeleteEquipo(c echo.Context) error {
	DB := db.DBManager()
	equipo := models.Equipo{}
	id := c.Param("id")
	DB.Delete(&equipo, id)
	if err := c.Bind(&equipo); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Equipo eliminado")
	}
	return c.NoContent(http.StatusOK)
}
