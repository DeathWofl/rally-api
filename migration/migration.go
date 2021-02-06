package migration

import (
	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/pkg/models"
)

//Migrate migracion a la base de datos
func Migrate() {

	DB := db.DBManager()

	DB.LogMode(true)

	//Las borras en caso de que existan
	DB.DropTableIfExists(&models.Equipo{}, &models.Estacion{}, &models.Respuesta{}, &models.Pregunta{}, &models.Usuario{}, &models.RegResp{}, &models.RegTiempo{})

	//Crea las tablas
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Equipo{}, &models.Estacion{}, &models.Respuesta{}, &models.Pregunta{}, &models.Usuario{}, &models.RegResp{}, &models.RegTiempo{})

	//Reaciones
	DB.Model(&models.Pregunta{}).AddForeignKey("estacion_id", "estacions(id)", "Cascade", "Cascade")
	DB.Model(&models.RegTiempo{}).AddForeignKey("estacion_id", "estacions(id)", "Cascade", "Cascade")
	DB.Model(&models.RegResp{}).AddForeignKey("equipo_id", "equipos(id)", "Cascade", "Cascade")
	DB.Model(&models.RegTiempo{}).AddForeignKey("equipo_id", "equipos(id)", "Cascade", "Cascade")
	DB.Model(&models.RegResp{}).AddForeignKey("pregunta_id", "pregunta(id)", "Cascade", "Cascade")
	DB.Model(&models.Respuesta{}).AddForeignKey("pregunta_id", "pregunta(id)", "Cascade", "Cascade")

	// Create a admin user
	user := models.Usuario{Nombre: "Admin", Username: "admin", Password: "admin"}
	DB.Create(&user)
	equip := models.Equipo{MatriculaE1: "2016-0001", MatriculaE2: "2016-0002", MatriculaE3: "2016-0003", ContraGrupo: "Pass123", CodigoGrupo: "123456789"}
	DB.Create(&equip)

}
