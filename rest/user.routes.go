package rest

import (
	"github.com/gin-gonic/gin"
)

func ServeUserRoutes(router *gin.RouterGroup, s *Server) {
	router.POST("/users", s.createUser)
	router.GET("/users/:id", s.getUser)
	router.PATCH("/users/:id", s.updateUser)
	router.DELETE("/users/:id", s.deleteUser)
}
