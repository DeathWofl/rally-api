package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	intern "github.com/spinales/internal/configs"
	"github.com/spinales/rally-api/pkg/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsc := intern.ConnectionString()
	DB, err := gorm.Open(mysql.Open(dsc), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion de la base de datos.")
	flag.Parse()

	if migrate == "yes" {
		fmt.Println("Comenzo la Migracion...")
		storage.Migrate(DB)
		fmt.Println("Termino la Migracion...")
	}

	service := &handler.Service{
		EquipoService:     &storage.EquipoService{DB},
		EstacionService:   &storage.EstacionService{DB: DB},
		PreguntaService:   &storage.PreguntaService{DB},
		PuntuacionService: &storage.PuntuacionService{DB},
		RegRespService:    &storage.RegRespService{DB},
		RegTiempoService:  &storage.RegTiempoService{DB},
		RespuestaService:  &storage.RespuestaService{DB},
		UsuarioService:    &storage.UsuarioService{DB},
	}

	e := echo.New()
	p := e.Group("/api")

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Rutas sin authorizacion
	p.POST("/loginalumnos", service.LoginEstu)  // probado
	p.POST("/loginmaestros", service.LoginUser) // Probado

	app := p.Group("/app")

	config := middleware.JWTConfig{
		Claims:     &handler.JWTCustomClaim{},
		SigningKey: ([]byte("itesarally")),
	}
	app.Use(middleware.JWTWithConfig(config))

	//// Authorization
	//Equipos
	app.GET("/equipo", service.GetAllEquipos)       // Probado
	app.GET("/equipo/:id", service.GetEquipo)       // Probado
	app.POST("/equipo", service.PostEquipo)         // Probado
	app.PUT("/equipo/:id", service.PutEquipo)       // Probado
	app.DELETE("/equipo/:id", service.DeleteEquipo) // Probado

	// Estaciones
	app.POST("/estacion", service.PostEstacion)         // Probado
	app.GET("/estacion", service.GetAllEstacion)        // Probado
	app.GET("/estacion/:id", service.GetEstacion)       // Probado
	app.PUT("/estacion/:id", service.PutEstacion)       // Probado
	app.DELETE("/estacion/:id", service.DeleteEstacion) // Probado

	// Respuesta
	app.GET("/respuestas/:id", service.GetRespuesta)      // Probado
	app.GET("/respuestas", service.GetAllRespuestas)      // Probado
	app.POST("/respuestas", service.PostRespuesta)        // Probado
	app.PUT("/respuestas/:id", service.PutRespuesta)      // Probado
	app.DELETE("/respuestas/:id", service.DeleteEstacion) // Probado

	// Preguntas
	app.GET("/preguntas", service.GetAllQuestion)        // Probado
	app.GET("/preguntas/:id", service.GetQuestion)       // Probado
	app.POST("/preguntas", service.PostQuestion)         // Probado
	app.PUT("/preguntas/:id", service.PutQuestion)       // Probado
	app.DELETE("/preguntas/:id", service.DeleteQuestion) // Probado

	// Usuarios
	app.GET("/usuarios", service.GetAllUsers)       // Probado
	app.GET("/usuarios/:id", service.GetUser)       // Probado
	app.POST("/usuarios", service.PostUser)         // Probado
	app.PUT("/usuarios/:id", service.PutUser)       // Probado
	app.DELETE("/usuarios/:id", service.DeleteUser) // Probado

	// Registro de respuestas
	app.POST("/regrespuesta", service.PostRegRespuesta) // Probado
	app.POST("/regrespuesta/all", service.PostAllRegRespuesta)
	app.GET("/regrespuesta", service.GetAllRegRespuesta)  // Probado
	app.GET("/regrespuesta/:id", service.GetRegRespuesta) // Probado

	// Registro de tiempos
	app.POST("/regtiempo", service.PostRegTiempo)   // Probado
	app.GET("/regtiempo", service.GetAllRegsTiempo) // Probado
	app.GET("/regtiempo/:id", service.GetRegTiempo) // Probado

	// Ganadores
	app.GET("/ganadores", service.GetPuntuaciones)

	e.Logger.Fatal(e.Start(":1323"))
}
