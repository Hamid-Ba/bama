package cache

import (
	"context"
	"time"

	"github.com/Hamid-Ba/bama/config"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitRedisClient(cfg *config.Config) {
	// Set up Redis client
	client = redis.NewClient(
		&redis.Options{
			Addr:         cfg.Redis.Host + ":" + cfg.Redis.Port,
			Password:     cfg.Redis.Password, // no password set
			DB:           cfg.Redis.Db,       // use default DB
			DialTimeout:  cfg.Redis.DialTimeout * time.Second,
			ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
			WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
			PoolSize:     cfg.Redis.PoolSize,
			PoolTimeout:  cfg.Redis.PoolTimeout,
		},
	)

	ctx := context.Background()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	println("PING Redis: " + cfg.Redis.Host + ":" + cfg.Redis.Port)
	println("Redis client initialized successfully")
}

func GetRedisClient() *redis.Client {
	if client == nil {
		panic("Redis client is not initialized. Please call SetRedisClient first.")
	}
	return client
}

func CloseRedisClient() {
	if client != nil {
		err := client.Close()
		if err != nil {
			panic("Failed to close Redis client: " + err.Error())
		}
		client = nil
	}
}
