package main

import (
	"flag"
	"fmt"

	"github.com/DeathWofl/rally-api/db"
	"github.com/DeathWofl/rally-api/migration.go"
)

func main() {
	db.Init()

	var migrate string

	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion de la base de datos.")

	flag.Parse()

	if migrate == "yes" {
		fmt.Println("Comenzo la Migracion...")
		migration.Migrate()
		fmt.Println("Termino la Migracion...")
	}
}
