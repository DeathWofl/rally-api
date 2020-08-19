package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB conexion en la BD
var DB *gorm.DB
var err error

//Init Inicia la conexion con la base de datos
func Init() *gorm.DB {

	userDB := "u7ngap3fewbbqahv"
	userPassword := "5yhtxXIgpHtb1k6Ba06h"
	DBHost := "by6b42vzghn5z3lwdaqg-mysql.services.clever-cloud.com"
	DBName := "by6b42vzghn5z3lwdaqg"
	// database string connection
	// dsc := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Database)
	dsc := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", userDB, userPassword, DBHost, DBName)

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
