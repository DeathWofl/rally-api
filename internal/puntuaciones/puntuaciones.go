package puntuaciones

import "gorm.io/gorm"

// Puntuacion representa cada puntuacion de cada equipo
// antiguamente representado como ganador
type Puntuacion struct {
	gorm.Model
	Transcurso string
	Puntaje    float32
}
