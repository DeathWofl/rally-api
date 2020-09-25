package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/migration"
	"github.com/DeathWofl/rally-api/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	p := e.Group("/api")

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Rutas sin authorizacion
	p.POST("/loginalumnos", route.LoginEstu)  // probado
	p.POST("/loginmaestros", route.LoginUser) // Probado

	app := p.Group("/app")

	config := middleware.JWTConfig{
		Claims:     &route.JWTCustomClaim{},
		SigningKey: ([]byte("itesarally")),
	}
	app.Use(middleware.JWTWithConfig(config))

	//// Authorization
	//Equipos
	app.GET("/equipo", route.GetAllEquipos)       // Probado
	app.GET("/equipo/:id", route.GetEquipo)       // Probado
	app.POST("/equipo", route.PostEquipo)         // Probado
	app.PUT("/equipo/:id", route.PutEquipo)       // Probado
	app.DELETE("/equipo/:id", route.DeleteEquipo) // Probado

	//Estaciones
	app.POST("/estacion", route.PostEstacion)         // Probado
	app.GET("/estacion", route.GetAllEstacion)        // Probado
	app.GET("/estacion/:id", route.GetEstacion)       // Probado
	app.PUT("/estacion/:id", route.PutEstacion)       // Probado
	app.DELETE("/estacion/:id", route.DeleteEstacion) // Probado

	//Respuesta
	app.GET("/respuestas/:id", route.GetRespuesta)      // Probado
	app.GET("/respuestas", route.GetAllRespuestas)      // Probado
	app.POST("/respuestas", route.PostRespuesta)        // Probado
	app.PUT("/respuestas/:id", route.PutRespuesta)      //Probado
	app.DELETE("/respuestas/:id", route.DeleteEstacion) // Probado

	// Preguntas
	app.GET("/preguntas", route.GetAllQuestion)        // Probado
	app.GET("/preguntas/:id", route.GetQuestion)       // Probado
	app.POST("/preguntas", route.PostQuestion)         // Probado
	app.PUT("/preguntas/:id", route.PutQuestion)       // Probado
	app.DELETE("/preguntas/:id", route.DeleteQuestion) // Probado

	// Usuarios
	app.GET("/usuarios", route.GetAllUsers)       // Probado
	app.GET("/usuarios/:id", route.GetUser)       // Probado
	app.POST("/usuarios", route.PostUser)         // Probado
	app.PUT("/usuarios/:id", route.PutUser)       // Probado
	app.DELETE("/usuarios/:id", route.DeleteUser) // Probado

	// Registro de respuestas
	app.POST("/regrespuesta", route.PostRegRespuesta) // Probado
	app.POST("/regrespuesta/all", route.PostAllRegRespuesta)
	app.GET("/regrespuesta", route.GetAllRegRespuesta)  // Probado
	app.GET("/regrespuesta/:id", route.GetRegRespuesta) // Probado

	//Registro de tiempos
	app.POST("/regtiempo", route.PostRegTiempo)   // Probado
	app.GET("/regtiempo", route.GetAllRegsTiempo) // Probado
	app.GET("/regtiempo/:id", route.GetRegTiempo) // Probado

	// Ganadores
	app.GET("/ganadores", route.GetGanadores)

	e.Logger.Fatal(e.Start(":1323"))
}
