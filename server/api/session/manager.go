package session

import (
	"log"
	"time"

	"github.com/syamsv/apollo/config"
	"github.com/syamsv/apollo/pkg/cache"
)

type Manager struct {
	Api    *cache.Cache
	Verify *cache.Cache
}

var manager Manager

func InitSession() {
	verifyConfig := cache.RedisConfig{
		Host:         config.REDIS_HOST,
		Port:         config.REDIS_PORT,
		Password:     config.REDIS_PASSWORD,
		DB:           0,
		MaxRetries:   3,
		RetryBackoff: 2 * time.Second,
	}
	apiConfig := cache.RedisConfig{
		Host:         config.REDIS_HOST,
		Port:         config.REDIS_PORT,
		Password:     config.REDIS_PASSWORD,
		DB:           1,
		MaxRetries:   3,
		RetryBackoff: 2 * time.Second,
	}

	apiCache, err := cache.NewCache(apiConfig)
	if err != nil {
		log.Fatal(err)
	}
	verifyCache, err := cache.NewCache(verifyConfig)
	if err != nil {
		log.Fatal(err)
	}

	manager = Manager{
		Api:    apiCache,
		Verify: verifyCache,
	}
}
