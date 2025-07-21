// @title           Rarible
// @version         1.0
// @description     Rarible client, and Service to interact with the Rarible API.

package main

import (
	_ "Rarible/cmd/api/docs"
	"Rarible/internal/config"
	"Rarible/internal/logger"
	"Rarible/internal/server"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	_ = godotenv.Load() // Підвантажує .env у ENV для локального запуску
	ctx := context.Background()

	// Load configuration
	cfg, err := config.New()
	if err != nil {
		logger.Fatal(ctx, err)
	}

	// Initialize server
	srv := server.NewServer(cfg)

	// Add Swagger UI endpoint
	srv.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO: Register Rarible handlers here

	// Start server
	go srv.Run(ctx)

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info(ctx, "Shutting down server...")
}
