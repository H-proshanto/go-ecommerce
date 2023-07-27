package svc

import "go-rest/utils"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) GetUser(id string) (*User, *utils.ServerError) {
	return s.userRepo.GetUser(id)
}

func (s *service) CreateUser(user *User) (string, *utils.ServerError) {
	return s.userRepo.CreateUser(user)
}

func (s *service) UpdateUser(id string, user *User) (*User, *utils.ServerError) {
	return s.userRepo.UpdateUser(id, user)
}

func (s *service) DeleteUser(id string) (string, *utils.ServerError) {
	return s.userRepo.DeleteUser(id)
}
