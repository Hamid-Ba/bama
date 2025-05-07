package main

import (
	"github.com/Hamid-Ba/bama/api"
	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/data/cache"
)

func main() {
	cfg := config.GetConfig()

	cache.InitRedisClient(cfg)

	api.InitServer(cfg)
}
