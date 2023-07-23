package main

import (
	"log"

	"github.com/syamsv/apollo/api/cmd"
	"github.com/syamsv/apollo/api/db"
	"github.com/syamsv/apollo/api/migrations"
	"github.com/syamsv/apollo/api/session"
	"github.com/syamsv/apollo/config"
)

func main() {
	config.LoadConfig()

	session.InitSession()

	if _, err := db.GetDbInstance(); err != nil {
		log.Fatalln("Error connecting to database", err)
	}
	db.InitDatabaseInstannce()
	if config.MIGRATE {
		migrations.Migrate()
	}

	cmd.StartServer()
}
