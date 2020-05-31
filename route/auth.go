package route

import (
	"net/http"
	"time"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JWTCustomClaim struct {
	iat int64 `json:"iat"`
	ID  uint  `json:"equipoID"` // ID asignado sin importar si es maestro o estudiante, tienen diferentes endpoints para loguearse
	jwt.StandardClaims
}

//LoginEstu ruta para la authentificacion de los estudiantes
func LoginEstu(c echo.Context) error {

	DB := db.DBManager()

	matriculaE1 := c.FormValue("MatriculaE1")
	matriculaE2 := c.FormValue("MatriculaE2")
	matriculaE3 := c.FormValue("MatriculaE3")

	result := models.Equipo{}
	DB.
		Where(models.Equipo{MatriculaE1: matriculaE1}).
		Or(models.Equipo{MatriculaE2: matriculaE1}).
		Or(models.Equipo{MatriculaE3: matriculaE1}).
		Where(models.Equipo{MatriculaE1: matriculaE2}).
		Or(models.Equipo{MatriculaE2: matriculaE2}).
		Or(models.Equipo{MatriculaE3: matriculaE2}).
		Where(models.Equipo{MatriculaE1: matriculaE3}).
		Or(models.Equipo{MatriculaE2: matriculaE3}).
		Or(models.Equipo{MatriculaE3: matriculaE3}).
		First(&result)

	if result.CodigoGrupo == "" {
		return c.String(http.StatusNotAcceptable, "Las matriculas no pertenecen a un mismo grupo")
	}

	if result.LoggedIn == true {
		return c.String(http.StatusNotAcceptable, "Ya esta loggeado.")
	}

	// claims
	claims := &JWTCustomClaim{
		time.Now().Unix(),
		result.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("itesarally"))
	if err != nil {
		return err
	}
	result.Token = t
	result.LoggedIn = true

	return c.JSON(http.StatusOK, result)
}

//LoginUser logearse maestros
func LoginUser(c echo.Context) error {

	DB := db.DBManager()

	username := c.FormValue("Username")
	password := c.FormValue("Password")

	result := models.Usuario{}
	DB.Where(&models.Usuario{Username: username, Password: password}).First(&result)

	if result.Nombre == "" {
		return echo.ErrNotFound
	}

	claims := &JWTCustomClaim{
		time.Now().Unix(),
		result.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("itesarally"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
