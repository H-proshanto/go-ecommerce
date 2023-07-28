package repo

import (
	"fmt"
	"go-rest/dto"
	"go-rest/svc"
	"go-rest/utils"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type UserRepo interface {
	svc.UserRepo
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(req *dto.UserRequestBody) (string, *utils.ServerError) {
	user := svc.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	result := r.db.Create(&user)

	if result.Error != nil {
		fmt.Println("something occured", result.Error)
		return "", &utils.ServerError{
			Message:    "something occured",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return "Created Succesfully", nil
}

func (r *userRepo) GetUser(req *dto.UserRequestBody) (*dto.UserResponseBody, *utils.ServerError) {
	var user *svc.User

	id := req.ID
	result := r.db.Where(fmt.Sprintf("id = %s", id)).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found")
			return nil, &utils.ServerError{
				Message:    result.Error.Error(),
				StatusCode: http.StatusNotFound,
			}
		} else {
			fmt.Println("Error while fetching user:", result.Error)
			return nil, &utils.ServerError{
				Message:    result.Error.Error(),
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	user.Password = ""
	return &dto.UserResponseBody{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *userRepo) UpdateUser(req *dto.UserRequestBody) (*dto.UserResponseBody, *utils.ServerError) {
	var user *svc.User

	id := req.ID
	result := r.db.Where(fmt.Sprintf("id = %s", id)).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found")
			return nil, &utils.ServerError{
				Message:    result.Error.Error(),
				StatusCode: http.StatusNotFound,
			}
		} else {
			fmt.Println("Error while fetching user:", result.Error)
			return nil, &utils.ServerError{
				Message:    result.Error.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}

	if req.LastName != "" {
		user.LastName = req.LastName
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	user.UpdatedAt = time.Now()

	r.db.Save(&user)

	user.Password = ""

	return &dto.UserResponseBody{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *userRepo) DeleteUser(req *dto.UserRequestBody) (string, *utils.ServerError) {
	var user *svc.User
	id := req.ID
	result := r.db.Where(fmt.Sprintf("id = %s", id)).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found")
			return "", &utils.ServerError{
				Message:    result.Error.Error(),
				StatusCode: http.StatusNotFound,
			}
		} else {
			fmt.Println("Error while fetching user:", result.Error)
			return "", &utils.ServerError{
				Message:    result.Error.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	result = result.Delete(&user)

	if result.Error != nil {
		fmt.Println("error while deleting:", result.Error)

		return "", &utils.ServerError{
			Message:    "Could not delete",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return "Deleted Successfully", nil
}
