package svc

import (
	"go-rest/dto"
	"go-rest/utils"
)

type UserRepo interface {
	CreateUser(req *dto.UserRequestBody) (string, *utils.ServerError)
	GetUser(req *dto.UserRequestBody) (*dto.UserResponseBody, *utils.ServerError)
	UpdateUser(req *dto.UserRequestBody) (*dto.UserResponseBody, *utils.ServerError)
	DeleteUser(req *dto.UserRequestBody) (string, *utils.ServerError)
}

type Service interface {
	CreateUser(req *dto.UserRequestBody) (string, *utils.ServerError)
	GetUser(req *dto.UserRequestBody) (*dto.UserResponseBody, *utils.ServerError)
	UpdateUser(req *dto.UserRequestBody) (*dto.UserResponseBody, *utils.ServerError)
	DeleteUser(req *dto.UserRequestBody) (string, *utils.ServerError)
}
