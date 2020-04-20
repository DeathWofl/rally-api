package main

import (
	"flag"
	"fmt"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/migration"
	"github.com/DeathWofl/rally-api/route"
	"github.com/labstack/echo"
)

func main() {
	db.Init()
	defer db.DB.Close()

	var migrate string

	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion de la base de datos.")

	flag.Parse()

	if migrate == "yes" {
		fmt.Println("Comenzo la Migracion...")
		migration.Migrate()
		fmt.Println("Termino la Migracion...")
	}

	e := echo.New()

	//Respuesta
	e.GET("/respuesta/:ID", route.GetRespuesta)
	e.GET("/respuesta", route.Respuestas)
	e.POST("/respuesta", route.PostRespuesta)
	e.PUT("/respuesta/:ID", route.PutRespuesta)
	e.GET("/respuesta", route.BuscarRespuestas)

	// //Registro de respuestas
	// e.POST("/regrespuesta", route.PostRegRespuesta)
	// e.GET("/regrespuesta", route.RegRespuestas)
	// e.GET("/regrespuesta/:ID", route.RegRespuesta)
	// e.GET("/regrespuesta", route.BuscarRegRespuesta)

	//Registro de tiempos
	//

	e.GET("/preguntas", route.GetAllQuestion)
	// e.GET("/usuarios", route.GetAllUsers)
	// e.POST("/usuarios", route.PostUser)
	// e.DELETE("/usuarios", route.DeleteUser)

	// Users
	e.GET("/usuarios", route.GetAllUsers)       // Funciona
	e.GET("/usuarios/:id", route.GetUser)       // Funciona
	e.POST("/usuarios", route.PostUser)         // Funciona
	e.PUT("/usuarios/:id", route.PutUser)       // Funciona
	e.DELETE("/usuarios/:id", route.DeleteUser) // Funciona

	// Questions
	e.GET("/preguntas", route.GetAllQuestion)        // Funciona
	e.GET("/preguntas/:id", route.GetQuestion)       // Funciona
	e.POST("/preguntas", route.PostQuestion)         // Funciona
	e.PUT("/preguntas/:id", route.PutQuestion)       // Funciona
	e.DELETE("/preguntas/:id", route.DeleteQuestion) // Funciona
	
	
	//Equipos
	e.GET("/equipo", route.GetAllEquipos)
	e.GET("/equipo/:id", route.GetEquiposID)
	e.POST("/equipo", route.PostEquipo)
	e.PUT("/equipo/:id", route.PutEquipo)
	e.DELETE("/equipo/:id", route.DeleteEquipo)

	//Estaciones
	e.GET("/estacion", route.GetAllEstacion)
	e.GET("/estacion/:id", route.GetEstacionID)
	e.POST("/estacion", route.PostEstacion)
	e.PUT("/estacion/:id", route.PutEstacion)
	e.DELETE("/estacion/:id", route.DeleteEstacion)

	e.Logger.Fatal(e.Start(":1323"))
}
