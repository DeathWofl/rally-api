package db

import (
	"fmt"
	"log"

	"github.com/DeathWofl/rally-api/configuration"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB conexion en la BD
var DB *gorm.DB
var err error

//Init Inicia la conexion con la base de datos
func Init() *gorm.DB {
	c := configuration.GetConfiguration()

	// database string connection
	// dsc := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Database)
	dsc := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Database)

	DB, err = gorm.Open("mysql", dsc)
	if err != nil {
		log.Fatal(err)
	}

	return DB
}

//DBManager retorna la  conexion abierta
func DBManager() *gorm.DB {
	return DB
}
