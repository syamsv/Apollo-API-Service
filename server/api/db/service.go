package db

import (
	"log"

	"github.com/syamsv/apollo/pkg/users"
)

var (
	User users.Interface
)

func InitDatabaseInstannce() {
	database, err := GetDbInstance()
	if err != nil {
		log.Fatal("error while connecting to database", err)
	}
	userRepo := users.NewRepository(database)
	User = users.NewService(userRepo)
}
