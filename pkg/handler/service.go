// Package handler aqui van todos los controladores o handler
// que corresponden a cada endpoint
package handler

import (
	"github.com/DeathWofl/rally-api/pkg/storage/mysql"
)

// Service Aqui van todas las dependencias necesarias para la app
// funcionar, digase capa de datos o algun servicio externo / tercero
type Service struct {
	EquipoService     *mysql.EquipoService
	EstacionService   *mysql.EstacionService
	PreguntaService   *mysql.PreguntaService
	PuntuacionService *mysql.PuntuacionService
	RegRespService    *mysql.RegRespService
	RegTiempoService  *mysql.RegTiempoService
	RespuestaService  *mysql.RespuestaService
	UsuarioService    *mysql.UsuarioService
}
