package usuarios

import (
	"github.com/jinzhu/gorm"
)

// Usuario representa cada entidad que tendra acceso al sistema.
// Digase maestros o administradores.
type Usuario struct {
	gorm.Model
	Nombre   string `json:"Nombre" gorm:"type:varchar(200);not null"`
	Username string `json:"Username" gorm:"type:varchar(100);not null"`
	Password string `json:"Password" gorm:"type:varchar(100); not null;"`
}

// type Usuarios struct {
//     logHandler logger.CustomLogger
//     store store
// }
