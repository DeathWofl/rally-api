package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spinales/rally-api/pkg/models"
)

// JWTCustomClaim datos encritados en el JWT
type JWTCustomClaim struct {
	iat int64 `json:"iat"`
	ID  uint  `json:"equipoID"` // ID asignado sin importar si es maestro o estudiante, tienen diferentes endpoints para loguearse
	jwt.StandardClaims
}

// LoginEstu ruta para la authentificacion de los estudiantes
func (s *Service) LoginEstu(c echo.Context) error {
	result := models.Equipo{}
	err := c.Bind(&result)
	if err != nil {
		log.Fatalln("login estu: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	sear := models.Equipo{}

	s.EquipoService.DB.Where(&models.Equipo{MatriculaE1: result.MatriculaE1}).
		Or(&models.Equipo{MatriculaE1: result.MatriculaE2}).
		Or(&models.Equipo{MatriculaE1: result.MatriculaE3}).
		First(&sear)

	if sear.ContraGrupo != result.ContraGrupo {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"response": "Usuario invalido, confirme la informacion enviada",
		})
	}

	// claims
	claims := &JWTCustomClaim{
		time.Now().Unix(),
		result.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("itesarally"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":       t,
		"CodigoGrupo": sear.CodigoGrupo,
	})
}

// LoginUser logearse maestros
func (s *Service) LoginUser(c echo.Context) error {
	// datos enviados, parseando json
	parse := models.Usuario{}
	err := c.Bind(&parse)
	if err != nil {
		log.Fatalln("login user: %w", err)
	}

	sear := models.Usuario{}
	s.UsuarioService.DB.Where(&models.Usuario{Username: parse.Username, Password: parse.Password}).First(&sear)

	if sear.Username != parse.Username || sear.Password != parse.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"response": "Usuario invalido",
		})
	}

	claims := &JWTCustomClaim{
		time.Now().Unix(),
		parse.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
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
