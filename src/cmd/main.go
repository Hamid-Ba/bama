package main

import (
	"fmt"
	"log"

	"github.com/Hamid-Ba/bama/api"
	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/infrastructure/cache"
	"github.com/Hamid-Ba/bama/infrastructure/db"
	"github.com/Hamid-Ba/bama/pkg/logging"
)

func main() {
	cfg := config.GetConfig()

	zap_log, err := logging.NewLogger(cfg.Logger)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}
	logging.Log = zap_log
	defer logging.Log.Sync()

	cache.InitRedisClient(cfg)
	defer cache.CloseRedisClient()

	err = db.InitDb(cfg)
	defer db.CloseDb()

	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
