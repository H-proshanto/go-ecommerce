package svc

import "go-rest/utils"

type UserRepo interface {
	CreateUser(user *User) (string, *utils.ServerError)
	GetUser(id string) (*User, *utils.ServerError)
	UpdateUser(id string, std *User) (*User, *utils.ServerError)
	DeleteUser(id string) (string, *utils.ServerError)
}

type Service interface {
	GetUser(id string) (*User, *utils.ServerError)
	CreateUser(user *User) (string, *utils.ServerError)
	UpdateUser(id string, std *User) (*User, *utils.ServerError)
	DeleteUser(id string) (string, *utils.ServerError)
}
