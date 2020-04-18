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
	user := models.Usuario{}
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Find(&user)
	return c.JSON(http.StatusOK, user)
}

// Selec one user
func GetUser(c echo.Context) error {
	DB := db.DBManager()
	id := c.Param("id")
	user := models.Usuario{}
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
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
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Create(&user)
	return c.JSON(http.StatusOK, user)
}

// Delete user usuario
func DeleteUser(c echo.Context) error {
	DB := db.DBManager()
	user := models.Usuario{}
	id := c.Param("id")
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	DB.Delete(&user, id)
	return c.String(http.StatusOK, "Usuario eliminado")
}
