package server

import (
	"log"
	"os"

	"github.com/0resama/go-clean-ddd/internal/domain/repository"
	"github.com/0resama/go-clean-ddd/internal/domain/usecase/user"
	"github.com/0resama/go-clean-ddd/internal/infrastructure/logger"
	"github.com/0resama/go-clean-ddd/internal/interface/http/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(db *gorm.DB, logger *logger.ZapLogger) *Server {
	userRepo := repository.NewUserGormRepository(db)
	createUesr := user.NewCreateUser(userRepo, logger)
	getUserById := user.NewGetUserByIDUseCase(userRepo, logger)
	userHandler := handler.NewUserHandler(createUesr, getUserById)

	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	engine.POST("/users", userHandler.CreateUser)
	engine.GET("/users/:id", userHandler.GetUserByID)

	return &Server{engine: engine}
}

func (s *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := s.engine.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}
