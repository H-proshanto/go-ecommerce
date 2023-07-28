package rest

import (
	"encoding/json"
	"go-rest/dto"
	"go-rest/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	response, responseError := s.svc.CreateUser(
		&dto.UserRequestBody{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	)

	if responseError != nil {
		ctx.JSON(responseError.StatusCode, gin.H{"error": responseError.Message})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": response,
		},
	)
}

func (s *Server) getUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	response, responseError := s.svc.GetUser(&dto.UserRequestBody{ID: userId})

	if responseError != nil {
		ctx.JSON(responseError.StatusCode, gin.H{"error": responseError.Message})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data":    response,
			"message": "User retrieved",
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

	response, responseError := s.svc.UpdateUser(
		&dto.UserRequestBody{
			ID:        userId,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	)

	if responseError != nil {
		ctx.JSON(responseError.StatusCode, gin.H{"error": responseError.Message})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data":    response,
			"message": "Updated Successfully",
		},
	)
}

func (s *Server) deleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	response, responseError := s.svc.DeleteUser(&dto.UserRequestBody{
		ID: userId,
	})

	if responseError != nil {
		ctx.JSON(responseError.StatusCode, gin.H{"error": responseError.Message})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": response,
		},
	)
}
