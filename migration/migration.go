package migration

import (
	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
)

//Migrate migracion a la base de datos
func Migrate() {

	DB := db.DBManager()

	DB.LogMode(true)

	//Equipos
	DB.DropTableIfExists(&models.Equipo{})
	DB.CreateTable(&models.Equipo{})

	// Estaciones
	DB.DropTableIfExists(&models.Estacion{})
	DB.CreateTable(&models.Estacion{})

	// Respuestas
	DB.DropTableIfExists(&models.Respuesta{})
	DB.CreateTable(&models.Respuesta{})

	// Preguntas
	DB.DropTableIfExists(&models.Pregunta{})
	DB.CreateTable(&models.Pregunta{})

	// Usuarios
	DB.DropTableIfExists(&models.Usuario{})
	DB.CreateTable(&models.Usuario{})
	user := models.Usuario{Nombre: "Admin", Username: "admin", Password: "admin"}
	DB.Create(&user)

	// Registros de respuestas
	DB.DropTableIfExists(&models.RegResp{})
	DB.CreateTable(&models.RegResp{})

	// Registros de Tiempos
	DB.DropTableIfExists(&models.RegTiempo{})
	DB.CreateTable(&models.RegTiempo{})

}
