package usuarios

import (
	"github.com/spinales/rally-api/internal/platform/logger"
	"gorm.io/gorm"
)

// Usuario representa cada entidad que tendra acceso al sistema.
// Digase maestros o administradores.
type Usuario struct {
	gorm.Model
	Nombre   string `json:"Nombre" gorm:"type:varchar(200);not null"`
	Username string `json:"Username" gorm:"type:varchar(100);not null"`
	Password string `json:"Password" gorm:"type:varchar(100); not null;"`
}

type Usuarios struct {
	logHandler logger.Logger
	store      store
}

func (us *Usuarios) UsuarioPorID(id uint) (*Usuario, error) {
	u, err := us.store.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *Usuarios) TodosUsuarios() (*[]Usuario, error) {
	u, err := us.store.All()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *Usuarios) CrearUsuario(u *Usuario) (*Usuario, error) {
	u, err := us.store.Create(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *Usuarios) ActualizarUsuario(id uint, u *Usuario) (*Usuario, error) {
	u, err := us.store.Update(id, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *Usuarios) EliminarUsuario(id uint) error {
	err := us.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewService(l logger.Logger, db *gorm.DB) (*Usuarios, error) {
	ustore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Usuarios{
		logHandler: l,
		store:      ustore,
	}, nil
}
