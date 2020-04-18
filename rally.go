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

	e.GET("/respuesta", route.GetRespuesta)

	// Users
	e.GET("/usuarios", route.GetAllUsers)
	e.GET("/usuarios/:id", route.GetUser)       // Funciona
	e.POST("/usuarios", route.PostUser)         // Funciona
	e.DELETE("/usuarios/:id", route.DeleteUser) // Funciona

	// Questions
	e.GET("/preguntas", route.GetAllQuestion)
	e.GET("/preguntas/:id", route.GetQuestion)       // Funciona
	e.POST("/preguntas", route.PostQuestion)         // Funciona
	e.DELETE("/preguntas/:id", route.DeleteQuestion) // Funciona
	e.Logger.Fatal(e.Start(":1323"))
}
