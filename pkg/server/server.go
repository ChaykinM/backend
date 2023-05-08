package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"main.go/pkg/database"
	// "../database"
)

type Server struct {
	database *database.Database
	router   *gin.Engine
}

func New(database *database.Database, router *gin.Engine) *Server {
	return &Server{
		database: database,
		router:   router,
	}
}

func (s *Server) Start() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	s.router.Use(cors.New(corsConfig))
	s.routerSettings()
	s.router.Run(":5050")
}
