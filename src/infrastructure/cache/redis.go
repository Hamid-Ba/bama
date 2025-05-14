package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/pkg/logging"
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
		logging.Log.Error("Failed to connect to Redis", logging.Field{Key: "error", Value: err})
		panic("Failed to connect to Redis: " + err.Error())
	}

	logging.Log.Info("PING Redis: " + cfg.Redis.Host + ":" + cfg.Redis.Port)
	println("PING Redis: " + cfg.Redis.Host + ":" + cfg.Redis.Port)
	logging.Log.Info("Redis client initialized successfully")
	println("Redis client initialized successfully")
}

func GetRedisClient() *redis.Client {
	if client == nil {
		logging.Log.Error("Redis client is not initialized. Please call SetRedisClient first.")
		panic("Redis client is not initialized. Please call SetRedisClient first.")
	}
	return client
}

func Get[T any](c *redis.Client, key string) (T, error) {
	var dest T = *new(T)
	v, err := c.Get(context.Background(), key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(v), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}

func Set[T any](c *redis.Client, key string, value T, duration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(context.Background(), key, v, duration).Err()
}

func CloseRedisClient() {
	if client != nil {
		err := client.Close()
		if err != nil {
			logging.Log.Error("Failed to close Redis client: " + err.Error())
			panic("Failed to close Redis client: " + err.Error())
		}
		client = nil
	}
}
