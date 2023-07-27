package rest

import (
	"fmt"
	"go-rest/config"
	"go-rest/svc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	svc       svc.Service
	appConfig *config.Application
}

func NewServer(svc svc.Service, appConfig *config.Application) (*Server, error) {
	server := &Server{
		svc:       svc,
		appConfig: appConfig,
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	s.router = router

	router.POST("/api/users", s.createUser)
	router.GET("/api/users/:id", s.getUser)
	router.PATCH("/api/users/:id", s.updateUser)
	router.DELETE("/api/users/:id", s.deleteUser)
}

func (s *Server) Start() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.appConfig.Host, s.appConfig.Port))
}
