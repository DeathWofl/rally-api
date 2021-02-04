package models

import "github.com/jinzhu/gorm"

//Usuario representa cada entidad que tendra acceso al sistema.
// Digase maestros o administradores.
type Usuario struct {
	gorm.Model
	Nombre   string `json:"Nombre" gorm:"type:varchar(200);not null"`
	Username string `json:"Username" gorm:"type:varchar(100);not null"`
	Password string `json:"Password" gorm:"type:varchar(100); not null;"`
}

// UsuarioService metodos disponibles para usuario
type UsuarioService interface {
	Usuario(id uint) (*Usuario, error)
	Usuarios() (*[]Usuario, error)
	CreateUsuario(u *Usuario) (*Usuario, error)
	UpdateUsuario(id uint, u *Usuario) (*Usuario, error)
	DeleteUsuario(id uint) error
}
