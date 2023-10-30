package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/api/handlers"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/api/routes"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServerHTTP(cfg *config.Config, userHandler handlers.UserHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.RegisterUserRoutes(engine.Group("/"), userHandler)
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"statuscode": 404,
			"message":    "invalid url",
		})
	})

	return &Server{
		Engine: engine,
		Port:   cfg.Port,
	}, nil
}

func (cfg *Server) Start() {
	cfg.Engine.Run(cfg.Port)
}
