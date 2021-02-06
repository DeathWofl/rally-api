package mysql

import (
	"github.com/DeathWofl/rally-api/pkg/models"
	"gorm.io/gorm"
)

//Migrate migracion a la base de datos
func Migrate(DB *gorm.DB) {
	//Las borras en caso de que existan
	DB.Migrator().DropTable(&models.Equipo{}, &models.Estacion{}, &models.Respuesta{}, &models.Pregunta{}, &models.Usuario{}, &models.RegResp{}, &models.RegTiempo{})

	//Crea las tablas
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Equipo{}, &models.Estacion{}, &models.Usuario{}, &models.Pregunta{})

	//Relacion
	// DB.Model(&models.Pregunta{}).AddForeignKey("estacion_id", "estacions(id)", "Cascade", "Cascade")
	DB.Migrator().CreateConstraint(&models.Estacion{}, "pregunta")
	DB.Migrator().CreateConstraint(&models.Estacion{}, "fk_estaciones_preguntas")

	DB.AutoMigrate(&models.Respuesta{})
	//Relacion
	// DB.Model(&models.Respuesta{}).AddForeignKey("pregunta_id", "pregunta(id)", "Cascade", "Cascade")
	DB.Migrator().CreateConstraint(&models.Pregunta{}, "Respuesta")
	DB.Migrator().CreateConstraint(&models.Pregunta{}, "fk_pregunta_respuestas")

	DB.AutoMigrate(&models.RegResp{}, &models.RegTiempo{})

	// DB.Model(&models.RegTiempo{}).AddForeignKey("estacion_id", "estacions(id)", "Cascade", "Cascade")
	DB.Migrator().CreateConstraint(&models.Estacion{}, "RegTiempos")
	DB.Migrator().CreateConstraint(&models.Estacion{}, "fk_estaciones_reg_tiempos")
	// DB.Model(&models.RegResp{}).AddForeignKey("equipo_id", "equipos(id)", "Cascade", "Cascade")
	DB.Migrator().CreateConstraint(&models.Equipo{}, "RegRespuestas")
	DB.Migrator().CreateConstraint(&models.Equipo{}, "fk_equipos_reg_respuestas")
	// DB.Model(&models.RegTiempo{}).AddForeignKey("equipo_id", "equipos(id)", "Cascade", "Cascade")
	DB.Migrator().CreateConstraint(&models.Equipo{}, "RegTiempos")
	DB.Migrator().CreateConstraint(&models.Equipo{}, "fk_equipos_reg_tiempos")
	// DB.Model(&models.RegResp{}).AddForeignKey("pregunta_id", "pregunta(id)", "Cascade", "Cascade")
	DB.Migrator().CreateConstraint(&models.Pregunta{}, "RegRespuestas")
	DB.Migrator().CreateConstraint(&models.Pregunta{}, "fk_pregunta_reg_respuesta")

	// Create a admin user
	user := models.Usuario{Nombre: "Admin", Username: "admin", Password: "admin"}
	DB.Create(&user)
	equip := models.Equipo{MatriculaE1: "2016-0001", MatriculaE2: "2016-0002", MatriculaE3: "2016-0003", ContraGrupo: "Pass123", CodigoGrupo: "123456789"}
	DB.Create(&equip)

}
