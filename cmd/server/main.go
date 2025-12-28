package main

import (
	"log"
	"w3labs/internal/adapters"
	"w3labs/internal/config"

	"github.com/caarlos0/env/v11"
)

func main() {

	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse configs. %+v", err)
	}

	logger := adapters.NewSlog(cfg.Logger)
	logger.Debug("APPLICATION HAS STARTED", nil)

}
