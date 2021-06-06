package puntuaciones

import (
	"gorm.io/gorm"
)

// puntuacionStore metodos disponibles para Puntuacion
type store interface {
	Puntuaciones() (*[]Puntuacion, error)
}

// puntuacionStore implementado metodos
type puntuacionStore struct {
	db *gorm.DB
}

// Puntuaciones retorna todas las puntuaciones de los equipos,
// ordenados por puntuacion y tiempo durado
func (s *puntuacionStore) Puntuaciones() (*[]Puntuacion, error) {
	Puntuaciones := []Puntuacion{}

	rows, err := s.db.Raw("SELECT e.id, TIMEDIFF(MAX(t.hora_llegada), MIN(t.hora_llegada)) AS transcurso,AVG(pr.calificacion) AS puntaje FROM equipos AS e INNER JOIN reg_tiempos AS t ON e.id=t.equipo_id INNER JOIN reg_resps AS pr ON e.id=pr.equipo_id group by e.id order by Puntaje DESC, transcurso ASC;").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		g := new(Puntuacion)
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

func newService(db *gorm.DB) (*puntuacionStore, error) {
	return &puntuacionStore{
		db: db,
	}, nil
}
