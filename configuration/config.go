package configuration

import (
	"encoding/json"
	"log"
	"os"
)

//GetConfiguration obtiene los datos para acceso a la BD
func GetConfiguration() Configuration {

	var c Configuration

	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
