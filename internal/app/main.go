// Package app configures and runs application.
package app

import (
	"github.com/asargin-dev/soundproof-backend-demo/config"
	"github.com/asargin-dev/soundproof-backend-demo/internal/middleware"
	"github.com/asargin-dev/soundproof-backend-demo/internal/models"
	"github.com/asargin-dev/soundproof-backend-demo/internal/routes"
	"github.com/asargin-dev/soundproof-backend-demo/pkg/logger"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {

	// Logger
	log := logger.New("info")

	// Server
	engine := gin.New()

	// Middleware
	engine.Use(middleware.JwtAuthMiddleware())

	// API Routes
	routes.AddRoutes(engine)

	// Database
	models.ConnectDataBase(cfg.PG.Url, log)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	err := engine.Run(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("App could not be started")
	}
}
