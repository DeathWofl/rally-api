package mysql

import (
	"github.com/DeathWofl/rally-api/pkg/models"
	"github.com/jinzhu/gorm"
)

type PuntuacionService struct {
	DB *gorm.DB
}

func (s *PuntuacionService) Puntuaciones() (*[]models.Puntuacion, error) {
	Puntuaciones := []models.Puntuacion{}

	rows, err := s.DB.Raw("SELECT e.id, TIMEDIFF(MAX(t.hora_llegada), MIN(t.hora_llegada)) AS transcurso,AVG(pr.calificacion) AS puntaje FROM equipos AS e INNER JOIN reg_tiempos AS t ON e.id=t.equipo_id INNER JOIN reg_resps AS pr ON e.id=pr.equipo_id group by e.id order by Puntaje DESC, transcurso ASC;").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		g := new(models.Puntuacion)
		err := rows.Scan(&g.ID, &g.Transcurso, &g.Puntaje)
		if err != nil {
			return nil, err
		}

		Puntuaciones = append(Puntuaciones, *g)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &Puntuaciones, nil
}
