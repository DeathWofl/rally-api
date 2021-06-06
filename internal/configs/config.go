package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	user     string
	password string
	port     string
	host     string
	name     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	port = os.Getenv("DB_PORT")
	host = os.Getenv("DB_HOST")
	name = os.Getenv("DB_NAME")
}

func ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
}
