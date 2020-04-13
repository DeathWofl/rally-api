package configuration

import (
	"encoding/json"
	"log"
	"os"
)

//GetConfiguration obtiene los datos para acceso a la BD
func GetConfiguration() {

	var c Configuration

	file, err := os.File("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	err := json.Encode(File).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
