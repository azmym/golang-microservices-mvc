package domain

import (
	"fmt"
	"golang-microservices-mvc/utils"
	"net/http"
)

type (
	usersList               map[int64]*User
	UserRepositoryInterface interface {
		FindUserById(int64) (*User, *utils.ApplicationError)
	}
	userRepository struct {
	}
)

var (
	users = usersList{
		123: {Id: 123, FirstName: "Hello", LastName: "world", Email: "Hello_world@yahoo.com"},
		124: {Id: 124, FirstName: "Hello", LastName: "world2", Email: "Hello_world2@yahoo.com"},
		125: {Id: 125, FirstName: "Hello", LastName: "world3", Email: "Hello_world3@yahoo.com"},
		126: {Id: 126, FirstName: "Hello", LastName: "world4", Email: "Hello_world4@yahoo.com"},
	}
	//expose the struct outside of the package
	UserRepository UserRepositoryInterface
)

// initialize the UserRepository
func init() {
	UserRepository = &userRepository{}
}
func (li usersList) findById(userId int64) (bool, *User) {
	user, v := li[userId]
	return v, user
}

func (ur *userRepository) FindUserById(userId int64) (*User, *utils.ApplicationError) {
	found, user := users.findById(userId)
	if !found {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("there is no user with id= %v", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
