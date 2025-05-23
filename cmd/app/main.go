package main

import (
	"log"

	"github.com/0resama/go-clean-ddd/internal/infrastructure/db"
	"github.com/0resama/go-clean-ddd/internal/infrastructure/logger"
	"github.com/0resama/go-clean-ddd/internal/infrastructure/server"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", zap.Error(err))
	}

	zapLogger := logger.NewZapLogger()
	database := db.InitPostgres()
	srv := server.NewServer(database, zapLogger)
	srv.Run()
}
