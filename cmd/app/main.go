package main

import (
	"github.com/asargin-dev/soundproof-backend-demo/config"
	"github.com/asargin-dev/soundproof-backend-demo/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	// Run
	app.Run(cfg)
}
