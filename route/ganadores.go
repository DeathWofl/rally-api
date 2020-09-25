package route

import (
	"encoding/json"
	"net/http"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/models"
	"github.com/labstack/echo"
)

func GetGanadores(c echo.Context) error {
	DB := db.DBManager()
	ganadores := []*models.Ganador{}

	rows, err := DB.Raw("SELECT e.id, TIMEDIFF(MAX(t.hora_llegada), MIN(t.hora_llegada)) AS transcurso,AVG(pr.calificacion) AS puntaje FROM equipos AS e INNER JOIN reg_tiempos AS t ON e.id=t.equipo_id INNER JOIN reg_resps AS pr ON e.id=pr.equipo_id group by e.id order by Puntaje DESC, transcurso ASC;").Rows()
	if err != nil {
		c.Logger().Print(err)
	}
	defer rows.Close()
	for rows.Next() {

		g := new(models.Ganador)
		err := rows.Scan(&g.ID, &g.Transcurso, &g.Puntaje)
		c.Logger().Print(g)
		if err != nil {
			c.Logger().Print(err)
		}

		ganadores = append(ganadores, g)
	}

	if err := rows.Err(); err != nil {
		c.Logger().Print(err)
	}

	if _, err := json.Marshal(ganadores); err != nil {
		c.Logger().Print(err)
	}

	return c.JSON(http.StatusOK, ganadores)
}

//GetEquipo Retorna equipo por ID
func GetGanadores2(c echo.Context) error {
	DB := db.DBManager()
	ganadores := []models.Ganador{}

	DB.Model(models.Equipo{}).
		Select("equipos.id AS ID, TIMEDIFF(MAX(reg_tiempos.hora_llegada),MIN(reg_tiempos.hora_llegada)) AS Transcurso, AVG(reg_resps.calificacion) AS Puntaje").
		Joins("INNER JOIN reg_tiempos ON equipos.id=reg_tiempos.equipo_id").
		Joins("INNER JOIN reg_resps ON equipos.id=reg_resps.equipo_id").
		Group("equipos.id").
		Order("Puntaje DESC, Transcurso ASC").
		Find(&ganadores)

	return c.JSON(http.StatusOK, ganadores)
}
