package svc

import (
	"go-rest/dto"
	"go-rest/utils"
)

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) CreateUser(req *dto.UserRequestBody) (string, *utils.ServerError) {
	return s.userRepo.CreateUser(req)
}

func (s *service) GetUser(req *dto.UserRequestBody) (*dto.UserResponseBody, *utils.ServerError) {
	return s.userRepo.GetUser(req)
}

func (s *service) UpdateUser(req *dto.UserRequestBody) (*dto.UserResponseBody, *utils.ServerError) {
	return s.userRepo.UpdateUser(req)
}

func (s *service) DeleteUser(req *dto.UserRequestBody) (string, *utils.ServerError) {
	return s.userRepo.DeleteUser(req)
}
