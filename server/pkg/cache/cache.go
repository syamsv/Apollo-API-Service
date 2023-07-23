package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host         string
	Port         string
	Password     string
	DB           int
	MaxRetries   int
	RetryBackoff time.Duration
}

type Cache struct {
	rdb *redis.Client
	ctx context.Context
}

func NewCache(config RedisConfig) (*Cache, error) {
	if config.MaxRetries <= 0 {
		config.MaxRetries = 3
	}
	if config.RetryBackoff <= 0 {
		config.RetryBackoff = 1 * time.Second
	}

	rdb, err := connectRedis(config)
	if err != nil {
		return nil, err
	}

	return &Cache{
		rdb: rdb,
		ctx: context.Background(),
	}, nil
}

func connectRedis(config RedisConfig) (*redis.Client, error) {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	}

	var rdb *redis.Client
	var err error

	for retries := 0; retries <= config.MaxRetries; retries++ {
		rdb = redis.NewClient(options)

		_, err = rdb.Ping(context.Background()).Result()
		if err == nil {
			log.Printf("Connected to Redis server successfully.")
			return rdb, nil
		}

		log.Printf("Failed to connect to Redis: %v", err)
		if retries < config.MaxRetries {
			log.Printf("Retrying in %v...", config.RetryBackoff)
			time.Sleep(config.RetryBackoff)
		}
	}

	return nil, fmt.Errorf("failed to establish connection to Redis after %d retries", config.MaxRetries)
}

func (c *Cache) SetValue(key, value string, expiration time.Duration) error {
	err := c.rdb.Set(c.ctx, key, value, expiration).Err()
	return err
}

func (c *Cache) GetValue(key string) (string, error) {
	val, err := c.rdb.Get(c.ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (c *Cache) DeleteValue(key string) error {
	err := c.rdb.Del(c.ctx, key).Err()
	if err == redis.Nil {
		return nil
	}
	return err
}

func (c *Cache) Close() error {
	if c.rdb != nil {
		return c.rdb.Close()
	}
	return nil
}
