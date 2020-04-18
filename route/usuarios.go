package route

import (
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

// Select all user
func GetAllUsers(c echo.Context) error {
	DB := db.DBManager()
	user := []models.Usuario{}
	DB.Find(&user)
	return c.JSON(http.StatusOK, user)
}

// Selec one user
func GetUser(c echo.Context) error {
	DB := db.DBManager()
	id := c.Param("id")
	user := models.Usuario{}
	DB.First(&user, id)
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "El usuario no existe.")
	}
	return c.JSON(http.StatusOK, user)
}

// Create new user
func PostUser(c echo.Context) error {
	DB := db.DBManager()
	user := models.Usuario{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Create(&user)
	return c.JSON(http.StatusOK, user)
}

//DeleteUser user usuario
func DeleteUser(c echo.Context) error {
	DB := db.DBManager()
	user := models.Usuario{}
	id := c.Param("id")
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Delete(&user, id)
	return c.String(http.StatusOK, "Usuario eliminado")
}

func PutUser(c echo.Context) error {
	DB := db.DBManager()
	id := c.Param("id")
	user := models.Usuario{}
	DB.Find(&user, id)
	putuser := new(models.Usuario)
	if err := c.Bind(putuser); err != nil {
		panic(err)
	}
	DB.Model(&user).Updates(&putuser)
	return c.JSON(http.StatusOK, user)
}
