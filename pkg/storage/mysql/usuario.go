package mysql

import (
	"github.com/DeathWofl/rally-api/pkg/models"
	"gorm.io/gorm"
)

// UsuarioService implementando metodos
type UsuarioService struct {
	DB *gorm.DB
}

// Usuario retorna el usuario pedido
func (s *UsuarioService) Usuario(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	s.DB.First(&usuario, id)
	return &usuario, nil
}

// Usuarios retorna todos los ususrios de la BD
func (s *UsuarioService) Usuarios() (*[]models.Usuario, error) {
	var usuarios []models.Usuario
	s.DB.Find(&usuarios)
	return &usuarios, nil
}

// CreateUsuario agrega un nuevo usuario en la BD
func (s *UsuarioService) CreateUsuario(u *models.Usuario) (*models.Usuario, error) {
	s.DB.Create(u)
	return u, nil
}

// UpdateUsuario actualiza un usuario en la BD
func (s *UsuarioService) UpdateUsuario(id uint, u *models.Usuario) (*models.Usuario, error) {
	var usuario models.Usuario
	s.DB.Find(&usuario, id)
	s.DB.Model(&usuario).Updates(u)
	return &usuario, nil
}

// DeleteUsuario elimina usuario de la BD
func (s *UsuarioService) DeleteUsuario(id uint) error {
	s.DB.Delete(&models.Usuario{}, id)
	return nil
}
