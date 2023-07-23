package db

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/syamsv/apollo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbMutex sync.Mutex

const maxRetries = 5
const retryDelay = 3 * time.Second

func GetDbInstance() (*gorm.DB, error) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return nil, fmt.Errorf("error getting DB connection: %w", err)
		}
		err = sqlDB.Ping()
		if err == nil {
			return db, nil
		}
		if err := sqlDB.Close(); err != nil {
			return nil, fmt.Errorf("error closing DB connection: %w", err)
		}
	}

	var err error
	for retry := 1; retry <= maxRetries; retry++ {
		db, err = connect()
		if err == nil {
			return db, nil
		}
		log.Printf("Database connection failed (Attempt %d/%d): %s", retry, maxRetries, err)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("failed to establish a database connection after %d attempts", maxRetries)
}

func connect() (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.POSTGRES_HOST, config.POSTGRES_USER, config.POSTGRES_PASS, config.POSTGRES_DB, config.POSTGRES_PORT)
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
}
