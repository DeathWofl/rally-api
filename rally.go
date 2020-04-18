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

	//Registro de respuestas
	e.POST("/regrespuesta", route.PostRegRespuesta)
	e.GET("/regrespuesta", route.RegRespuestas)
	e.GET("/regrespuesta/:ID", route.RegRespuesta)
	e.GET("/regrespuesta", route.BuscarRegRespuesta)

	//Registro de tiempos
	e.POST("/regtiempo", route.PostRegTiempo)
	e.GET("/regtiempo", route.RegsTiempo)
	e.GET("/regtiempo/:ID", route.RegTiempo)
	e.GET("/regtiempo", route.BuscarRegRespuesta)

	e.GET("/preguntas", route.GetAllQuestion)
	// e.GET("/usuarios", route.GetAllUsers)
	// e.POST("/usuarios", route.PostUser)
	// e.DELETE("/usuarios", route.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
