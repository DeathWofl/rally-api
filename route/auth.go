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

	result := models.Equipo{}
	c.Bind(&result)

	sear := models.Equipo{}
	DB.Where(&models.Equipo{MatriculaE1: result.MatriculaE1, MatriculaE2: result.MatriculaE2, MatriculaE3: result.MatriculaE3}).First(&sear)

	if sear.CodigoGrupo == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"response": "Usuario invalido, confirme la informacion enviada",
		})
	}

	if sear.LoggedIn == true {
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
	sear.Token = t
	sear.LoggedIn = true

	return c.JSON(http.StatusOK, sear)
}

//LoginUser logearse maestros
func LoginUser(c echo.Context) error {

	DB := db.DBManager()

	result := models.Usuario{}
	c.Bind(&result)

	sear := models.Usuario{}
	DB.Where(&models.Usuario{Username: result.Username, Password: result.Password}).First(&sear)

	if sear.Username != result.Username || sear.Password != result.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"response": "Usuario invalido",
		})
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
