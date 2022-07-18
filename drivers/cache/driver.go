package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
)

type CacheConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

func (c *CacheConfig) InitCacheDB() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Username: c.Username,
		Password: c.Password,
		Addr:     fmt.Sprintf("%s:%s", c.Host, c.Password),
		DB:       0,
	})

	// TODO Add wait
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	return client
}
