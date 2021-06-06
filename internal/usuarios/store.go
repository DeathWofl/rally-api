package usuarios

import (
	"github.com/jinzhu/gorm"
)

// service metodos disponibles para usuario
type store interface {
	Usuario(id uint) (*Usuario, error)
	Usuarios() (*[]Usuario, error)
	CreateUsuario(u *Usuario) (*Usuario, error)
	UpdateUsuario(id uint, u *Usuario) (*Usuario, error)
	DeleteUsuario(id uint) error
}

type usuarioStore struct {
	db *gorm.DB
}

// Usuario retorna el usuario pedido
func (s *usuarioStore) Usuario(id uint) (*Usuario, error) {
	var usuario Usuario
	s.db.First(&usuario, id)
	return &usuario, nil
}

// Usuarios retorna todos los ususrios de la BD
func (s *usuarioStore) Usuarios() (*[]Usuario, error) {
	var usuarios []Usuario
	s.db.Find(&usuarios)
	return &usuarios, nil
}

// CreateUsuario agrega un nuevo usuario en la BD
func (s *usuarioStore) CreateUsuario(u *Usuario) (*Usuario, error) {
	s.db.Create(u)
	return u, nil
}

// UpdateUsuario actualiza un usuario en la BD
func (s *usuarioStore) UpdateUsuario(id uint, u *Usuario) (*Usuario, error) {
	var usuario Usuario
	s.db.Find(&usuario, id)
	s.db.Model(&usuario).Updates(u)
	return &usuario, nil
}

// DeleteUsuario elimina usuario de la BD
func (s *usuarioStore) DeleteUsuario(id uint) error {
	s.db.Delete(&Usuario{}, id)
	return nil
}

func newService(db *gorm.DB) (*usuarioStore, error) {
	return &usuarioStore{
		db: db,
	}, nil
}
