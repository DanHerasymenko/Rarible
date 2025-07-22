// @title           Rarible
// @version         1.0
// @description     Rarible client, and Service to interact with the Rarible API.

package main

import (
	_ "RaribleAPI/cmd/api/docs"
	"RaribleAPI/internal/client"
	"RaribleAPI/internal/config"
	"RaribleAPI/internal/handler"
	"RaribleAPI/internal/logger"
	"RaribleAPI/internal/server"
	"RaribleAPI/internal/service"
	"context"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Load .env to local ENV
	_ = godotenv.Load()

	ctx := context.Background()

	// Load config
	cfg, err := config.New()
	if err != nil {
		logger.Fatal(ctx, err)
	}

	// Initialize Client, Service, Handler
	Client := client.NewRaribleClient(cfg.RaribleAPIURL, cfg.RaribleAPIKey)
	Service := service.NewRaribleService(Client)
	Handler := handler.NewRaribleHandler(Service)

	// Initialize server
	srv := server.NewServer(cfg)

	// Add Swagger endpoint
	srv.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	Handler.RegisterRoutes(srv.Router)

	// Start server with graceful shutdown
	srv.Run(ctx)
}
