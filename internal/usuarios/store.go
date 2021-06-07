package usuarios

import (
	"gorm.io/gorm"
)

type store interface {
	ReadByID(id uint) (*Usuario, error)
	All() (*[]Usuario, error)
	Create(u *Usuario) (*Usuario, error)
	Update(id uint, u *Usuario) (*Usuario, error)
	Delete(id uint) error
}

type usuarioStore struct {
	db *gorm.DB
}

func (s *usuarioStore) ReadByID(id uint) (*Usuario, error) {
	var usuario Usuario
	s.db.First(&usuario, id)
	return &usuario, nil
}

func (s *usuarioStore) All() (*[]Usuario, error) {
	var usuarios []Usuario
	s.db.Find(&usuarios)
	return &usuarios, nil
}

func (s *usuarioStore) Create(u *Usuario) (*Usuario, error) {
	s.db.Create(u)
	return u, nil
}

func (s *usuarioStore) Update(id uint, u *Usuario) (*Usuario, error) {
	var usuario Usuario
	s.db.Find(&usuario, id)
	s.db.Model(&usuario).Updates(u)
	return &usuario, nil
}

func (s *usuarioStore) Delete(id uint) error {
	s.db.Delete(&Usuario{}, id)
	return nil
}

func newStore(db *gorm.DB) (*usuarioStore, error) {
	return &usuarioStore{
		db: db,
	}, nil
}
