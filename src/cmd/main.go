package main

import (
	"log"

	"github.com/Hamid-Ba/bama/api"
	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/data/cache"
	"github.com/Hamid-Ba/bama/data/db"
)

func main() {
	cfg := config.GetConfig()

	cache.InitRedisClient(cfg)
	defer cache.CloseRedisClient()

	err := db.InitDb(cfg)
	defer db.CloseDb()

	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
