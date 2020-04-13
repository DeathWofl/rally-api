package migration

import (
	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
)

//Migrate migracion a la base de datos
func Migrate() {

	db := db.DBManager

	db.LogMode(true)

	//Equipos
	db.DropTableIfExists(&models.Equipo{})
	db.CreateTable(&models.Equipo{})

	//Estaciones
	db.DropTableIfExists(&models.Estacion{})
	db.CreateTable(&models.Estacion{})

	//Respuestas
	db.DropTableIfExists(&models.Respuesta{})
	db.CreateTable(&models.Respuesta{})

	//Preguntas
	db.DropTableIfExists(&models.Pregunta{})
	db.CreateTable(&models.Pregunta{})

	//Registros de respuestas
	db.DropTableIfExists(&models.RegResp{})
	db.CreateTable(&models.RegResp{})

	//Registros de Tiempos
	db.DropTableIfExists(&models.RegTiempo{})
	db.CreateTable(&models.RegTiempo{})

	// Usuarios
	db.DropTableIfExists(&models.Usuario{})
	db.CreateTable(&models.Usuario{})
	user := models.Usuario{Name: "Admin", Username: "admin", Password: "admin"}
	db.Create(&user)
}
