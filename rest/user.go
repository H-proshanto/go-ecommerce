package rest

import (
	"encoding/json"
	"go-rest/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := s.svc.GetUser(userId)

	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data":    user,
			"success": true,
			"message": "User retrieved",
		},
	)
}

func (s *Server) createUser(ctx *gin.Context) {
	var user svc.User
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	message, responseError := s.svc.CreateUser(&user)

	if responseError != nil {
		ctx.JSON(responseError.StatusCode, gin.H{"error": responseError.Message})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": message,
		},
	)
}

func (s *Server) updateUser(ctx *gin.Context) {
	var user svc.User

	userId := ctx.Param("id")
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	updatedUser, responseError := s.svc.UpdateUser(userId, &user)

	if responseError != nil {
		ctx.JSON(responseError.StatusCode, gin.H{"error": responseError.Message})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data":    updatedUser,
			"success": true,
			"message": "Updated Successfully",
		},
	)
}

func (s *Server) deleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	message, err := s.svc.DeleteUser(userId)

	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": message,
		},
	)
}
