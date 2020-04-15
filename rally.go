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
	e.GET("/preguntas", route.GetAllQuestion)
	e.GET("/usuarios", route.GetAllUsers)
	e.POST("/usuarios", route.PostUser)
	e.DELETE("/usuarios", route.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
