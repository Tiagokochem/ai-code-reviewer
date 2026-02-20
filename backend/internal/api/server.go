package api

import (
	"github.com/gin-gonic/gin"
	"github.com/test-go/ai-code-reviewer/internal/config"
	"github.com/test-go/ai-code-reviewer/internal/handlers"
	"github.com/test-go/ai-code-reviewer/internal/services"
)

type Server struct {
	router  *gin.Engine
	config  *config.Config
	handler *handlers.CodeReviewHandler
}

func NewServer(cfg *config.Config) *Server {
	// Configurar Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// CORS middleware
	router.Use(corsMiddleware())

	// Criar serviços
	aiService := services.NewAIService(cfg)
	n8nService := services.NewN8NService(cfg)
	reviewService := services.NewReviewService(aiService, n8nService)

	// Criar handlers
	handler := handlers.NewCodeReviewHandler(reviewService)

	// Rotas
	setupRoutes(router, handler)

	return &Server{
		router:  router,
		config:  cfg,
		handler: handler,
	}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func setupRoutes(router *gin.Engine, handler *handlers.CodeReviewHandler) {
	api := router.Group("/api/v1")
	{
		api.POST("/review", handler.ReviewCode)
		api.GET("/health", handler.HealthCheck)
	}

	// Swagger (será configurado depois)
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
