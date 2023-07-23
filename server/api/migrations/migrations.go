package migrations

import (
	"log"

	"github.com/syamsv/apollo/api/db"
	"github.com/syamsv/apollo/pkg/models"
)

func Migrate() {
	database, err := db.GetDbInstance()
	if err != nil {
		log.Fatal("[Migrations] Migrations Failed : Try again later")
	}
	database.AutoMigrate(&models.Users{})
}
