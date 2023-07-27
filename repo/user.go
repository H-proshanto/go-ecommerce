package repo

import (
	"fmt"
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

func (r *userRepo) CreateUser(user *svc.User) (string, *utils.ServerError) {
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

func (r *userRepo) GetUser(id string) (*svc.User, *utils.ServerError) {
	var user *svc.User

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
	return user, nil
}

func (r *userRepo) UpdateUser(id string, newUser *svc.User) (*svc.User, *utils.ServerError) {
	var user *svc.User
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

	if newUser.FirstName != "" {
		user.FirstName = newUser.FirstName
	}

	if newUser.LastName != "" {
		user.LastName = newUser.LastName
	}

	if newUser.Email != "" {
		user.Email = newUser.Email
	}

	user.UpdatedAt = time.Now()

	r.db.Save(&user)

	user.Password = ""

	return user, nil
}

func (r *userRepo) DeleteUser(id string) (string, *utils.ServerError) {
	var user *svc.User
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
