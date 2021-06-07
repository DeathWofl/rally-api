package preguntas

import "gorm.io/gorm"

type store interface {
	ReadByID(id uint) (*Pregunta, error)
	All() (*[]Pregunta, error)
	Create(p *Pregunta) (*Pregunta, error)
	Update(id uint, p *Pregunta) (*Pregunta, error)
	Delete(id uint) error
}

type preguntaStore struct {
	db *gorm.DB
}

func (s *preguntaStore) ReadByID(id uint) (*Pregunta, error) {
	var pregunta Pregunta
	s.db.First(&pregunta, id)
	return &pregunta, nil
}

func (s *preguntaStore) All() (*[]Pregunta, error) {
	var preguntas []Pregunta
	s.db.Find(&preguntas)
	return &preguntas, nil
}

func (s *preguntaStore) Create(p *Pregunta) (*Pregunta, error) {
	s.db.Create(p)
	return p, nil
}

func (s *preguntaStore) Update(id uint, p *Pregunta) (*Pregunta, error) {
	var pregunta Pregunta
	s.db.Find(&pregunta, id)
	s.db.Model(&pregunta).Updates(p)
	return &pregunta, nil
}

func (s *preguntaStore) Delete(id uint) error {
	s.db.Delete(&Pregunta{}, id)
	return nil
}

func newStore(db *gorm.DB) (*preguntaStore, error) {
	return &preguntaStore{
		db: db,
	}, nil
}
